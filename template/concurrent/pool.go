package concurrent

import (
	"sync"
)

type ThreadPool struct {
	semaphore *Semaphore
}

func NewThreadPool(maxWorkers int) *ThreadPool {
	tp := &ThreadPool{
		semaphore: NewSemaphore(maxWorkers),
	}
	return tp
}

func (tp *ThreadPool) Submit(handler func() (interface{}, error)) *Future {
	tp.semaphore.Acquire()
	defer tp.semaphore.Release()
	wg := sync.WaitGroup{}
	wg.Add(1)
	var data interface{}
	var err error

	go func() {
		defer wg.Done()
		data, err = handler()
	}()
	return NewFuture(&data, &err, &wg)
}

func (tp *ThreadPool) SubmitAndReturnErrorOnly(handler func() error) *Future {
	tp.semaphore.Acquire()
	defer tp.semaphore.Release()
	wg := sync.WaitGroup{}
	wg.Add(1)
	var err error

	go func() {
		defer wg.Done()
		err = handler()
	}()
	return NewFuture(nil, &err, &wg)
}
