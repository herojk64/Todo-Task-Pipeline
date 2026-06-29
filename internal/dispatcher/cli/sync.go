package cli

import (
	"fmt"

	"taskpilot/internal/core"
	"taskpilot/internal/plugins/fake"
	"taskpilot/internal/plugins/jira"
	"taskpilot/internal/plugins/taskwarrior"
	"taskpilot/internal/rules"
)

func RunSync(cfg core.Config) error {
	var provider core.Provider

	switch cfg.Provider {
	case "jira":
		provider = jira.NewProvider(cfg)
	case "fake":
		provider = fake.NewProvider()

	default:
		return fmt.Errorf(
			"unknown provider: %s",
			cfg.Provider,
		)
	}

	sink := taskwarrior.NewSink()
	policy := rules.NewPolicy()

	engine := core.NewEngine(
		provider,
		sink,
		policy,
	)

	return engine.Sync()
}
