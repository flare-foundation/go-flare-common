// Generically typed heap implementation based on "container/heap"

package heapt

import "sort"

// The Interface type describes the requirements
// for a type using the routines in this package.
// Any type that implements it may be used as a
// min-heap with the following invariants (established after
// [Init] has been called or if the data is empty or sorted):
//
//	!h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
//
// Note that [Push] and [Pop] in this interface are for package heap's
// implementation to call. To add and remove things from the heap,
// use [heap.Push] and [heap.Pop].
type Interface[T any] interface {
	sort.Interface
	Push(item T) // add item as element Len()
	Pop() T      // remove and return element Len() - 1.
}

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
func Init[T any](h Interface[T]) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

// Push pushes the item onto the heap.
// The complexity is O(log n) where n = h.Len().
func Push[T any](h Interface[T], item T) {
	h.Push(item)
	up(h, h.Len()-1)
}

// Pop removes and returns the minimum element (according to Less) from the heap
// if it is nonempty. Also returns a boolean indicator of existence of the element.
//
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to [Remove](h, 0).
func Pop[T any](h Interface[T]) (T, bool) {
	var item T
	if h.Len() < 1 {
		return item, false
	}

	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop(), true
}

// Remove removes and returns the element at index i from the heap.
//
// Panics if i is out of bounds.
// The complexity is O(log n) where n = h.Len().
func Remove[T any](h Interface[T], i int) T {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Ineffective if i is out of bounds.
//
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling [Remove](h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func Fix[T any](h Interface[T], i int) {
	// do nothing if index our of bounds and avoid panic
	if i < 0 || i >= h.Len() {
		return
	}

	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

func up[T any](h Interface[T], j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || i < 0 || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down[T any](h Interface[T], i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
