package mock_caldera

import (
	"soarca/internal/capability/caldera"
	"soarca/internal/capability/caldera/api/models"
)

type MockCalderaConnectionFactory struct{}
type MockCalderaConnection struct{}

func (f MockCalderaConnectionFactory) Create() (caldera.ICalderaConnection, error) {
	return &MockCalderaConnection{}, nil
}

func (m MockCalderaConnection) CreateAbility(ability *models.Ability) (string, error) {
	return "abilityID", nil
}

func (m MockCalderaConnection) CreateOperation(
	agentGroupId string,
	abilityId string,
) (string, error) {
	return "operationID", nil
}

func (m MockCalderaConnection) DeleteAbility(abilityId string) error {
	return nil
}

func (m MockCalderaConnection) IsOperationFinished(operationId string) (bool, error) {
	return true, nil
}

func (m MockCalderaConnection) RequestFacts(operationId string) (caldera.CalderaFacts, error) {
	return caldera.CalderaFacts{}, nil
}
