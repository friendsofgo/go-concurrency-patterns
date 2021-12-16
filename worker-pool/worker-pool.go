package main

import (
	"fmt"
	"time"
)

type task interface {
	Run()
}

type simpleTask struct {
	id int
}

func (t simpleTask) Run() {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("<task %d> finished\n", t.id)
}

func main() {
	tasks := make(chan task)

	poolSize := 10
	for w := 1; w <= poolSize; w++ {
		w := w
		go func() {
			for t := range tasks {
				fmt.Printf("<worker %d> doing task\n", w)
				t.Run()
				fmt.Printf("<worker %d> finished task\n", w)
			}
		}()
	}

	numTasks := 100
	for t := 1; t <= numTasks; t++ {
		tasks <- simpleTask{id: t}
	}

	close(tasks)
}
