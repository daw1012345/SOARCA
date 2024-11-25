package caldera

import (
	"soarca/internal/capability/caldera"
	"soarca/models/cacao"
	"soarca/models/execution"
	mock_caldera "soarca/test/unittest/mocks/mock_utils/caldera"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCapabilityName(t *testing.T) {
	calderaCapability := caldera.New(&mock_caldera.MockCalderaConnectionFactory{})
	assert.Equal(t, calderaCapability.GetType(), "soarca-caldera-cmd")
}

func TestExecute(t *testing.T) {
	calderaCapability := caldera.New(&mock_caldera.MockCalderaConnectionFactory{})

	results, err := calderaCapability.Execute(
		execution.Metadata{},
		cacao.Command{Type: "caldera-cmd", Command: "id: abilityID"},
		cacao.AuthenticationInformation{},
		cacao.AgentTarget{},
		cacao.NewVariables())

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, cacao.NewVariables(), results)
}

func TestExecuteB64(t *testing.T) {
	calderaCapability := caldera.New(&mock_caldera.MockCalderaConnectionFactory{})

	results, err := calderaCapability.Execute(
		execution.Metadata{},
		cacao.Command{Type: "caldera-cmd", CommandB64: "e30="},
		cacao.AuthenticationInformation{},
		cacao.AgentTarget{},
		cacao.NewVariables())

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, cacao.NewVariables(), results)
}

func TestExecuteErrorConnection(t *testing.T) {
	calderaCapability := caldera.New(&mock_caldera.MockBadCalderaConnectionFactory{})

	_, err := calderaCapability.Execute(
		execution.Metadata{},
		cacao.Command{Type: "caldera-cmd", CommandB64: "e30="},
		cacao.AuthenticationInformation{},
		cacao.AgentTarget{},
		cacao.NewVariables())

	if err == nil {
		t.Fail()
	}

	assert.Equal(t, "Error Creating Ability", err.Error())
}
