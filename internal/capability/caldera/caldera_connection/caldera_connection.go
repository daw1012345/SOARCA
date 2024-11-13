package caldera_connection

import (
	"fmt"
	"soarca/internal/capability/caldera/api/client/abilities"
	"soarca/internal/capability/caldera/api/client/operationsops"
	"soarca/internal/capability/caldera/api/models"
)

type connectionState int

const (
	uninitiated connectionState = iota
	building
	built
	runningOperation
	finishedOperation
	receivedFacts
	cleared
	done
)

type CalderaConnection struct {
	instance *calderaInstance
	state    connectionState
}
type CalderaFacts map[string]string

func New() (*CalderaConnection, error) {
	var instance, err = getCalderaInstance()
	if err != nil {
		return nil, err
	}
	return &CalderaConnection{instance, uninitiated}, nil
}

func (cc CalderaConnection) CreateAbility(ability *models.Ability) (string, error) {
	response, err := cc.instance.send.Abilities.PostAPIV2Abilities(
		abilities.NewPostAPIV2AbilitiesParams().WithBody(ability),
	)
	if err != nil {
		return "", err
	}
	return response.GetPayload().AbilityID, nil
}
func (cc CalderaConnection) DeleteAbility(abilityId string) error {
	_, err := cc.instance.send.Abilities.DeleteAPIV2AbilitiesAbilityID(
		abilities.NewDeleteAPIV2AbilitiesAbilityIDParams().WithAbilityID(abilityId),
	)
	return err
}
func (cc CalderaConnection) CreateOperation(
	agentGroupId string,
	abilityId string,
) (string, error) {
	response, err := cc.instance.send.Operationsops.PostAPIV2Operations(
		operationsops.NewPostAPIV2OperationsParams().WithBody(&models.Operation{
			Adversary: &models.Adversary{
				AtomicOrdering: []string{abilityId},
			},
			Group:      agentGroupId,
			Autonomous: 1,
		}),
	)
	if err != nil {
		return "", err
	}
	return response.GetPayload().ID, nil
}
func (cc CalderaConnection) IsOperationFinished(operationId string) (bool, error) {
	response, err := cc.instance.send.Operationsops.GetAPIV2OperationsID(
		operationsops.NewGetAPIV2OperationsIDParams().WithID(operationId),
	)
	if err != nil {
		return false, err
	}
	if response.GetPayload().State == "finished" {
		return true, nil
	}
	return false, nil
}
func (cc CalderaConnection) RequestFacts(operationId string) (CalderaFacts, error) {
	response, err := cc.instance.send.Operationsops.GetAPIV2OperationsIDLinks(
		operationsops.NewGetAPIV2OperationsIDLinksParams().WithID(operationId),
	)
	if err != nil {
		return nil, err
	}
	var facts = make(CalderaFacts)
	for _, link := range response.GetPayload() {
		for _, fact := range link.Facts {
			facts[fmt.Sprint(link.Paw, "-", fact.Name)] = fmt.Sprint(fact.Value)
		}
	}
	return facts, nil
}
