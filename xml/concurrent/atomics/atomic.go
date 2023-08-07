package atomics

import (
	"sync"
)

type AtomicReference struct {
	dataPointer interface{}
	mu          *sync.Mutex
}

func NewAtomicReference(dataPointer interface{}) *AtomicReference {
	return &AtomicReference{
		dataPointer: dataPointer,
		mu:          &sync.Mutex{},
	}
}

func (a *AtomicReference) Mutate(handler func(objPointer interface{})) interface{} {
	a.mu.Lock()
	defer a.mu.Unlock()
	handler(a.dataPointer)
	return a.dataPointer
}

func (a *AtomicReference) Get() interface{} {
	pointer := a.dataPointer.(*interface{})
	return *pointer
}
