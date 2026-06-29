package main

import (
	"log"

	"taskpilot/internal/app"
	"taskpilot/internal/core"
)

func main() {
	cfg := core.LoadConfig()

	application := app.New(cfg)

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
