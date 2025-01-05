// package caldera adds the caldera capability for the soarca core.
package caldera

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"strings"
	"time"

	"sigs.k8s.io/yaml"

	"soarca/internal/logger"
	"soarca/pkg/core/capability"
	"soarca/pkg/core/capability/caldera/api/models"
	"soarca/pkg/models/cacao"
	"soarca/pkg/models/execution"
)

// calderaCapability is a struct that implements the ICapability interface for the caldera-cmd capability.
// It is initiated with an ICalderaConnectionFactory, which dictates which kind of connection to use.
// By default, it uses the calderaConnectionFactory.
type calderaCapability struct {
	Factory ICalderaConnectionFactory
}

type Empty struct{}

const (
	calderaResult  = "__soarca_caldera_cmd_result__"
	calderaError   = "__soarca_caldera_cmd_error__"
	capabilityName = "soarca-caldera-cmd"
)

var (
	component = reflect.TypeOf(Empty{}).PkgPath()
	log       *logger.Log
)

func init() {
	log = logger.Logger(component, logger.Info, "", logger.Json)
}

// New is the constructor method for the calderaCapability.
// It takes an ICalderaConnectionFactory which determines the kind of connection to use.
// By default, it uses the calderaConnectionFactory.
func New(factory ICalderaConnectionFactory) *calderaCapability {
	if factory == nil {
		factory = calderaConnectionFactory{}
	}
	return &calderaCapability{factory}
}

// GetType returns the name of the capability.
func (capability *calderaCapability) GetType() string {
	return capabilityName
}

// Execute performs the caldera-cmd capability and handles the business logic of the capability.
// It takes a execution Metadata object and a capability Context object which describe the details of the step.
// It first creates an ability if the step defines a custom ability. Afther that, it creates a single action adversary and starts an operation on the defined agents.
// After that, it waits for the operation to complete and parses the generated facts.
// It returns the generated facts as cacao Variables.
// If anything goes wrong, it returns an empty array of variables and the error.
func (c *calderaCapability) Execute(
	metadata execution.Metadata,
	context capability.Context) (cacao.Variables, error) {

	command := context.Command
	target := context.Target

	connection, err := c.Factory.Create()
	if err != nil {
		log.Error("Could not create a connection to caldera")
		return cacao.NewVariables(), err
	}

	abilityId := ""
	groupName := ""
	createdAbility := false

	// parse the command and create the ability if neccesary
	if command.CommandB64 != "" {
		bytes, err := base64.StdEncoding.DecodeString(command.CommandB64)
		if err != nil {
			return cacao.NewVariables(), err
		}
		// First try unmarshalling the ability with json decoding
		var ability *models.Ability
		if json.Valid(bytes) {
			ability = ParseJsonAbility(bytes)
		} else {
			// If that fails, try unmarshalling the ability with yaml decoding
			ability = ParseYamlAbility(bytes)
		}
		abilityId, err = connection.CreateAbility(ability)
		if err != nil {
			log.Error("Could not create custom Ability")
			return cacao.NewVariables(), err
		}
		createdAbility = true
	} else {
		abilityId = strings.Replace(command.Command, "id: ", "", -1)
	}

	// parse the specific target group
	if target.Type == "security-category" && slices.Contains(target.Category, "caldera") {
		groupName = target.Name
	}

	globalFactSourceId, err := createGlobalFactSourceIfMissing(connection, metadata)

	if err != nil {
		log.Error("Could not create/find a global fact source", err)
		return cacao.NewVariables(), err
	}

	// create an adversary
	adversaryId, err := connection.CreateAdversary(abilityId)
	if err != nil {
		log.Error("Could not create the Adversary", err)
		return cacao.NewVariables(), err
	}

	// start the operation
	operationId, err := connection.CreateOperation(groupName, adversaryId)
	if err != nil {
		log.Error("Could not start the Operation", err)
		return cacao.NewVariables(), err
	}

	// poll for operation status
	for finished, err := connection.IsOperationFinished(operationId); true; {
		if err != nil {
			log.Warn("Could not poll for operation status, retrying in 3 seconds")
			time.Sleep(3 * time.Second)
			continue
		}
		if finished {
			break
		}
		time.Sleep(3 * time.Second)
	}

	facts, err := processFacts(connection, operationId, globalFactSourceId, context)
	if err != nil {
		log.Error("Could not process facts", err)
		return cacao.NewVariables(), err
	}

	// remove any artifacts
	if createdAbility {
		cleanup(connection, abilityId)
	} else {
		cleanup(connection, "")
	}

	return facts, nil
}

