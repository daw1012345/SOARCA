package caldera

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"slices"
	"time"

	"soarca/internal/capability/caldera/api/models"
	"soarca/logger"
	"soarca/models/cacao"
	"soarca/models/execution"
)

type calderaCapability struct {
	factory ICalderaConnectionFactory
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

func New(factory ICalderaConnectionFactory) *calderaCapability {
	if factory == nil {
		factory = calderaConnectionFactory{}
	}
	return &calderaCapability{factory}
}

func (capability *calderaCapability) GetType() string {
	return capabilityName
}

func (capability *calderaCapability) Execute(
	metadata execution.Metadata,
	command cacao.Command,
	authentication cacao.AuthenticationInformation,
	target cacao.AgentTarget,
	variables cacao.Variables) (cacao.Variables, error) {

	log.Info("Successfully called execute on the caldera capability")
	connection, err := capability.factory.Create()
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
		ability := parseYamlAbility(bytes)
		abilityId, err = connection.CreateAbility(ability)
		if err != nil {
			log.Error("Could not create custom Ability")
			return cacao.NewVariables(), err
		}
		createdAbility = true
	} else {
		abilityId = command.Command
	}

	// parse the specific target group
	if target.Type == "security-category" && slices.Contains(target.Category, "caldera") {
		groupName = target.Name
	}

	// start the operation
	operationId, err := connection.CreateOperation(groupName, abilityId)
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
	facts, err := connection.RequestFacts(operationId)
	if err != nil {
		log.Error("Could not fetch Facts from Operation")
		return cacao.NewVariables(), err
	}

	// remove any artifacts
	if createdAbility {
		cleanup(connection, abilityId)
	} else {
		cleanup(connection, "")
	}

	return parseFacts(facts), nil
}

func cleanup(cc ICalderaConnection, abilityId string) {
	if abilityId != "" {
		err := cc.DeleteAbility(abilityId)
		if err != nil {
			log.Warn("Could not cleanup artifacts from command")
		}
	}
}

func parseFacts(facts CalderaFacts) cacao.Variables {
	variables := make(cacao.Variables, len(facts))
	for name, value := range facts {
		variables[name] = cacao.Variable{
			Name:  fmt.Sprintf("__%s__", name),
			Type:  cacao.VariableTypeString,
			Value: value,
		}
	}
	return variables
}

func parseYamlAbility(bytes []byte) *models.Ability {
	return &models.Ability{
		AbilityID: string(bytes),
	}
}
