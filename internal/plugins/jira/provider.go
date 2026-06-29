package jira

import (
	"encoding/json"

	"taskpilot/internal/core"
)

type Provider struct {
	client *Client
	cfg    core.Config
}

func NewProvider(cfg core.Config) *Provider {
	return &Provider{
		client: NewClient(
			cfg.JiraBase,
			cfg.JiraEmail,
			cfg.JiraToken,
		),
		cfg: cfg,
	}
}

func (p *Provider) FetchTasks() ([]core.Task, error) {
	issues, err := p.client.FetchIssues(p.cfg.JiraJQL)
	if err != nil {
		return nil, err
	}

	tasks := make([]core.Task, 0, len(issues))

	for _, i := range issues {
		descBytes, _ := json.Marshal(i.Fields.Description)
		tasks = append(tasks, core.Task{
			ID:          i.Key,
			Title:       i.Fields.Summary,
			Status:      i.Fields.Status.Name,
			Project:     i.Fields.Project.Key,
			Description: string(descBytes),
		})
	}
	return tasks, nil
}
