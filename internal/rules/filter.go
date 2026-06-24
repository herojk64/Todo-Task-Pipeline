package rules

import "jira-task-sync/internal/core"

type Policy struct{}

func NewPolicy() *Policy {
	return &Policy{}
}

func (p *Policy) Apply(tasks []core.Task) []core.Task {
	out := make([]core.Task, 0, len(tasks))

	for _, t := range tasks {
		if t.Title == "" {
			continue
		}

		out = append(out, t)
	}

	return out
}
