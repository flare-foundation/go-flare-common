package priority

import (
	"sync"

	heapt "github.com/flare-foundation/go-flare-common/pkg/heapt"
)

type weight[T any] interface {
	Self() T
	Less(T) bool
}

type Item[T any, W weight[W]] struct {
	value  T         // The value of the item.
	weight weight[W] // The weight of the item in the queue.
	// The index is needed by update and is maintained by the heapt.Interface methods.
	index int // The index of the item in the heap.
}

// A Queue implements heapt.Interface
type Queue[T any, W weight[W]] []*Item[T, W]

type QueueMutex[T any, W weight[W]] struct {
	Queue[T, W]
	empty chan bool
	sync.RWMutex
}

func (q Queue[T, W]) Len() int { return len(q) }

func (q Queue[T, W]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, weight so we use greater than here.
	return q[j].weight.Less(q[i].weight.Self())
}

// Swap SHOULD NOT BE USED DIRECTLY
func (q Queue[T, W]) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

// Pop SHOULD NOT BE USED DIRECTLY use heapt.Pop(q) instead
func (q *Queue[T, W]) Pop() *Item[T, W] {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*q = old[0 : n-1]
	return item
}

// Push SHOULD NOT BE USED DIRECTLY use heapt.Push(q, item) instead
func (q *Queue[T, W]) Push(item *Item[T, W]) {
	n := len(*q)
	item.index = n
	*q = append(*q, item)
}

// Update updates the weight of an item and updates the queue accordingly.
//
// Does not effect the queue if the item is not in it.
func (q *Queue[T, W]) UpdateWeight(item *Item[T, W], weight W) {
	item.weight = weight
	heapt.Fix(q, item.index)
}

// AddValue creates Item with value and weight, and adds it to the queue.
func (q *Queue[T, W]) AddValue(value T, weight W) {
	item := new(Item[T, W])
	item.value = value
	item.weight = weight
	heapt.Push(q, item)
}
