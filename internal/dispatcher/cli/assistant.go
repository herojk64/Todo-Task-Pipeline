package cli

import (
	"fmt"

	"taskpilot/internal/core"
	"taskpilot/internal/plugins/assistants/claude"
	"taskpilot/internal/plugins/assistants/ollama"
	"taskpilot/internal/plugins/assistants/openai"
)

func RunAssistant(name string, task core.Task) error {
	var assistant core.Assistant

	switch name {
	case "claude":
		assistant = claude.New()

	case "ollama":
		assistant = ollama.New()

	case "openai":
		assistant = openai.New()

	default:
		return fmt.Errorf(
			"unknown assistant: %s",
			name,
		)
	}

	return assistant.Run(task)
}
