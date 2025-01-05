package caldera

import (
	"errors"
	"soarca/pkg/core/capability/caldera/api/models"
)

type MockCalderaConnectionFactory struct{}
type MockCalderaConnection struct{}

// CreateFactSource implements caldera.ICalderaConnection.
func (m *MockCalderaConnection) CreateFactSource(factSourceName string) (string, error) {
	return "factSource", nil
}

// DeleteFactSource implements caldera.ICalderaConnection.
func (m *MockCalderaConnection) DeleteFactSource(factSourceId string) error {
	return nil
}

// GetFactSource implements caldera.ICalderaConnection.
func (m *MockCalderaConnection) GetFactSource(factSourceId string) (*models.PartialSource, error) {
	return &models.PartialSource{
		Facts: []*models.Fact{
			&models.Fact{
				Name:  "sample.fact.from.source",
				Value: "some_value",
			},
		},
	}, nil
}

// GetFactSources implements caldera.ICalderaConnection.
func (m *MockCalderaConnection) GetFactSources() ([]*models.PartialSource, error) {
	return make([]*models.PartialSource, 0), nil
}

// SetFactSourceFacts implements caldera.ICalderaConnection.
func (m *MockCalderaConnection) SetFactSourceFacts(factSourceName string, factSourceData *models.PartialSource) error {
	return nil
}

func (f MockCalderaConnectionFactory) Create() (ICalderaConnection, error) {
	return &MockCalderaConnection{}, nil
}

func (m MockCalderaConnection) CreateAbility(ability *models.Ability) (string, error) {
	return "abilityID", nil
}

func (m MockCalderaConnection) CreateAdversary(name string, abilityId string) (string, error) {
	return "adversaryID", nil
}

func (m MockCalderaConnection) CreateOperation(
	name string,
	agentGroupId string,
	abilityId string,
	factSourceId string,
) (string, error) {
	return "operationID", nil
}

func (m MockCalderaConnection) DeleteAbility(abilityId string) error {
	return nil
}

func (m MockCalderaConnection) IsOperationFinished(operationId string) (bool, error) {
	return true, nil
}

func (m MockCalderaConnection) RequestFacts(operationId string) ([]*models.PartialLink, error) {
	return []*models.PartialLink{
		&models.PartialLink{
			Facts: []*models.Fact{
				&models.Fact{
					Name:  "sample.operation.fact",
					Value: "some_value_1",
				},
			},
		},
	}, nil
}

type MockBadCalderaConnectionFactory struct{}
type MockBadCalderaConnection struct{}

func (f MockBadCalderaConnectionFactory) Create() (ICalderaConnection, error) {
	return &MockBadCalderaConnection{}, nil
}

func (m MockBadCalderaConnection) CreateAbility(ability *models.Ability) (string, error) {
	return "", errors.New("Error Creating Ability")
}

func (m MockBadCalderaConnection) CreateOperation(
	name string,
	agentGroupId string,
	abilityId string,
	factSourceId string,
) (string, error) {
	return "", errors.New("Error Creating Operation")
}

func (m MockBadCalderaConnection) DeleteAbility(abilityId string) error {
	return errors.New("Error Deleting Ability")
}

func (m MockBadCalderaConnection) IsOperationFinished(operationId string) (bool, error) {
	return false, errors.New("Error Fetching Finished")
}

func (m MockBadCalderaConnection) RequestFacts(operationId string) ([]*models.PartialLink, error) {
	return make([]*models.PartialLink, 0), errors.New("Error Requesting Facts")
}

func (m MockBadCalderaConnection) CreateAdversary(name string, abilityId string) (string, error) {
	return "", errors.New("Error Creating Adversary")
}

// CreateFactSource implements caldera.ICalderaConnection.
func (m *MockBadCalderaConnection) CreateFactSource(factSourceName string) (string, error) {
	return "", errors.New("Can't Create Source")
}

// DeleteFactSource implements caldera.ICalderaConnection.
func (m *MockBadCalderaConnection) DeleteFactSource(factSourceId string) error {
	return errors.New("Cannot Delete Source")
}

// GetFactSource implements caldera.ICalderaConnection.
func (m *MockBadCalderaConnection) GetFactSource(factSourceId string) (*models.PartialSource, error) {
	return nil, errors.New("Cannot Get Source")
}

// GetFactSources implements caldera.ICalderaConnection.
func (m *MockBadCalderaConnection) GetFactSources() ([]*models.PartialSource, error) {
	return nil, errors.New("Cannot Get Sources")
}

// SetFactSourceFacts implements caldera.ICalderaConnection.
func (m *MockBadCalderaConnection) SetFactSourceFacts(factSourceName string, factSourceData *models.PartialSource) error {
	return errors.New("Cannot Set Facts")
}
