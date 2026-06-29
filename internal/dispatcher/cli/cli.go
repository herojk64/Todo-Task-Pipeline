package cli

import (
	"os"

	"taskpilot/internal/core"
)

func Dispatch(cfg core.Config) error {
	args := os.Args
	if len(args) < 2 {
		Usage()
		os.Exit(1)
	}

	switch args[1] {
	case "sync":
		return RunSync(cfg)

	case "work":
		return RunWork(cfg)
	}

	return nil
}
