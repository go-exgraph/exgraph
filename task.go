package exgraph

type Task interface {
	Run(Context) error
}
