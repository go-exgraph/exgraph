package main

import (
	"context"
	"fmt"

	exgraph "github.com/go-exgraph/exgraph"
)

func main() {
	sub1 := exgraph.NewGraph().Pipeline(NewTask("1"), NewTask("2"))
	sub2 := exgraph.NewGraph().Spawns(NewTask("3"), NewTask("4"))
	d := exgraph.NewGraph()
	d.Spawns(sub1, sub2).Pipeline(NewTask("5"), NewTask("6"))
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
