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

// Create implements caldera.ICalderaConnection.
func (f MockCalderaConnectionFactory) Create() (ICalderaConnection, error) {
	return &MockCalderaConnection{}, nil
}

// CreateAbility implements caldera.ICalderaConnection.
func (m MockCalderaConnection) CreateAbility(ability *models.Ability) (string, error) {
	return "abilityID", nil
}

// CreateAdversary implements caldera.ICalderaConnection.
func (m MockCalderaConnection) CreateAdversary(name string, abilityId string) (string, error) {
	return "adversaryID", nil
}

// CreateOperation implements caldera.ICalderaConnection.
func (m MockCalderaConnection) CreateOperation(
	name string,
	agentGroupId string,
	abilityId string,
	factSourceId string,
) (string, error) {
	return "operationID", nil
}

// DeleteAbility implements caldera.ICalderaConnection.
func (m MockCalderaConnection) DeleteAbility(abilityId string) error {
	return nil
}

// IsOperationFinished implements caldera.ICalderaConnection.
func (m MockCalderaConnection) IsOperationFinished(operationId string) (bool, error) {
	return true, nil
}

// RequestFacts implements caldera.ICalderaConnection.
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

// GetStaleOperations implements caldera.ICalderaConnection.
func (m MockCalderaConnection) GetStaleOperations() ([]*models.PartialOperation, error) {
	return []*models.PartialOperation{
		&models.PartialOperation{
			Name: "soarca-6b2f53c8-4b08-47de-9c5c-aeacfe0ee723",
		},
		&models.PartialOperation{
			Name: "d4774b76-f78c-4e09-bf2b-a0715a8d3015",
		},
		&models.PartialOperation{
			Name: "Some random, unrelated operation",
		},
	}, nil
}

// DeleteOperation implements caldera.ICalderaConnection.
func (m MockCalderaConnection) DeleteOperation(operationId string) error {
	return nil
}

// DeleteAdversary implements caldera.ICalderaConnection.
func (m MockCalderaConnection) DeleteAdversary(adversaryId string) error {
	return nil
}

// GetAdversaries implements caldera.ICalderaConnection.
func (m MockCalderaConnection) GetAdversaries() ([]*models.PartialAdversary, error) {
	return []*models.PartialAdversary{
		&models.PartialAdversary{
			Name: "soarca-6b2f53c8-4b08-47de-9c5c-aeacfe0ee723",
		},
	}, nil
}

// GetAbilities implements caldera.ICalderaConnection.
func (m MockCalderaConnection) GetAbilities() ([]*models.PartialAbility, error) {
	name := "soarca-6b2f53c8-4b08-47de-9c5c-aeacfe0ee723"
	return []*models.PartialAbility{
		&models.PartialAbility{
			Name: &name,
		},
	}, nil
}

type MockBadCalderaConnectionFactory struct{}
type MockBadCalderaConnection struct{}

// Create implements caldera.ICalderaConnection.
func (f MockBadCalderaConnectionFactory) Create() (ICalderaConnection, error) {
	return &MockBadCalderaConnection{}, nil
}

// CreateAbility implements caldera.ICalderaConnection.
func (m MockBadCalderaConnection) CreateAbility(ability *models.Ability) (string, error) {
	return "", errors.New("Error Creating Ability")
}

// CreateOperation implements caldera.ICalderaConnection.
func (m MockBadCalderaConnection) CreateOperation(
	name string,
	agentGroupId string,
	abilityId string,
	factSourceId string,
) (string, error) {
	return "", errors.New("Error Creating Operation")
}

// DeleteAbility implements caldera.ICalderaConnection.
func (m MockBadCalderaConnection) DeleteAbility(abilityId string) error {
	return errors.New("Error Deleting Ability")
}

// IsOperationFinished implements caldera.ICalderaConnection.
func (m MockBadCalderaConnection) IsOperationFinished(operationId string) (bool, error) {
	return false, errors.New("Error Fetching Finished")
}

// RequestFacts implements caldera.ICalderaConnection.
func (m MockBadCalderaConnection) RequestFacts(operationId string) ([]*models.PartialLink, error) {
	return make([]*models.PartialLink, 0), errors.New("Error Requesting Facts")
}

// CreateAdversary implements caldera.ICalderaConnection.
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

// GetStaleOperations implements caldera.ICalderaConnection.
func (m MockBadCalderaConnection) GetStaleOperations() ([]*models.PartialOperation, error) {
	return nil, errors.New("Unable to retrieve or determine stale operations")
}

// DeleteOperation implements caldera.ICalderaConnection.
func (m MockBadCalderaConnection) DeleteOperation(operationId string) error {
	return errors.New("Unable to delete the Operation")
}

// DeleteAdversary implements caldera.ICalderaConnection.
func (m MockBadCalderaConnection) DeleteAdversary(adversaryId string) error {
	return errors.New("Unable to delete the Adversary")
}

// GetAdversaries implements caldera.ICalderaConnection.
func (m MockBadCalderaConnection) GetAdversaries() ([]*models.PartialAdversary, error) {
	return nil, errors.New("Unable to fetch all Adversaries")
}

// GetAbilities implements caldera.ICalderaConnection.
func (m MockBadCalderaConnection) GetAbilities() ([]*models.PartialAbility, error) {
	return nil, errors.New("Unable to fetch all Abilities")
}
