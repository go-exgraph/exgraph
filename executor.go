package exgraph

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Executor interface {
	Exec(ctx context.Context, dag Graph) error
}

type executor struct {
}

func NewExecutor() Executor {
	return &executor{}
}

func (e *executor) Exec(ctx context.Context, dag Graph) error {
	c := newContext(ctx)
	for _, v := range dag.(*graph).jobs {
		e.execJob(c, v)
	}
	return nil
}

func (e *executor) execJob(c *contextImpl, job *Job) (err error) {
	var taskfuncs []TaskFunc
	for _, v := range job.tasks {
		taskfuncs = append(taskfuncs, v.Run)
	}

	if job.sequential {
		return e.execTasksSync(c, taskfuncs)
	}
	return e.execTasksAsync(c, taskfuncs)
}

func (e *executor) execTasksSync(c *contextImpl, taskfuncs []TaskFunc) (err error) {
	for _, v := range taskfuncs {
		if err := v(c); err != nil {
			return err
		}
	}
	return nil
}

func (e *executor) execTasksAsync(c *contextImpl, taskfuncs []TaskFunc) (err error) {
	g, _ := errgroup.WithContext(c.ctx)
	for i := 0; i < len(taskfuncs); i++ {
		f := taskfuncs[i]
		g.Go(func() error {
			return f(c)
		})
	}
	return g.Wait()
}
