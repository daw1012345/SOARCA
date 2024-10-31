package caldera

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"slices"
	"time"

	"soarca/internal/capability/caldera/caldera_connection"
	"soarca/logger"
	"soarca/models/cacao"
	"soarca/models/execution"
)

type calderaCapability struct{}

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

func New() *calderaCapability {
	return &calderaCapability{}
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
	connection, err := caldera_connection.New()
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
		abilityId, err = connection.CreateAbility(bytes)
		if err != nil {
			log.Error("Could not create custom Ability")
			return cacao.NewVariables(), err
		}
		createdAbility = true
	} else {
		abilityId = command.Command
	}

	// create the adversary
	adversaryId, err := connection.CreateAdversary(abilityId)
	if err != nil {
		log.Error("Could not create Adversary")
		return cacao.NewVariables(), err
	}

	// parse the specific target group
	if target.Type == "security-category" && slices.Contains(target.Category, "caldera") {
		groupName = target.Name
	}

	// start the operation
	operationId, err := connection.CreateOperation(groupName, adversaryId)
	if err != nil {
		log.Error("Could not start the Operation")
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
		cleanup(connection, adversaryId, abilityId)
	} else {
		cleanup(connection, adversaryId, "")
	}

	return parseFacts(facts), nil
}

func cleanup(cc *caldera_connection.CalderaConnection, adversaryId string, abilityId string) error {
	if abilityId != "" {
		err := cc.DeleteAbility(abilityId)
		if err != nil {
			log.Warn("Could not cleanup artifacts from command")
			return err
		}
	}
	return cc.DeleteAdversary(adversaryId)
}

func parseFacts(facts caldera_connection.CalderaFacts) cacao.Variables {
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
