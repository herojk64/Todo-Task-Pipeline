package jira

import (
	"os"

	"jira-task-sync/internal/core"
)

type Provider struct {
	client *Client
}

func NewProvider() *Provider {
	return &Provider{
		client: NewClient(
			os.Getenv("JIRA_BASE"),
			os.Getenv("JIRA_EMAIL"),
			os.Getenv("JIRA_TOKEN"),
		),
	}
}

func (p *Provider) FetchTasks() ([]core.Task, error) {
	issues, err := p.client.FetchIssues()
	if err != nil {
		return nil, err
	}

	tasks := make([]core.Task, 0, len(issues))

	for _, i := range issues {
		tasks = append(tasks, core.Task{
			ID:     i.Key,
			Title:  i.Fields.Summary,
			Status: i.Fields.Status.Name,
		})
	}

	return tasks, nil
}
