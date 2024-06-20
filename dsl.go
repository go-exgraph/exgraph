package exgraph

type pipelineDSL struct {
	graph Graph
}

func (dsl *pipelineDSL) Spawns(tasks ...Task) *spawnsDSL {
	dsl.graph.Spawns(tasks...)
	return &spawnsDSL{
		dsl.graph,
	}
}

func (dsl *pipelineDSL) Run(ctx Context) error {
	return dsl.graph.Run(ctx)
}

type spawnsDSL struct {
	graph Graph
}

func (dsl *spawnsDSL) Pipeline(tasks ...Task) *pipelineDSL {
	dsl.graph.Pipeline(tasks...)
	return &pipelineDSL{
		dsl.graph,
	}
}

func (dsl *spawnsDSL) Run(ctx Context) error {
	return dsl.graph.Run(ctx)
}
