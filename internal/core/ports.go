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