// cleanup handles the removal of the created artifacts on the caldera server using the same connection.
func cleanup(cc ICalderaConnection, abilityId string) {
	if abilityId != "" {
		err := cc.DeleteAbility(abilityId)
		if err != nil {
			log.Warn("Could not cleanup artifacts from command")
		}
	}
}

func createGlobalFactSourceIfMissing(cc ICalderaConnection, metadata execution.Metadata) (string, error) {
	existingSources, err := cc.GetFactSources()
	expectedFactSourceName := fmt.Sprintf("soarca-%s-global", metadata.ExecutionId)

	if err != nil {
		return "", err
	}

	for _, source := range existingSources {
		if source.Name == expectedFactSourceName {
			return source.ID, nil
		}
	}

	id, err := cc.CreateFactSource(expectedFactSourceName)

	if err != nil {
		return "", err
	}

	return id, nil
}

func processFacts(cc ICalderaConnection, operationId string, globalFactSourceId string, context capability.Context) (cacao.Variables, error) {
	operationFacts, err := cc.RequestFacts(operationId)
	if err != nil {
		log.Error("Could not fetch Facts from Operation")
		return nil, err
	}

	existingFacts, err := cc.GetFactSource(globalFactSourceId)

	if err != nil {
		log.Error("Could not fetch facts from global source!")
		return nil, err
	}

	for _, f := range operationFacts {
		existingFacts.Facts = append(existingFacts.Facts, f.Facts...)
		existingFacts.Relationships = append(existingFacts.Relationships, f.Relationships...)
	}

	err = cc.SetFactSourceFacts(globalFactSourceId, existingFacts)

	if err != nil {
		return nil, err
	}
	variables := make(cacao.Variables)
	for _, link := range operationFacts {
		for _, fact := range link.Facts {
			factName := fmt.Sprint("__", fact.Name, "__")
			if context.Step.InArgs == nil || !slices.Contains(context.Step.InArgs, factName) {
				continue
			}
			variables[factName] = cacao.Variable{
				Name:  factName,
				Type:  cacao.VariableTypeString,
				Value: fmt.Sprint(fact.Value),
			}
		}
	}

	return variables, nil
}

// ParseJsonAbility converts a byte array into a caldera Ability.
// It tries to Unmarshal the byte array as a json struct and casts it to the wanted struct.
// If it fails to do that, or if the byte array is not a valid json struct, it logs an error and returns an empty caldera Ability struct.
func ParseJsonAbility(bytes []byte) *models.Ability {
	var ability models.Ability
	err := json.Unmarshal(bytes, &ability)
	if err != nil {
		log.Error("Could not convert ability JSON to a valid Ability object")
		return &models.Ability{}
	}
	return &ability
}

// ParseYamlAbility converts a byte array into a caldera Ability.
// It tries to Unmarshal the byte array as a yaml struct and casts it to the wanted struct.
// If it fails to do that, or if the byte array is not a valid yaml struct, it logs an error and returns an empty caldera Ability struct.
func ParseYamlAbility(bytes []byte) *models.Ability {
	var ability models.Ability
	err := yaml.Unmarshal(bytes, &ability)
	if err != nil {
		log.Error("Could not convert ability YAML to a valid Ability object")
		return &models.Ability{}
	}
	return &ability
}
