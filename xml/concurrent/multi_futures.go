package concurrent

import "sync"

type MultiFutures struct {
	FList []*Future
}

func EmptyMultiFutures() *MultiFutures {
	return &MultiFutures{
		FList: make([]*Future, 0),
	}
}

func AllOf(fList []*Future) *MultiFutures {
	return &MultiFutures{
		FList: fList,
	}
}

func (f *MultiFutures) AddFuture(future *Future) {
	f.FList = append(f.FList, future)
}

func (f *MultiFutures) Wait() {
	for _, future := range f.FList {
		future.Wait()
	}
}

func (f *MultiFutures) Result() AllResults {
	result := AllResults{
		Successes: make([]Result, 0),
		Failures:  make([]Result, 0),
		IsSucceed: true,
	}

	f.Wait()
	id := 0
	for _, future := range f.FList {
		id++
		if future.Error() != nil {
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

func (f *MultiFutures) ToFuture() *Future {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var err error
	var data interface{}

	go func() {
		defer wg.Done()
		result := AllResults{
			Successes: make([]Result, 0),
			Failures:  make([]Result, 0),
			IsSucceed: true,
		}
		data = &result

		f.Wait()
		id := 0
		for _, future := range f.FList {
			id++
			if future.Error() != nil {
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
	}()
	return NewFuture(&data, &err, &wg)
}
