package claude

import (
	"fmt"

	"taskpilot/internal/core"
)

func BuildPrompt(task core.Task) string {
	return fmt.Sprintf(`
	Task ID: %s
	Project: %s
	Title: %s

	You are assisting with implementation planning.

	DO NOT make code changes immediately.
	Analyze the task.
	Produce a step-by-step implementation plan.
	Wait for user approval before editing files.
	`,
		task.ID,
		task.Project,
		task.Title,
	)
}
