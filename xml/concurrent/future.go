package concurrent

import (
	"sync"
)

type Future struct {
	dataPointer interface{}
	err         *error
	wg          *sync.WaitGroup
	Id          int
	Name        string
	readyToWait *sync.WaitGroup
}

type AllResults struct {
	Successes []Result
	Failures  []Result
	IsSucceed bool
}

type Result struct {
	Id   int
	Name string
	Data interface{}
	Err  error
}

func NewEmptyFuture() *Future {
	var dataPointer interface{}
	var err error
	readyToWait := &sync.WaitGroup{}
	readyToWait.Add(1)
	return &Future{
		dataPointer: dataPointer,
		err:         &err,
		wg:          nil,
		readyToWait: readyToWait,
	}
}

func NewFuture(data interface{}, err *error, wg *sync.WaitGroup) *Future {
	future := Future{
		dataPointer: data,
		err:         err,
		wg:          wg,
		readyToWait: &sync.WaitGroup{},
	}
	future.readyToWait.Add(1)
	future.ReadyToWait()
	return &future
}

func NewFutureWithId(data *interface{}, err *error, wg *sync.WaitGroup, name string, id int) *Future {
	return &Future{
		dataPointer: data,
		err:         err,
		wg:          wg,
		Id:          id,
		Name:        name,
	}
}

func (f *Future) Wait() {
	f.readyToWait.Wait()
	if f.wg != nil {
		f.wg.Wait()
	}
}

func (f *Future) ReadyToWait() {
	f.readyToWait.Done()
}

func (f *Future) Data() interface{} {
	if f.dataPointer == nil {
		return nil
	}
	pointer := f.dataPointer.(*interface{})
	return *pointer
}

func (f *Future) SetDataPointer(pointer interface{}) {
	f.dataPointer = pointer
}

func (f *Future) SetErrorPointer(pointer *error) {
	f.err = pointer
}

func (f *Future) Error() error {
	if f.err == nil {
		return nil
	}
	return *f.err
}

func WaitAllOf(futures []Future) AllResults {
	result := AllResults{
		Successes: make([]Result, 0),
		Failures:  make([]Result, 0),
		IsSucceed: true,
	}

	for _, future := range futures {
		future.Wait()

		if future.err != nil {
			result.IsSucceed = false
			result.Failures = append(result.Failures, Result{
				Id:   future.Id,
				Name: future.Name,
				Data: future.Data(),
				Err:  future.Error(),
			})
		} else {
			result.Successes = append(result.Successes, Result{
				Id:   future.Id,
				Name: future.Name,
				Data: future.Data(),
				Err:  future.Error(),
			})
		}
	}
	return result
}

func WaitAllOfList(futures List) AllResults {
	result := AllResults{
		Successes: make([]Result, 0),
		Failures:  make([]Result, 0),
		IsSucceed: true,
	}

	id := 0
	for item := range futures.Values() {
		id++
		future := item.(Future)
		future.Wait()

		if future.err != nil {
			result.IsSucceed = false
			result.Failures = append(result.Failures, Result{
				Id:   id,
				Name: future.Name,
				Data: future.Data(),
				Err:  future.Error(),
			})
		} else {
			result.Successes = append(result.Successes, Result{
				Id:   id,
				Name: future.Name,
				Data: future.Data(),
				Err:  future.Error(),
			})
		}
	}
	return result
}

func (f *Future) thenSupplyAsync(handler func(interface{}) (interface{}, error)) *Future {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var data interface{}
	var err error

	go func() {
		defer wg.Done()
		f.wg.Wait()
		if f.Error() == nil {
			data, err = handler(f.Data())
		} else {
			data, err = nil, f.Error()
		}
	}()
	return NewFuture(&data, &err, &wg)
}

func (f *Future) thenSupplyAsyncWithPool(handler func(interface{}) (interface{}, error), pool *ThreadPool) *Future {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var data interface{}
	var err error
	pool.semaphore.Acquire()
	defer pool.semaphore.Release()

	go func() {
		defer wg.Done()
		f.wg.Wait()
		if f.Error() == nil {
			data, err = handler(f.Data())
		} else {
			data, err = nil, f.Error()
		}
	}()
	return NewFuture(&data, &err, &wg)
}

func (f *Future) thenRunAsync(handler func() error) *Future {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var err error

	go func() {
		defer wg.Done()
		f.wg.Wait()
		if f.Error() == nil {
			err = handler()
		} else {
			err = f.Error()
		}
	}()
	return NewFuture(nil, &err, &wg)
}

func (f *Future) thenRunAsyncWithPool(handler func() error, pool *ThreadPool) *Future {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var err error
	pool.semaphore.Acquire()
	defer pool.semaphore.Release()

	go func() {
		defer wg.Done()
		f.wg.Wait()
		if f.Error() == nil {
			err = handler()
		} else {
			err = f.Error()
		}
	}()
	return NewFuture(nil, &err, &wg)
}
