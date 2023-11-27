package wait_group

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

const DefaultTimeout time.Duration = 60 * time.Second

type Job struct {
	handler       func(interface{}, chan interface{})
	inputToHandle interface{}
	// time out in ms
	timeout time.Duration
	result  interface{}
	worker  *Worker
}

func NewJob(handler func(interface{}, chan interface{}), inputToHandle interface{}, timeout time.Duration) *Job {
	if timeout <= 0 {
		timeout = DefaultTimeout
	}
	return &Job{
		handler:       handler,
		inputToHandle: inputToHandle,
		timeout:       timeout,
		result:        nil,
		worker:        nil,
	}
}

func (job *Job) ByWorker(worker *Worker) {
	job.worker = worker
}

func (job *Job) Handling() {
	go job.handler(job.inputToHandle, job.worker.resultChannel)
	fmt.Printf("\nHandling with timeout %v [worker %v]", job.timeout, job.worker.id)
	job.WaitForResult()
}

func (job *Job) WaitForResult() {
	select {
	case job.result = <-job.worker.resultChannel:
		fmt.Printf("\nReceived from resultChannel %v [worker %v]", job.result, job.worker.id)
		return
	case <-time.After(job.timeout):
		fmt.Printf("\nReached timeout after %d ms [worker %v]", job.timeout/time.Millisecond, job.worker.id)
		job.result = errors.New(fmt.Sprintf("Worker %v received timeout error (%v ms)", job.worker.id, job.timeout))
		return
	}
}
