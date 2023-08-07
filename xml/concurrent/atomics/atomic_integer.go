package atomics

import (
	"strconv"
)

type AtomicInteger struct {
	data *AtomicReference
}

func NewAtomicInteger() *AtomicInteger {
	number := int64(0)
	reference := NewAtomicReference(&number)
	return &AtomicInteger{
		data: reference,
	}
}

func NewAtomicIntegerWithInitial(initValue int64) *AtomicInteger {
	number := initValue
	reference := NewAtomicReference(&number)
	return &AtomicInteger{
		data: reference,
	}
}

func (a *AtomicInteger) AddAndGet(delta int64) int64 {
	return *a.data.Mutate(func(objPointer interface{}) {
		number := objPointer.(*int64)
		*number = *number + delta
	}).(*int64)
}

func (a *AtomicInteger) IncrementAndGet() int64 {
	return a.AddAndGet(1)
}

func (a *AtomicInteger) DecrementAndGet() int64 {
	return a.AddAndGet(-1)
}

func (a *AtomicInteger) Get() int64 {
	pointer := a.data.dataPointer.(*int64)
	return *pointer
}

func (a *AtomicInteger) String() string {
	pointer := a.data.dataPointer.(*int64)
	return strconv.FormatInt(*pointer, 10)
}
