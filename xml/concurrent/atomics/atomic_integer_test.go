package atomics

import (
	"github.com/gernhan/xml/concurrent"
	"log"
	"strconv"
	"testing"
)

func TestAtomicInteger_AddAndGet(t *testing.T) {
	number := NewAtomicInteger()

	nThreads := 101
	futures := concurrent.EmptyMultiFutures()
	for i := 0; i < nThreads; i++ {
		futures.AddFuture(concurrent.RunAsync(func() error {
			log.Printf(strconv.FormatInt(number.AddAndGet(1), 10))
			return nil
		}).thenRunAsync(func() error {
			return nil
		}))
	}

	result := futures.Result()

	if !result.IsSucceed {
		err := result.Failures[0].Err
		t.Errorf("Caught error %v", err)
		return
	}

	log.Printf("value: %v", number)
	if number.Get() != int64(nThreads) {
		t.Errorf("Expected %v but got %v", nThreads, number)
	}
}

func TestAtomicInteger_IncrementAndGet(t *testing.T) {
	number := NewAtomicInteger()

	nThreads := 101
	futures := concurrent.EmptyMultiFutures()
	for i := 0; i < nThreads; i++ {
		futures.AddFuture(concurrent.RunAsync(func() error {
			log.Printf(strconv.FormatInt(number.IncrementAndGet(), 10))
			return nil
		}).thenRunAsync(func() error {
			return nil
		}))
	}

	result := futures.Result()

	if !result.IsSucceed {
		err := result.Failures[0].Err
		t.Errorf("Caught error %v", err)
		return
	}

	log.Printf("value: %v", number)
	if number.Get() != int64(nThreads) {
		t.Errorf("Expected %v but got %v", nThreads, number)
	}
}

func TestAtomicInteger_DecrementAndGet(t *testing.T) {
	nThreads := 100001
	number := NewAtomicIntegerWithInitial(int64(nThreads))

	futures := concurrent.EmptyMultiFutures()
	for i := 0; i < nThreads; i++ {
		futures.AddFuture(concurrent.RunAsync(func() error {
			number.DecrementAndGet()
			return nil
		}).thenRunAsync(func() error {
			return nil
		}))
	}

	result := futures.Result()

	if !result.IsSucceed {
		err := result.Failures[0].Err
		t.Errorf("Caught error %v", err)
		return
	}

	log.Printf("value: %v", number)
	if number.Get() != 0 {
		t.Errorf("Expected %v but got %v", 0, number)
	}
}
