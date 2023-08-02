package concurrent

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewSemaphore(t *testing.T) {
	// Create a semaphore with a maximum of 3 concurrent permits.
	startTime := time.Now()
	sem := NewSemaphore(3)

	// Create a wait group to synchronize goroutines.
	var wg sync.WaitGroup
	count := int64(0)

	// Start multiple goroutines.
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Acquire the permit from the semaphore.
			sem.Acquire()
			atomic.AddInt64(&count, 1)
			defer sem.Release()

			// Simulate some work.
			fmt.Printf("Goroutine %d started\n", i)
			fmt.Printf("number of concurrent threads: %v\n", count)
			// In a real scenario, you would perform some actual work here.
			time.Sleep(time.Second)

			atomic.AddInt64(&count, -1)
			fmt.Printf("Goroutine %d finished\n", i)
		}(i)
	}

	// Wait for all goroutines to finish.
	wg.Wait()

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Time consumed: %v\n", executionTime)
}
