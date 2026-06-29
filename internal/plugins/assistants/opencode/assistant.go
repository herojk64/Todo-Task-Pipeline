package opencode

import (
	"fmt"
	"os"
	"os/exec"

	"taskpilot/internal/core"
)

type Assistant struct{}

func New() *Assistant {
	return &Assistant{}
}

func (a *Assistant) Run(task core.Task) error {
	fmt.Println("Opencode assistant called")
	fmt.Println("Task ID:", task.ID)

	prompt := BuildPrompt(task)

	cmd := exec.Command(
		"opencode",
		"--prompt",
		prompt,
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
