package main

import (
	"context"
	"fmt"

	exgraph "github.com/go-exgraph/exgraph"
)

func main() {
	d := exgraph.NewGraph()
	d.Pipeline(NewTask("hello"), NewTask("world")).
		Spawns(NewTask("foo"), NewTask("bar"))
	exgraph.NewExecutor().Exec(context.Background(), d)
}

type Task struct {
	Name string
}

func NewTask(name string) *Task {
	return &Task{Name: name}
}

func (task *Task) Run(ctx exgraph.Context) (err error) {
	fmt.Println(task.Name)
	return nil
}
