package concurrent

import (
	"sync"
)

func SupplyAsync(handler func() (interface{}, error)) *Future {
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

func SupplyAsyncWithPool(handler func() (interface{}, error), pool *ThreadPool) *Future {
	return pool.Submit(handler)
}

func RunAsync(handler func() error) *Future {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var err error

	go func() {
		defer wg.Done()
		err = handler()
	}()
	return NewFuture(nil, &err, &wg)
}

func RunAsyncWithPool(handler func() error, pool *ThreadPool) *Future {
	return pool.SubmitAndReturnErrorOnly(handler)
}
