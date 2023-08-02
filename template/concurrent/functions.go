package concurrent

import "sync"

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
