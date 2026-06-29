package taskwarrior

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"taskpilot/internal/core"
)

type taskwarriorTask struct {
	Description string `json:"description"`
	Project     string `json:"project"`
	Status      string `json:"status"`

	Annotations []struct {
		Description string `json:"description"`
	} `json:"annotations"`

	Tags []string `json:"tags"`
}

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Get(id string) (*core.Task, error) {
	cmd := exec.Command("task", "+jira", "export")

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var tasks []taskwarriorTask

	if err := json.Unmarshal(output, &tasks); err != nil {
		return nil, err
	}

	jiraRef := "jira:" + id

	for _, t := range tasks {
		if !hasAnnotation(t, jiraRef) {
			continue
		}

		return &core.Task{
			ID:      id,
			Title:   t.Description,
			Project: t.Project,
			Status:  t.Status,
		}, nil
	}

	return nil, fmt.Errorf(
		"task not found: %s",
		id,
	)
}

func hasAnnotation(task taskwarriorTask, target string) bool {
	for _, a := range task.Annotations {
		if a.Description == target {
			return true
		}
	}
	return false
}
