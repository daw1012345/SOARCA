// This file bundles the abstract connection request methods used by the main business logic of the Caldera capability.
package caldera

import (
	"strings"
	"soarca/pkg/core/capability/caldera/api/client/abilities"
	"soarca/pkg/core/capability/caldera/api/client/adversaries"
	"soarca/pkg/core/capability/caldera/api/client/operationsops"
	"soarca/pkg/core/capability/caldera/api/client/sources"
	"soarca/pkg/core/capability/caldera/api/models"
)

// ICalderaConnectionFactory is a factory struct that defines a Create method to construct a specific type of ICalderaConnection.
type ICalderaConnectionFactory interface {
	Create() (ICalderaConnection, error)
}

// calderaConnectionFactory is the default ICalderaConnectionFactory and builds a calderaConnection.
type calderaConnectionFactory struct{}

// Create builds a calderaConnection.
// It first tries to get the Caldera server instance, and returns a calderaConnection struct connected to the instance.
//
// If it fails to retrieve the Caldera instance, it returns nil with the error.
func (c calderaConnectionFactory) Create() (ICalderaConnection, error) {
	var instance, err = GetCalderaInstance()
	if err != nil {
		return nil, err
	}
	return &calderaConnection{instance}, nil
}

// ICalderaConnection is the inferface that defines the higher level abstract connection requests.
//
// These methods are used for the main business logic of the Caldera capability.
// Each method creates a request to the Caldera instance and returns either the parsed response or an error if something went wrong.
type ICalderaConnection interface {
	CreateAbility(ability *models.Ability) (string, error)
	DeleteAbility(abilityId string) error
	CreateOperation(name string, agentGroupId string, adversaryId string, factSourceId string) (string, error)
	IsOperationFinished(operationId string) (bool, error)
	RequestFacts(operationId string) ([]*models.PartialLink, error)
	CreateAdversary(name string, abilityId string) (string, error)
	SetFactSourceFacts(factSourceName string, factSourceData *models.PartialSource) error
	GetFactSources() ([]*models.PartialSource, error)
	GetFactSource(factSourceId string) (*models.PartialSource, error)
	CreateFactSource(factSourceName string) (string, error)
	DeleteFactSource(factSourceId string) error
	GetStaleOperations() ([]*models.PartialOperation, error)
	DeleteOperation(operationId string) error
	GetAdversaries() ([]*models.PartialAdversary, error)
	DeleteAdversary(adversaryId string) error
	GetAbilities() ([]*models.PartialAbility, error)
}

// calderaConnection is the default Caldera connection struct built by the calderaConnectionFactory.
// It contains a pointer to the Caldera instance it is connected to.
//
// This connection is used for the actual connection requests to a Caldera instance.
type calderaConnection struct {
	instance *calderaInstance
}

// CalderaFacts is a map between the fact name as string to the value, also as a string.
//
// Plans are to be able to have different types of values, like integers or arrays.
type CalderaFacts map[string]string

// CreateAbility initiates a request to the Caldera instance to create a Caldera Ability.
//
// It returns the id of the created ability, or an error if it fails.
func (cc calderaConnection) CreateAbility(ability *models.Ability) (string, error) {
	response, err := cc.instance.send.Abilities.PostAPIV2Abilities(
		abilities.NewPostAPIV2AbilitiesParams().WithBody(ability),
		authenticateCaldera,
	)
	if err != nil {
		return "", err
	}
	return response.GetPayload().AbilityID, nil
}

// DeleteAbility initiates a request to the Caldera instance to delete a Caldera Ability.
//
// It returns an error if it fails.
func (cc calderaConnection) DeleteAbility(abilityId string) error {
	_, err := cc.instance.send.Abilities.DeleteAPIV2AbilitiesAbilityID(
		abilities.NewDeleteAPIV2AbilitiesAbilityIDParams().WithAbilityID(abilityId),
		authenticateCaldera,
	)
	return err
}

