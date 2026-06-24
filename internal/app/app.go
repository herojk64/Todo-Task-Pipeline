package app

import (
	"fmt"

	"jira-task-sync/internal/core"
	"jira-task-sync/internal/plugins/fake"
	"jira-task-sync/internal/plugins/jira"
	"jira-task-sync/internal/plugins/taskwarrior"
	"jira-task-sync/internal/rules"
)

type App struct {
	engine *core.Engine
}

func New(cfg core.Config) *App {
	var provider core.Provider

	switch cfg.Provider {
	case "jira":
		provider = jira.NewProvider()

	case "fake":
		provider = fake.NewProvider()

	default:
		panic(fmt.Sprintf("unknown provider: %s", cfg.Provider))
	}

	sink := taskwarrior.NewSink()
	policy := rules.NewPolicy()

	engine := core.NewEngine(provider, sink, policy)

	return &App{engine: engine}
}

func (a *App) Run() error {
	return a.engine.Sync()
}
