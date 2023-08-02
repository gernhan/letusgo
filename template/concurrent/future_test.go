package concurrent

import (
	"log"
	"testing"
)

func TestSupplyAsyncWithPool(t *testing.T) {
	tp := NewThreadPool(3)
	nThreads := 1001
	mtp := EmptyMultiFutures()
	for i := 0; i < nThreads; i++ {
		finalI := i
		mtp.AddFuture(SupplyAsyncWithPool(func() (interface{}, error) {
			log.Printf("i: %v", finalI)
			return finalI, nil
		}, tp).thenSupplyAsyncWithPool(func(input interface{}) (interface{}, error) {
			value := input.(int)
			log.Printf("log again i: %v", value)
			return nil, nil
		}, tp))
	}
	mtp.Wait()
}

func TestRunAsyncWithPool(t *testing.T) {
	tp := NewThreadPool(3)
	nThreads := 1001
	mtp := EmptyMultiFutures()
	for i := 0; i < nThreads; i++ {
		finalI := i
		mtp.AddFuture(RunAsyncWithPool(func() error {
			log.Printf("i: %v", finalI)
			return nil
		}, tp).thenRunAsyncWithPool(func() error {
			log.Printf("log again i: %v", finalI)
			return nil
		}, tp))
	}
	mtp.Wait()
}
