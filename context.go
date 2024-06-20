package exgraph

import "context"

type Context interface {
	Context() context.Context
}

type TaskFunc func(Context) error

type contextImpl struct {
	ctx context.Context
}

func (c *contextImpl) Context() context.Context {
	return c.ctx
}

func newContext(ctx context.Context) *contextImpl {
	return &contextImpl{ctx: ctx}
}
