package concurrent

import (
	"sync"
	"testing"
)

func TestConcurrentList(t *testing.T) {
	list, listSize := getConcurrentList()

	if list.Size() != listSize {
		t.Errorf("Expected list size %d, actual: %v", listSize, list.Size())
	}
	for i := 0; i < listSize; i++ {
		_, err := list.Get(i)
		if err != nil {
			t.Errorf("Error getting value at index %d: %v", i, err)
		}
	}
}

func getConcurrentList() (*List, int) {
	list := NewConcurrentList()

	var wg sync.WaitGroup
	listSize := 1002
	for i := 0; i < listSize; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			list.Add(i)
		}(i)
	}
	wg.Wait()
	return list, listSize
}
