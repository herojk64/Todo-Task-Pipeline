package ollama

import (
	"fmt"
	"os"
	"os/exec"

	"taskpilot/internal/core"
)

type Assistant struct {
	model string
}

func New(model string) *Assistant {
	if model == "" {
		model = "llama3"
	}

	return &Assistant{model: model}
}

func (a *Assistant) Run(task core.Task) error {
	fmt.Println("Ollama assistant called (model: %s) \n", a.model)
	fmt.Println("Task ID:", task.ID)

	prompt := BuildPrompt(task)

	cmd := exec.Command("ollama", "run", a.model, prompt)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