// CreateOperation initiates a request to the Caldera instance to create a Caldera Operation.
//
// It returns the id of the created operation, or an error if it fails.
func (cc calderaConnection) CreateOperation(
	name string,
	agentGroupId string,
	adversaryId string,
	factSourceId string,
) (string, error) {
	response, err := cc.instance.send.Operationsops.PostAPIV2Operations(
		operationsops.NewPostAPIV2OperationsParams().WithBody(&models.Operation{
			Adversary: &models.Adversary{
				AdversaryID:    adversaryId,
				AtomicOrdering: []string{},
				Tags:           []string{},
			},
			Group:      agentGroupId,
			Autonomous: 1,
			AutoClose:  true,
			Name:       &name,
			Source: &models.Source{
				ID:            factSourceId,
				Adjustments:   make([]*models.Adjustment, 0),
				Rules:         make([]*models.Rule, 0),
				Facts:         make([]*models.Fact, 0),
				Relationships: make([]*models.Relationship, 0),
			},
		}),
		authenticateCaldera,
	)
	if err != nil {
		return "", err
	}
	return response.GetPayload().ID, nil
}

// CreateAdversary initiates a request to the Caldera instance to create a Caldera Adversary.
//
// It returns the id of the created adversary, or an error if it fails.
func (cc calderaConnection) CreateAdversary(name string, abilityId string) (string, error) {
	response, err := cc.instance.send.Adversaries.PostAPIV2Adversaries(
		adversaries.NewPostAPIV2AdversariesParams().WithBody(&models.Adversary{
			Name:           name,
			AtomicOrdering: []string{abilityId},
			Tags:           []string{},
		}),
		authenticateCaldera,
	)
	if err != nil {
		return "", err
	}
	return response.GetPayload().AdversaryID, nil
}

// IsOperationFinished initiates a request to the Caldera instance to check if a given Caldera Operation is finished.
//
// It returns true if the operation is finished, false if it is not, or an error if it fails.
func (cc calderaConnection) IsOperationFinished(operationId string) (bool, error) {
	response, err := cc.instance.send.Operationsops.GetAPIV2OperationsID(
		operationsops.NewGetAPIV2OperationsIDParams().WithID(operationId),
		authenticateCaldera,
	)
	if err != nil {
		log.Error(err)
		return false, err
	}
	if response.GetPayload().State == "finished" {
		return true, nil
	}
	return false, nil
}

// RequestFacts initiates a request to the Caldera instance to get the facts of a given Caldera Operation.
//
// It returns a list of Caldera links, which contain the Caldera facts, or an error if it fails.
func (cc calderaConnection) RequestFacts(operationId string) ([]*models.PartialLink, error) {
	response, err := cc.instance.send.Operationsops.GetAPIV2OperationsIDLinks(
		operationsops.NewGetAPIV2OperationsIDLinksParams().WithID(operationId),
		authenticateCaldera,
	)
	if err != nil {
		return nil, err
	}
	return response.GetPayload(), nil
}

// SetFactSourceFacts initiates a request to the Caldera instance to update a given Caldera Fact Source with new facts and relationships.
//
// Only return a possible error if the request fails.
func (cc calderaConnection) SetFactSourceFacts(factSourceName string, factSourceData *models.PartialSource) error {
	_, err := cc.instance.send.Sources.PatchAPIV2SourcesID(sources.NewPatchAPIV2SourcesIDParams().WithID(factSourceName).WithBody(factSourceData), authenticateCaldera)

	return err
}

// CreateFactSource initiates a request to the Caldera instance to create a new Caldera Fact Source.
//
// It returns the id of the created fact source, or an error if it fails.
func (cc calderaConnection) CreateFactSource(factSourceName string) (string, error) {
	body := models.Source{}
	body.Name = factSourceName
	body.Facts = []*models.Fact{}
	body.Relationships = []*models.Relationship{}
	body.Rules = []*models.Rule{}

	response, err := cc.instance.send.Sources.PostAPIV2Sources(sources.NewPostAPIV2SourcesParams().WithBody(&body), authenticateCaldera)

	if err != nil {
		return "", err
	}

	return response.GetPayload().ID, nil
}

