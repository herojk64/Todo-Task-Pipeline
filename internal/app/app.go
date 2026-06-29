package app

import (
	"taskpilot/internal/core"
	"taskpilot/internal/dispatcher/cli"
)

type App struct {
	cfg core.Config
}

func New(cfg core.Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (a *App) Run() error {
	return cli.Dispatch(a.cfg)
}
