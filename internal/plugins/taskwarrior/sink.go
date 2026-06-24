package taskwarrior

import (
	"fmt"

	"jira-task-sync/internal/core"
)

type Sink struct{}

func NewSink() *Sink {
	return &Sink{}
}

func (s *Sink) Write(tasks []core.Task) error {
	for _, t := range tasks {
		fmt.Printf("%s | %s | %s\n", t.ID, t.Title, t.Status)
	}
	return nil
}
