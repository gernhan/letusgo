package main

import (
	"github.com/google/uuid"
	"go_tutorial/tour/concurrency/wait_group"
	"time"
)

func main() {
	workGroup := wait_group.WorkerGroup{}
	workGroup.Work(
		wait_group.NewWorker(
			uuid.New(),
			func(input interface{}, output chan interface{}) {
				time.Sleep(4 * time.Second)
				output <- "Done"
			},
			nil,
			0,
		),
		wait_group.NewWorker(
			uuid.New(),
			func(input interface{}, output chan interface{}) {
				time.Sleep(5 * time.Second)
				output <- "Done"
			},
			nil,
			0,
		),
	)
}
