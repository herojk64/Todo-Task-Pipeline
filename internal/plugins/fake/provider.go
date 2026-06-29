package fake

import "taskpilot/internal/core"

type Provider struct{}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) FetchTasks() ([]core.Task, error) {
	return []core.Task{
		{ID: "1", Title: "Test Task A", Status: "Todo"},
		{ID: "2", Title: "Test Task B", Status: "InProgress"},
	}, nil
}
