package taskwarrior

import (
	"fmt"
	"os/exec"

	"taskpilot/internal/core"
)

type Sink struct{}

func NewSink() *Sink {
	return &Sink{}
}

func (s *Sink) createTask(t core.Task) error {
	cmd := exec.Command(
		"task",
		"add",
		fmt.Sprintf("project:%s", t.Project),
		"+jira",
		t.Title,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("taskwarrior: %w: %s", err, string(output))
	}

	return nil
}

func (s *Sink) annotateTask(t core.Task) error {
	cmd := exec.Command(
		"task",
		"+LATEST",
		"annotate",
		fmt.Sprintf("jira:%s", t.ID),
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("taskwarrior annotate: %w: %s", err, string(output))
	}

	return nil
}

func (s *Sink) Write(tasks []core.Task) error {
	for _, t := range tasks {
		if err := s.createTask(t); err != nil {
			return err
		}
		if err := s.annotateTask(t); err != nil {
			return err
		}
	}
	return nil
}
