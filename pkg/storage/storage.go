// Package storage provides a thread-safe cyclic key-value storage with fixed capacity.
package storage

import (
	"fmt"
	"sync"
)

// Integer matches any integer type. It mirrors x/exp's constraints.Integer.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type cyclicItem[K Integer, T any] struct {
	key   K
	value T
}

// Cyclic is a limited size storage. Keys are of type K, items are of type T. Item with key n is stored to n (mod size) together with the key.
//
// The zero value is not usable; construct it with New.
type Cyclic[K Integer, T any] struct {
	values []*cyclicItem[K, T]
	mu     *sync.RWMutex
}

// Size is the size of cyclic storage.
func (s *Cyclic[K, T]) Size() int {
	return len(s.values)
}

// Store stores value with key to key (mod size).
func (s *Cyclic[K, T]) Store(key K, value T) {
	keyMod := key % K(len(s.values))

	// make sure 0 <= keyMod < Size
	if keyMod < 0 {
		keyMod += K(len(s.values))
	}

	storedItem := &cyclicItem[K, T]{key: key, value: value}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[keyMod] = storedItem
}

// Get retrieves element from key (mod size) if the stored element has the matching key.
func (s *Cyclic[K, T]) Get(key K) (T, bool) {
	var v T
	keyMod := key % K(len(s.values))

	// make sure 0 <= keyMod < Size
	if keyMod < 0 {
		keyMod += K(len(s.values))
	}

	s.mu.RLock()
	defer s.mu.RUnlock()
	storedItem := s.values[keyMod]

	if storedItem == nil {
		return v, false
	}

	if storedItem.key != key {
		return v, false
	}

	return storedItem.value, true
}

// NewCyclic initializes a Cyclic storage with size. Panics if size is not
// positive or cannot be represented losslessly by K (the modulo arithmetic
// in Store/Get otherwise produces out-of-range indices and a runtime
// slice-OOB panic on first access).
//
// Deprecated: Use New instead.
func NewCyclic[K Integer, T any](size int) Cyclic[K, T] {
	p := New[K, T](size)
	if p == nil {
		panic(fmt.Sprintf("storage.NewCyclic: invalid size %d for key type", size))
	}
	return *p
}

// New initializes a Cyclic storage with size.
//
// Returns nil if size is not positive, or if size cannot be represented
// losslessly by K (e.g., size=200 with K=int8 overflows to -56; the modulo
// arithmetic in Store/Get would then produce out-of-range indices and the
// slice access would panic at runtime). The roundtrip int(K(size)) == size
// catches both signed-narrowing and unsigned-truncation cases.
func New[K Integer, T any](size int) *Cyclic[K, T] {
	if size <= 0 {
		return nil
	}
	if int(K(size)) != size {
		return nil
	}
	return &Cyclic[K, T]{values: make([]*cyclicItem[K, T], size), mu: new(sync.RWMutex)}
}
