package priority

import (
	"sync"

	heapt "github.com/flare-foundation/go-flare-common/pkg/heapt"
)

type Item[T any] struct {
	value  T   // The value of the item; arbitrary.
	weight int // The weight of the item in the queue.
	// The index is needed by update and is maintained by the heapt.Interface methods.
	index int // The index of the item in the heap.
}

// A Queue implements heapt.Interface
type Queue[T any] []*Item[T]

type QueueMutex[T any] struct {
	Queue[T]
	empty chan bool
	sync.RWMutex
}

func (q Queue[T]) Len() int { return len(q) }

func (q Queue[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, weight so we use greater than here.
	return q[i].weight > q[j].weight
}

func (q Queue[T]) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *Queue[T]) Pop() *Item[T] {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*q = old[0 : n-1]
	return item
}

func (q *Queue[T]) Push(item *Item[T]) {
	n := len(*q)
	item.index = n
	*q = append(*q, item)
}

func (q *Queue[T]) Update(item *Item[T], value T, priority int) {
	item.value = value
	item.weight = priority
	heapt.Fix(q, item.index)
}

// AddValue creates Item with value and weight, and adds it to the queue
func (q *Queue[T]) AddValue(value T, weight int) {
	item := new(Item[T])
	item.value = value
	item.weight = weight
	heapt.Push(q, item)
}
