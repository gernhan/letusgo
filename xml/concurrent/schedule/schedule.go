package schedule

import (
	"container/heap"
	"time"

	"github.com/gernhan/xml/concurrent"
	"github.com/gernhan/xml/models/schedule"
	"github.com/google/uuid"
)

var systemScheduler *Scheduler

type Scheduler struct {
	pool      concurrent.ThreadPool
	taskQueue *schedule.PriorityQueue
	signals   chan struct{}
}

func NewScheduler(maxWorkers int) *Scheduler {
	queue := make(schedule.PriorityQueue, 0)
	// Channel to signals the arrival of new items.
	newItem := make(chan struct{})
	s := &Scheduler{
		pool:      *concurrent.NewThreadPool(maxWorkers),
		taskQueue: &queue,
		signals:   newItem,
	}
	heap.Init(s.taskQueue)
	s.run()
	return s
}

func Init() {
	systemScheduler = NewScheduler(500)
}

func Schedule(task func() (interface{}, error), delay time.Duration) *concurrent.Future {
	return systemScheduler.Schedule(task, delay)
}

func (s *Scheduler) Schedule(task func() (interface{}, error), delay time.Duration) *concurrent.Future {
	id := uuid.NewString()
	item := schedule.ScheduledItem{
		ID:       id,
		Name:     id,
		DateTime: time.Now().Add(delay),
		Task:     task,
	}

	time.AfterFunc(delay, func() {
		heap.Push(s.taskQueue, &item)
		s.signals <- struct{}{} // Send a signals to the channel.
	})

	item.Result = concurrent.NewEmptyFuture()
	return item.Result
}

func (s *Scheduler) run() {
	// Start a goroutine to process new items.
	go func() {
		for {
			// Wait for a signal on the channel.
			<-s.signals

			// Process the new item.
			for s.taskQueue.Len() > 0 {
				item := heap.Pop(s.taskQueue).(*schedule.ScheduledItem)
				concurrent.RunAsyncWithPool(func() error {
					future := item.Result
					result, err := item.Task()
					future.SetErrorPointer(&err)
					future.SetDataPointer(&result)
					future.ReadyToWait()
					return err
				}, &s.pool)
			}
		}
	}()
}
