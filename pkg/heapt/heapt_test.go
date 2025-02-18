package heapt_test

import (
	"fmt"
	"testing"

	heapt "github.com/flare-foundation/go-flare-common/pkg/heapt"
	"github.com/stretchr/testify/require"
)

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v int) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v int) {
	*h = append(*h, v)
}

func (h myHeap) verify(t *testing.T, i int) {
	t.Helper()
	n := h.Len()
	j1 := 2*i + 1
	j2 := 2*i + 2
	if j1 < n {
		require.Truef(t, !h.Less(j1, i), "heap invariant invalidated [%d] = %d > [%d] = %d", i, h[i], j1, h[j1])

		h.verify(t, j1)
	}
	if j2 < n {
		require.Truef(t, !h.Less(j1, i), "heap invariant invalidated [%d] = %d > [%d] = %d", i, h[i], j2, h[j2])
		h.verify(t, j2)
	}
}

func TestInitEmpty(t *testing.T) {
	h := new(myHeap)
	heapt.Init(h)
	h.verify(t, 0)

	_, exists := heapt.Pop(h)

	require.False(t, exists, "empty heap pops")
}

func TestInit0(t *testing.T) {
	h := new(myHeap)
	for i := 20; i > 0; i-- {
		h.Push(0) // all elements are the same
	}
	heapt.Init(h)
	h.verify(t, 0)

	for i := 1; h.Len() > 0; i++ {
		x, exists := heapt.Pop(h)
		require.Equal(t, 0, x)
		require.True(t, exists, "non empty heap did not pop")

		h.verify(t, 0)
	}
}

func TestInit1(t *testing.T) {
	h := new(myHeap)
	for i := 20; i > 0; i-- {
		h.Push(i) // all elements are different
	}
	heapt.Init(h)
	h.verify(t, 0)

	for i := 1; h.Len() > 0; i++ {
		x, _ := heapt.Pop(h)
		h.verify(t, 0)
		require.Equalf(t, i, x, "%d.th pop got %d; want %d", i, x, i)
	}
}

func Test(t *testing.T) {
	h := new(myHeap)
	h.verify(t, 0)

	for i := 20; i > 10; i-- {
		h.Push(i)
	}
	heapt.Init(h)
	h.verify(t, 0)

	for i := 10; i > 0; i-- {
		heapt.Push(h, i)
		h.verify(t, 0)
	}

	for i := 1; h.Len() > 0; i++ {
		x, _ := heapt.Pop(h)
		if i < 20 {
			heapt.Push(h, 20+i)
		}
		h.verify(t, 0)
		require.Equalf(t, i, x, "%d.th pop got %d; want %d", i, x, i)
	}
}

func TestRemove0(t *testing.T) {
	h := new(myHeap)
	for i := 0; i < 10; i++ {
		h.Push(i)
	}
	heapt.Init(h)
	h.verify(t, 0)
	for h.Len() > 0 {
		i := h.Len() - 1
		x := heapt.Remove(h, i)
		require.Equalf(t, i, x, "%d.th remove got %d; want %d", i, x, i)
		h.verify(t, 0)
	}
}

func TestRemove1(t *testing.T) {
	h := new(myHeap)
	for i := 0; i < 10; i++ {
		h.Push(i)
	}
	heapt.Init(h)
	h.verify(t, 0)

	for i := 0; h.Len() > 0; i++ {
		x := heapt.Remove(h, 0)
		fmt.Printf("h: %v\n", h)
		require.Equalf(t, i, x, "Remove(0) got %d; want %d", i, x, i)
		h.verify(t, 0)
	}
}

func TestRemove2(t *testing.T) {
	N := 10

	h := new(myHeap)
	for i := 0; i < N; i++ {
		h.Push(i)
	}
	heapt.Init(h)
	h.verify(t, 0)

	m := make(map[int]bool)
	for h.Len() > 0 {
		m[heapt.Remove(h, (h.Len()-1)/2)] = true
		h.verify(t, 0)
	}

	require.Equalf(t, N, len(m), "len(m) = %d; want %d", len(m), N)
	for i := 0; i < len(m); i++ {
		require.Truef(t, m[i], "m[%d] doesn't exist", i)

	}
}

func TestFix(t *testing.T) {
	N := 10

	h := new(myHeap)
	for i := 0; i < N; i++ {
		heapt.Push(h, i)
	}
	h.verify(t, 0)

	(*h)[N/2] = -N

	heapt.Fix(h, N/2)

	h.verify(t, 0)
}

func TestFixInexistentIndex(t *testing.T) {
	N := 10

	h := new(myHeap)
	for i := 0; i < N; i++ {
		heapt.Push(h, i)
	}
	h.verify(t, 0)

	heapt.Fix(h, -1)

	h.verify(t, 0)

	heapt.Fix(h, 2*N)

	h.verify(t, 0)
}
