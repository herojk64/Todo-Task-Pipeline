package cli

import (
	"fmt"
	"os"

	"taskpilot/internal/core"
	"taskpilot/internal/plugins/taskwarrior"
)

func RunWork(cfg core.Config) error {
	args := os.Args

	if len(args) < 4 {
		return fmt.Errorf(
			"usage: taskpilot work <assistant> <task-id>",
		)
	}

	assistantName := args[2]
	taskID := args[3]

	repo := taskwarrior.NewRepository()

	task, err := repo.Get(taskID)
	if err != nil {
		return err
	}

	return RunAssistant(
		assistantName,
		*task,
		cfg,
	)
}
