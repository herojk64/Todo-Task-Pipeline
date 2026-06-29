package cli

import "fmt"

func Usage() {
	fmt.Println("TaskPilot")
	fmt.Println("")

	fmt.Println("Usage:")
	fmt.Println("\ttaskpilot sync <TODO-MODEL>")
	fmt.Println("\ttaskpilot work <ASSISTANT> <TASK-ID>")
	fmt.Println("")

	fmt.Println("TODO-MODELS:")
	fmt.Println("\tjira")
	fmt.Println("")

	fmt.Println("ASSISTANTS:")
	fmt.Println("\tclaude")
	fmt.Println("\topenai")
	fmt.Println("\tollama")
	fmt.Println("")

	fmt.Println("Examples:")
	fmt.Println("\ttaskpilot sync jira")
	fmt.Println("\ttaskpilot work claude JAT-1")
	fmt.Println("\ttaskpilot work openai DEMO-4")
	fmt.Println("\ttaskpilot work ollama JAT-1")
	fmt.Println("")

	fmt.Println("Workflow:")
	fmt.Println("\t1. Sync tasks from the configured provider")
	fmt.Println("\t2. Select a task to work on")
	fmt.Println("\t3. Launch an AI assistant with task context")
	fmt.Println("\t4. Review and approve changes inside the assistant")
}
