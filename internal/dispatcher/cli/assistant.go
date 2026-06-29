package cli

import (
	"fmt"

	"taskpilot/internal/core"
	"taskpilot/internal/plugins/assistants/claude"
	"taskpilot/internal/plugins/assistants/ollama"
	"taskpilot/internal/plugins/assistants/opencode"
)

func RunAssistant(name string, task core.Task, cfg core.Config) error {
	var assistant core.Assistant

	switch name {
	case "claude":
		assistant = claude.New()

	case "ollama":
		assistant = ollama.New(cfg.OllamaModel)

	case "opencode":
		assistant = opencode.New()

	default:
		return fmt.Errorf(
			"unknown assistant: %s",
			name,
		)
	}

	return assistant.Run(task)
}
