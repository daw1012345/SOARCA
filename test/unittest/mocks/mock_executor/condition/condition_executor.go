package mock_condition_executor

import (
	"soarca/pkg/models/cacao"
	"soarca/pkg/models/execution"

	"github.com/stretchr/testify/mock"
)

type Mock_Condition struct {
	mock.Mock
}

func (executor *Mock_Condition) Execute(metadata execution.Metadata,
	step cacao.Step,
	variables cacao.Variables) (string, bool, error) {
	args := executor.Called(metadata, step, variables)
	return args.String(0), args.Bool(1), args.Error(2)
}
