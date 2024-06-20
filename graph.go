package exgraph

import "context"

// Graph represents directed acyclic graph
type Graph interface {
	Pipeline(tasks ...Task) *pipelineDSL
	Spawns(tasks ...Task) *spawnsDSL
	Exec(ctx context.Context) error
	Task
}

type graph struct {
	jobs []*Job
}

// New creates new Graph
func NewGraph() Graph {
	return &graph{
		jobs: make([]*Job, 0),
	}
}

func (g *graph) Run(ctx Context) error {
	return g.Exec(ctx.Context())
}

func (g *graph) Exec(ctx context.Context) error {
	return NewExecutor().Exec(ctx, g)
}

// Pipeline executes tasks sequentially
func (g *graph) Pipeline(tasks ...Task) *pipelineDSL {

	job := &Job{
		tasks:      tasks,
		sequential: true,
	}

	g.jobs = append(g.jobs, job)

	return &pipelineDSL{
		g,
	}
}

// Spawns executes tasks concurrently
func (g *graph) Spawns(tasks ...Task) *spawnsDSL {

	job := &Job{
		tasks:      tasks,
		sequential: false,
	}

	g.jobs = append(g.jobs, job)

	return &spawnsDSL{
		g,
	}
}
