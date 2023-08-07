package atomics

import (
	"github.com/gernhan/xml/concurrent"
	"github.com/gernhan/xml/tools"
	"log"
	"testing"
)

func TestAtomicInteger(t *testing.T) {
	c := 0
	aList := NewAtomicReference(&c)

	nThreads := 101
	futures := concurrent.EmptyMultiFutures()
	for i := 0; i < nThreads; i++ {
		futures.AddFuture(concurrent.RunAsync(func() error {
			aList.Mutate(func(obj interface{}) {
				a := obj.(*int)
				*a = *a + 1
			})
			return nil
		}).thenRunAsync(func() error {
			log.Printf("")
			return nil
		}))
	}

	result := futures.Result()

	if !result.IsSucceed {
		err := result.Failures[0].Err
		t.Errorf("Caught error %v", err)
		return
	}

	log.Printf("value: %v", c)
	if c != nThreads {
		t.Errorf("Expected %v but got %v", nThreads, c)
	}
}

func TestAtomicList(t *testing.T) {
	c := make([]string, 0)
	aList := NewAtomicReference(&c)
	ig := tools.NewInputGenerator()

	nThreads := 101
	futures := concurrent.EmptyMultiFutures()
	for i := 0; i < nThreads; i++ {
		futures.AddFuture(concurrent.RunAsync(func() error {
			aList.Mutate(func(obj interface{}) {
				list := obj.(*[]string)
				*list = append(*list, ig.GenerateString(4))
			})
			return nil
		}))
	}
	result := futures.Result()

	if !result.IsSucceed {
		err := result.Failures[0].Err
		t.Errorf("Caught error %v", err)
	}

	ls := len(c)
	log.Printf("list size: %v", ls)
	if ls != nThreads {
		t.Errorf("Expected list size %v but got %v", nThreads, ls)
	}
	for i, data := range c {
		log.Printf("i : %v, dataPointer : %v", i, data)
	}
}

func TestWithoutAtomicList(t *testing.T) {
	nThreads := 102
	ls := getListSizeWithoutUsingAtomic(t, nThreads)
	ls2 := getListSizeWithoutUsingAtomic(t, nThreads)
	ls3 := getListSizeWithoutUsingAtomic(t, nThreads)
	if ls*3 == ls+ls2+ls3 && ls+ls2+ls3 == nThreads*3 && ls2 == ls3 {
		t.Error("Without atomic reference, list is still thread-safe")
	}
	log.Printf("list size: %v, %v, %v", ls, ls2, ls3)
}

func getListSizeWithoutUsingAtomic(t *testing.T, nThreads int) int {
	c := make([]string, 0)
	ig := tools.NewInputGenerator()

	futures := concurrent.EmptyMultiFutures()
	for i := 0; i < nThreads; i++ {
		futures.AddFuture(concurrent.RunAsync(func() error {
			c = append(c, ig.GenerateString(4))
			return nil
		}))
	}
	result := futures.Result()

	if !result.IsSucceed {
		err := result.Failures[0].Err
		t.Errorf("Caught error %v", err)
	}

	return len(c)
}
