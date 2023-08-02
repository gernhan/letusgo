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

func NewFuture(data interface{}, err *error, wg *sync.WaitGroup) *Future {
	return &Future{
		dataPointer: data,
		err:         err,
		wg:          wg,
	}
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
	f.wg.Wait()
}

func (f *Future) Data() interface{} {
	if f.dataPointer == nil {
		return nil
	}
	pointer := f.dataPointer.(*interface{})
	return *pointer
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
