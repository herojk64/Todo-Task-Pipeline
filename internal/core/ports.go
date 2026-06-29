package core

type Provider interface {
	FetchTasks() ([]Task, error)
}

type Sink interface {
	Write([]Task) error
}

type Policy interface {
	Apply([]Task) []Task
}

type Assistant interface {
	Run(task Task) error
}

type TaskRepository interface {
	Get(id string) (*Task, error)
}
