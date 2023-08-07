package schedule

import (
	"github.com/gernhan/xml/concurrent"
	"time"
)

type ScheduledItem struct {
	ID       string
	Name     string
	DateTime time.Time
	Task     func() (interface{}, error)
	Error    error
	Result   *concurrent.Future
	// Add other fields as needed
}

// PriorityQueue implements heap.Interface to act as a priority queue.
type PriorityQueue []*ScheduledItem

// Len returns the number of items in the queue.
func (pq *PriorityQueue) Len() int { return len(*pq) }

// Less checks if item i has a higher priority than item j.
func (pq *PriorityQueue) Less(i, j int) bool { return (*pq)[i].DateTime.Before((*pq)[j].DateTime) }

// Swap swaps the positions of items i and j.
func (pq *PriorityQueue) Swap(i, j int) { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }

// Push adds an item to the queue.
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*ScheduledItem)
	*pq = append(*pq, item)
}

// Pop removes and returns the item with the highest priority.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
