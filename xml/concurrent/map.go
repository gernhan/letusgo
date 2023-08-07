package concurrent

import (
	"github.com/gernhan/xml/utils"
	"sync"
)

type Map struct {
	data map[interface{}]interface{}
	mu   *sync.Mutex
}

func NewConcurrentMap() *Map {
	return &Map{
		data: make(map[interface{}]interface{}),
		mu:   &sync.Mutex{},
	}
}

func (cl *Map) Put(key interface{}, element interface{}) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.data[key] = element
}

func (cl *Map) Get(key interface{}) interface{} {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	return cl.data[key]
}

func (cl *Map) Contain(key interface{}) bool {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	return cl.data[key] != nil
}

func (cl *Map) Size() int {
	return len(cl.data)
}

func (cl *Map) String() string {
	json, err := utils.ParseJSON(cl.data)
	if err != nil {
		return ""
	}
	return json
}

func (cl *Map) Values() <-chan interface{} {
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

func (cl *Map) Clear() {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	cl.data = make(map[interface{}]interface{})
}
