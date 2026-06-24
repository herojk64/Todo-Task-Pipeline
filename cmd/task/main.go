package main

import (
	"log"

	"jira-task-sync/internal/app"
	"jira-task-sync/internal/core"
)

func main() {
	cfg := core.LoadConfig()

	application := app.New(cfg)

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
