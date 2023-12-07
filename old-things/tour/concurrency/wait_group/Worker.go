package wait_group

import (
	"github.com/google/uuid"
	"sync"
	"time"
)

type Worker struct {
	id                               uuid.UUID
	commitment                       *sync.WaitGroup
	job                              *Job
	willReceiveResultFromOtherThread bool
	resultChannel                    chan interface{}
}

func NewWorker(
	id uuid.UUID,
	handler func(interface{}, chan interface{}),
	inputToHandle interface{},
	timeout time.Duration, // time out in ms
) *Worker {
	worker := Worker{
		id:            id,
		job:           NewJob(handler, inputToHandle, timeout),
		resultChannel: make(chan interface{}),
	}
	worker.job.ByWorker(&worker)
	return &worker
}

func (worker *Worker) WithCommitment(commitment *sync.WaitGroup) {
	worker.commitment = commitment
	worker.commitment.Add(1)
}

func (worker *Worker) Work() {
	worker.job.Handling()
	worker.commitment.Done()
}

func (worker *Worker) CompleteTheJobWithResultFromOtherThread(result interface{}) {
	worker.resultChannel <- result
}

type WorkerService interface {
	Work() (interface{}, error)
}
