package caldera

import (
	"soarca/pkg/core/capability"
	calderaModels "soarca/pkg/core/capability/caldera/api/models"
	"soarca/pkg/models/cacao"
	"soarca/pkg/models/execution"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
)

func TestCapabilityName(t *testing.T) {
	calderaCapability := New(&MockCalderaConnectionFactory{})
	assert.Equal(t, calderaCapability.GetType(), "soarca-caldera-cmd")
}

func TestExecute(t *testing.T) {
	calderaCapability := New(&MockCalderaConnectionFactory{})

	results, err := calderaCapability.Execute(
		execution.Metadata{},
		capability.Context{
			Command: cacao.Command{Type: "caldera-cmd", Command: "id: abilityID"},
		},
	)

	if err != nil {
		t.Fail()
	}

	retVal, exists := results.Find("__sample.operation.fact__")

	assert.Equal(t, true, exists)
	assert.Equal(t, "some_value_1", retVal.Value)
}

func TestExecuteB64(t *testing.T) {
	calderaCapability := New(&MockCalderaConnectionFactory{})
	execId := uuid.New()

	results, err := calderaCapability.Execute(
		execution.Metadata{
			ExecutionId: execId,
		},
		capability.Context{
			Command: cacao.Command{Type: "caldera-cmd", CommandB64: "e30="},
		},
	)

	if err != nil {
		t.Fail()
	}

	retVal, exists := results.Find("__sample.operation.fact__")

	assert.Equal(t, true, exists)
	assert.Equal(t, "some_value_1", retVal.Value)
}

func TestParseYamlAbility(t *testing.T) {
	resultingAbility := ParseYamlAbility([]byte("ability_id: 9a30740d-3aa8-4c23-8efa-d51215e8a5b9"))
	assert.Equal(t, resultingAbility.AbilityID, "9a30740d-3aa8-4c23-8efa-d51215e8a5b9")
}

func TestParseJsonAbility(t *testing.T) {
	resultingAbility := ParseJsonAbility([]byte("{\"ability_id\": \"9a30740d-3aa8-4c23-8efa-d51215e8a5b9\"}"))
	assert.Equal(t, resultingAbility.AbilityID, "9a30740d-3aa8-4c23-8efa-d51215e8a5b9")
}

func TestParseJsonAbilityWithException(t *testing.T) {
	// This should not crash, just produce an empty Ability
	resultingAbility := ParseJsonAbility([]byte("  / very %$#% invalid json"))
	assert.Equal(t, resultingAbility.AbilityID, "")
}

func TestParseYamlAbilityWithException(t *testing.T) {
	// This should not crash, just produce an empty Ability
	resultingAbility := ParseYamlAbility([]byte("  / very %$#% invalid yaml"))
	assert.Equal(t, resultingAbility.AbilityID, "")
}

func TestCleanup(t *testing.T) {
	capability := New(&MockCalderaConnectionFactory{})
	connection, err := capability.Factory.Create()
	assert.Equal(t, err, nil)
	cleanup(connection)
}

func TestExecuteErrorConnection(t *testing.T) {
	calderaCapability := New(&MockBadCalderaConnectionFactory{})
	_, err := calderaCapability.Execute(
		execution.Metadata{},
		capability.Context{
			Command: cacao.Command{Type: "caldera-cmd", CommandB64: "e30="},
		},
	)

	if err == nil {
		t.Fail()
	}

	assert.Equal(t, "Error Creating Ability", err.Error())
}

func TestGetCalderaInstance(t *testing.T) {
	_, err := GetCalderaInstance()
	assert.Equal(t, err, nil)
}

func TestCalderaHandleFacts(t *testing.T) {
	vars, err := processFacts(&MockCalderaConnection{}, "operation-0001", "soarca-0000-global")
	assert.Equal(t, vars["__sample.operation.fact__"].Name, "__sample.operation.fact__")
	assert.Equal(t, err, nil)

	_, err1 := processFacts(&MockBadCalderaConnection{}, "operation-0001", "soarca-0000-global")

	assert.NotEqual(t, err1, nil)

}

func TestBadCalderaConnection(t *testing.T) {
	capability := New(nil)
	connection, err := capability.Factory.Create()
	assert.Equal(t, err, nil)

	_, err1 := connection.CreateOperation("soarca-1234", "agentGroup-0001", "adversary-0001", "source-0001")
	assert.NotEqual(t, err1, nil)

	_, err2 := connection.CreateAdversary("soarca-1234", "ability-0001")
	assert.NotEqual(t, err2, nil)

	_, err3 := connection.IsOperationFinished("operation-0001")
	assert.NotEqual(t, err3, nil)

	_, err4 := connection.RequestFacts("operation-0001")
	assert.NotEqual(t, err4, nil)

	_, err5 := connection.CreateAbility(&calderaModels.Ability{})
	assert.NotEqual(t, err5, nil)

	err6 := connection.DeleteAbility("ability-0001")
	assert.NotEqual(t, err6, nil)

	_, err7 := connection.GetFactSource("source-0001")
	assert.NotEqual(t, err7, nil)

	err8 := connection.DeleteFactSource("source-0001")
	assert.NotEqual(t, err8, nil)

	_, err9 := connection.GetFactSources()
	assert.NotEqual(t, err9, nil)

	_, err10 := connection.CreateFactSource("source-0001")
	assert.NotEqual(t, err10, nil)

	err11 := connection.SetFactSourceFacts("source-0001", &calderaModels.PartialSource{})
	assert.NotEqual(t, err11, nil)

	err12 := connection.DeleteOperation("soarca-c045fe91-7dc6-4a2a-bbd6-b531ac183537")
	assert.NotEqual(t, err12, nil)

	err13 := connection.DeleteAdversary("soarca-c045fe91-7dc6-4a2a-bbd6-b531ac183537")
	assert.NotEqual(t, err13, nil)

	_, err14 := connection.GetAdversaries()
	assert.NotEqual(t, err14, nil)

	_, err15 := connection.GetAbilities()
	assert.NotEqual(t, err15, nil)

	_, err16 := connection.GetStaleOperations()
	assert.NotEqual(t, err16, nil)
}
