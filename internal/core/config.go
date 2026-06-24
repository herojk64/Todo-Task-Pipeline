package core

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Provider string

	JiraBase  string
	JiraEmail string
	JiraToken string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("warning: .env not found, using system environment")
	}

	return Config{
		Provider: "jira",

		JiraBase:  os.Getenv("JIRA_BASE"),
		JiraEmail: os.Getenv("JIRA_EMAIL"),
		JiraToken: os.Getenv("JIRA_TOKEN"),
	}
}
