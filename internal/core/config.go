package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Provider string

	JiraBase  string
	JiraEmail string
	JiraToken string

	JiraJQL string
}

func LoadConfig() Config {
	if execPath, err := os.Executable(); err == nil {
		envPath := filepath.Join(filepath.Dir(execPath), ".env")
		if godotenv.Load(envPath) != nil {
			if err := godotenv.Load(); err != nil {
				fmt.Println("warning: .env not found, using system environment")
			}
		}
	} else {
		if err := godotenv.Load(); err != nil {
			fmt.Println("warning: .env not found, using system environment")
		}
	}

	return Config{
		Provider: os.Getenv("TASKPILOT_PROVIDER"),

		JiraBase:  os.Getenv("JIRA_BASE"),
		JiraEmail: os.Getenv("JIRA_EMAIL"),
		JiraToken: os.Getenv("JIRA_TOKEN"),

		JiraJQL: os.Getenv("JIRA_JQL"),
	}
}
