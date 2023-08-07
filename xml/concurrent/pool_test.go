package concurrent

import (
	"fmt"
	"testing"
)

func TestNewThreadPool(t *testing.T) {
	// Create a thread pool with a maximum of 3 concurrent workers.
	pool := NewThreadPool(3)

	mtp := EmptyMultiFutures()
	for i := 1; i <= 10; i++ {
		taskID := i
		mtp.AddFuture(pool.Submit(func() (interface{}, error) {
			fmt.Printf("Task %d executed\n", taskID)
			return nil, nil
		}))
	}
	mtp.Wait()
}
