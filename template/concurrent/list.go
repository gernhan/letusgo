package concurrent

import (
	"fmt"
	"sync"
)

type List struct {
	data []interface{}
	mu   *sync.Mutex
}

func NewConcurrentList() *List {
	return &List{
		data: make([]interface{}, 0),
		mu:   &sync.Mutex{},
	}
}

func (cl *List) Add(element interface{}) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.data = append(cl.data, element)
}

func (cl *List) Get(index int) (interface{}, error) {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	if index < 0 || index >= len(cl.data) {
		return 0, fmt.Errorf("index out of range")
	}

	return cl.data[index], nil
}

func (cl *List) Size() int {
	return len(cl.data)
}

func (cl *List) Values() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		cl.mu.Lock()
		defer cl.mu.Unlock()
		for _, value := range cl.data {
			ch <- value
		}
	}()
	return ch
}
