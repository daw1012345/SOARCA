package executors

import (
	"soarca/pkg/models/cacao"
	"soarca/pkg/models/execution"
)

type IPlaybookExecutor interface {
	Execute(execution.Metadata,
		cacao.Step,
		cacao.Variables) (cacao.Variables, error)
}
