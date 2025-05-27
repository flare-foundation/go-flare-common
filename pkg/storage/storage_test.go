package storage_test

import (
	"testing"

	"github.com/flare-foundation/go-flare-common/pkg/storage"
	"github.com/stretchr/testify/require"
)

func TestCyclicSimple(t *testing.T) {
	const size int = 4

	stg := storage.NewCyclic[int, uint64](size)

	if stg.Size() != size {
		t.Error("wrong size")
	}

	getEmpty, exists := stg.Get(123)
	if exists || getEmpty != 0 {
		t.Error("unset but exists")
	}

	stg.Store(1, 1)

	getOne, exists := stg.Get(1)
	if !exists || getOne != 1 {
		t.Error("not stored")
	}

	getOneWithFive, exists := stg.Get(5)

	if exists || getOneWithFive != 0 {
		t.Error("returned with wrong key")
	}

	//overwrite
	stg.Store(5, 5)

	getOneWithFive, exists = stg.Get(5)

	if !exists || getOneWithFive != 5 {
		t.Error("not overwritten")
	}

	getOne, exists = stg.Get(1)

	if exists || getOne != 0 {
		t.Error("not overwritten with larger")
	}
}

func TestCyclicArray(t *testing.T) {
	const size int = 4

	stg := storage.NewCyclic[int, []uint64](size)

	_, exists := stg.Get(123)

	if exists {
		t.Error("unset but exists")
	}

	ar := []uint64{1}

	stg.Store(1, ar)

	getOne, exists := stg.Get(1)
	if !exists || len(getOne) != 1 {
		t.Error("not stored")
	}

	ar = append(ar, 1)

	getOne, exists = stg.Get(1)
	if !exists || len(getOne) != 1 || len(ar) != 2 {
		t.Error("copy not made 1")
	}

	getOne = append(getOne, 1)
	if len(getOne) != 2 {
		t.Error("element no added")
	}

	getOne, exists = stg.Get(1)

	if !exists || len(getOne) != 1 {
		t.Error("copy not made 2")
	}
}

func TestCyclicArrayPointer(t *testing.T) {
	const size int = 4

	stg := storage.NewCyclic[int, *[]uint64](size)

	_, exists := stg.Get(123)

	if exists {
		t.Error("unset but exists")
	}

	ar := []uint64{1}

	stg.Store(1, &ar)

	getOne, exists := stg.Get(1)

	if !exists || len(*getOne) != 1 {
		t.Error("not stored")
	}

	ar = append(ar, 1)

	getOne, exists = stg.Get(1)

	if !exists || len(*getOne) != 2 || len(ar) != 2 {
		t.Error("copy made")
	}

	*getOne = append(*getOne, 1)

	getOne, exists = stg.Get(1)

	if !exists || len(*getOne) != 3 {
		t.Error("copy not made 2")
	}
}

func TestCyclicMap(t *testing.T) {
	const size int = 4

	stg := storage.NewCyclic[int, map[uint64]uint64](size)

	_, exists := stg.Get(123)
	if exists {
		t.Error("unset but exists")
	}

	mp := map[uint64]uint64{1: 1}

	stg.Store(1, mp)

	_, exists = stg.Get(1)
	if !exists {
		t.Error("not stored")
	}

	mp[2] = 3

	getOne, exists := stg.Get(1)
	if !exists || getOne[2] != 3 {
		t.Error("map not modified")
	}
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

func TestNew(t *testing.T) {
	stg := storage.New[int, uint64](0)
	require.Nil(t, stg)

	stg = storage.New[int, uint64](-10)
	require.Nil(t, stg)
}
