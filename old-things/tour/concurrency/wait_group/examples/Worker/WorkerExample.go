package main

import (
	"github.com/google/uuid"
	"go_tutorial/tour/concurrency/wait_group"
	"sync"
	"time"
)

func main() {
	var commitment sync.WaitGroup
	worker1 := wait_group.NewWorker(
		uuid.New(),
		func(input interface{}, output chan interface{}) {
			time.Sleep(4 * time.Second)
			output <- "Done"
		},
		nil,
		0,
	)
	worker1.WithCommitment(&commitment)

	worker2 := wait_group.NewWorker(
		uuid.New(),
		func(input interface{}, output chan interface{}) {
			time.Sleep(5 * time.Second)
			output <- "Done"
		},
		nil,
		0,
	)
	worker2.WithCommitment(&commitment)
	go worker1.Work()
	go worker2.Work()

	commitment.Wait()
}
