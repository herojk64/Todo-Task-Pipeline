package core

type Engine struct {
	provider Provider
	sink     Sink
	policy   Policy
}

func NewEngine(p Provider, s Sink, pol Policy) *Engine {
	return &Engine{
		provider: p,
		sink:     s,
		policy:   pol,
	}
}

func (e *Engine) Sync() error {
	tasks, err := e.provider.FetchTasks()
	if err != nil {
		return err
	}

	filtered := e.policy.Apply(tasks)

	return e.sink.Write(filtered)
}
