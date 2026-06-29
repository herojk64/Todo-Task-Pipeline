package ollama

import (
	"taskpilot/internal/core"
)

type Assistant struct{}

func New() *Assistant {
	return &Assistant{}
}

func (a *Assistant) Run(task core.Task) error {
	return nil
}