// GetFactSources initiates a request to the Caldera instance to get all Caldera Fact Sources.
//
// It returns a list of Caldera Fact Sources, or an error if it fails.
func (cc calderaConnection) GetFactSources() ([]*models.PartialSource, error) {
	response, err := cc.instance.send.Sources.GetAPIV2Sources(sources.NewGetAPIV2SourcesParams(), authenticateCaldera)

	if err != nil {
		return nil, err
	}

	return response.GetPayload(), nil
}

// DeleteFactSource initiates a request to the Caldera instance to delete a Caldera Fact Source.
//
// It only returns an error if the request fails.
func (cc calderaConnection) DeleteFactSource(factSourceId string) error {
	_, err := cc.instance.send.Sources.DeleteAPIV2SourcesID(sources.NewDeleteAPIV2SourcesIDParams().WithID(factSourceId), authenticateCaldera)
	if strings.Contains(err.Error(), "(status 204)") {
		// The HTTP unfortunately treats HTTP 204 responses (No Content) as
		// invalid/error responses. We want to ignore those errors.
		return nil
	}
	return err
}

// GetFactSource initiates a request to the Caldera instance to get a specific Caldera Fact Source.
//
// It returns the Caldera Fact Source, or an error if it fails.
func (cc calderaConnection) GetFactSource(factSourceId string) (*models.PartialSource, error) {
	response, err := cc.instance.send.Sources.GetAPIV2SourcesID(sources.NewGetAPIV2SourcesIDParams().WithID(factSourceId), authenticateCaldera)

	if err != nil {
		return nil, err
	}

	return response.GetPayload(), nil
}

// GetStaleOperations provides a slice of all Operations that are - based on their name -
// created by SOARCA and are finished. These operations should in principle be deleted.
// Unfortunately, Caldera does not allow this type of filtering to take place server-side.
func (cc calderaConnection) GetStaleOperations() ([]*models.PartialOperation, error) {
	response, err := cc.instance.send.Operationsops.GetAPIV2Operations(
		operationsops.NewGetAPIV2OperationsParams(),
		authenticateCaldera,
	)
	if err != nil {
		return nil, err
	}
	fullResult := response.GetPayload()
	filteredResult := make([]*models.PartialOperation, 0)
	for i := 0; i < len(fullResult); i++ {
		if !strings.HasPrefix(fullResult[i].Name, "soarca-") {
			// This Operation is not created by SOARCA, ignore
			continue
		}
		if fullResult[i].State == "finished" {
			filteredResult = append(filteredResult, fullResult[i])
		}
	}
	return filteredResult, nil
}

// Deletes an Operation by its ID. Return an error upon failure.
func (cc calderaConnection) DeleteOperation(operationId string) error {
	_, err := cc.instance.send.Operationsops.DeleteAPIV2OperationsID(
		operationsops.NewDeleteAPIV2OperationsIDParams().WithID(operationId),
		authenticateCaldera,
	)
	if strings.Contains(err.Error(), "(status 204)") {
		// The HTTP unfortunately treats HTTP 204 responses (No Content) as
		// invalid/error responses. We want to ignore those errors.
		return nil
	}
	return err
}

// Deletes an Adversary by its ID. Return an error upon failure.
func (cc calderaConnection) DeleteAdversary(adversaryId string) error {
	_, err := cc.instance.send.Adversaries.DeleteAPIV2AdversariesAdversaryID(
		adversaries.NewDeleteAPIV2AdversariesAdversaryIDParams().WithAdversaryID(adversaryId),
		authenticateCaldera,
	)
	return err
}

// Retrieves all available Abilities
func (cc calderaConnection) GetAbilities() ([]*models.PartialAbility, error) {
	response, err := cc.instance.send.Abilities.GetAPIV2Abilities(abilities.NewGetAPIV2AbilitiesParams(), authenticateCaldera)
	if err != nil {
		return nil, err
	}
	return response.GetPayload(), nil
}

// Retrieves all available Adversaries
func (cc calderaConnection) GetAdversaries() ([]*models.PartialAdversary, error) {
	response, err := cc.instance.send.Adversaries.GetAPIV2Adversaries(adversaries.NewGetAPIV2AdversariesParams(), authenticateCaldera)
	if err != nil {
		return nil, err
	}
	return response.GetPayload(), nil
}