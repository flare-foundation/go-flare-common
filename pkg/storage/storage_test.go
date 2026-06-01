package storage_test

import (
	"sync"
	"testing"

	"github.com/flare-foundation/go-flare-common/pkg/storage"
	"github.com/stretchr/testify/require"
)

func TestCyclicSimple(t *testing.T) {
	const size int = 4

	stg := storage.NewCyclic[int, uint64](size)

	require.Equal(t, size, stg.Size())

	getEmpty, exists := stg.Get(123)
	require.False(t, exists)
	require.Equal(t, uint64(0), getEmpty)

	stg.Store(1, 1)

	getOne, exists := stg.Get(1)
	require.True(t, exists)
	require.Equal(t, uint64(1), getOne)

	getOneWithFive, exists := stg.Get(5)
	require.False(t, exists)
	require.Equal(t, uint64(0), getOneWithFive)

	// overwrite
	stg.Store(5, 5)

	getOneWithFive, exists = stg.Get(5)
	require.True(t, exists)
	require.Equal(t, uint64(5), getOneWithFive)

	getOne, exists = stg.Get(1)
	require.False(t, exists)
	require.Equal(t, uint64(0), getOne)
}

func TestCyclicArray(t *testing.T) {
	const size int = 4

	stg := storage.NewCyclic[int, []uint64](size)

	_, exists := stg.Get(123)
	require.False(t, exists)

	ar := []uint64{1}

	stg.Store(1, ar)

	getOne, exists := stg.Get(1)
	require.True(t, exists)
	require.Len(t, getOne, 1)

	ar = append(ar, 1)

	getOne, exists = stg.Get(1)
	require.True(t, exists)
	require.Len(t, getOne, 1)
	require.Len(t, ar, 2)

	getOne = append(getOne, 1)
	require.Len(t, getOne, 2)

	getOne, exists = stg.Get(1)
	require.True(t, exists)
	require.Len(t, getOne, 1)
}

func TestCyclicArrayPointer(t *testing.T) {
	const size int = 4

	stg := storage.NewCyclic[int, *[]uint64](size)

	_, exists := stg.Get(123)
	require.False(t, exists)

	ar := []uint64{1}

	stg.Store(1, &ar)

	getOne, exists := stg.Get(1)
	require.True(t, exists)
	require.Len(t, *getOne, 1)

	ar = append(ar, 1)

	getOne, exists = stg.Get(1)
	require.True(t, exists)
	require.Len(t, *getOne, 2)
	require.Len(t, ar, 2)

	*getOne = append(*getOne, 1)

	getOne, exists = stg.Get(1)
	require.True(t, exists)
	require.Len(t, *getOne, 3)
}

func TestCyclicMap(t *testing.T) {
	const size int = 4

	stg := storage.NewCyclic[int, map[uint64]uint64](size)

	_, exists := stg.Get(123)
	require.False(t, exists)

	mp := map[uint64]uint64{1: 1}

	stg.Store(1, mp)

	_, exists = stg.Get(1)
	require.True(t, exists)

	mp[2] = 3

	getOne, exists := stg.Get(1)
	require.True(t, exists)
	require.Equal(t, uint64(3), getOne[2])
}

func TestCyclicNegative(t *testing.T) {
	const size int = 4

	stg := storage.New[int, uint64](size)
	require.NotNil(t, stg)
	require.Equal(t, size, stg.Size())

	for i := -10; i < 10; i++ {
		in := uint64(100 + i)

		stg.Store(i, in)

		out, ok := stg.Get(i)
		require.True(t, ok)
		require.Equal(t, in, out)
	}
}

// TestCyclicConcurrent exercises parallel Store/Get to catch races and
// verify that any successful Get observes the value paired with the queried key.
func TestCyclicConcurrent(t *testing.T) {
	const (
		ringSize   = 16
		keySpace   = 1024
		goroutines = 32
		opsEach    = 2000
	)

	stg := storage.New[int, int](ringSize)
	require.NotNil(t, stg)

	value := func(k int) int { return k*31 + 7 }

	var wg sync.WaitGroup
	wg.Add(2 * goroutines)

	for i := range goroutines {
		seed := i
		go func() {
			defer wg.Done()
			for j := range opsEach {
				k := (seed*opsEach + j) % keySpace
				stg.Store(k, value(k))
			}
		}()
		go func() {
			defer wg.Done()
			for j := range opsEach {
				k := (seed*opsEach + j) % keySpace
				if v, ok := stg.Get(k); ok {
					require.Equal(t, value(k), v)
				}
			}
		}()
	}

	wg.Wait()
}

func TestNew(t *testing.T) {
	stg := storage.New[int, uint64](0)
	require.Nil(t, stg)

	stg = storage.New[int, uint64](-10)
	require.Nil(t, stg)
}

// TestNewRejectsKeyTypeOverflow verifies that New refuses construction when K
// is narrower than int and cannot represent size, where K(size) wraps. For signed
// K the resulting K(len) is negative; key % K(len) can then produce
// out-of-range slice indices, and Store/Get panic on first access. New
// must refuse the construction instead of returning a panic-on-use struct.
func TestNewRejectsKeyTypeOverflow(t *testing.T) {
	// int8 max is 127; size 200 wraps to -56.
	require.Nil(t, storage.New[int8, uint64](200))

	// uint8 max is 255; size 300 wraps to 44.
	require.Nil(t, storage.New[uint8, uint64](300))

	// int16 max is 32767; size 40000 wraps.
	require.Nil(t, storage.New[int16, uint64](40000))

	// Boundary: int8 size 127 is the max representable, must be accepted.
	require.NotNil(t, storage.New[int8, uint64](127))
}

// TestNewCyclicPanicsOnInvalidSize verifies that the deprecated NewCyclic
// panics eagerly at construction rather than returning a struct that panics
// on first key % K(0).
func TestNewCyclicPanicsOnInvalidSize(t *testing.T) {
	require.Panics(t, func() { storage.NewCyclic[int, uint64](0) })
	require.Panics(t, func() { storage.NewCyclic[int, uint64](-5) })
	require.Panics(t, func() { storage.NewCyclic[int8, uint64](200) })
}
