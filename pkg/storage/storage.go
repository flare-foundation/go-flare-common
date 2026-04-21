// Package storage provides a thread-safe cyclic key-value storage with fixed capacity.
package storage

import (
	"sync"

	"golang.org/x/exp/constraints"
)

type cyclicItem[K constraints.Integer, T any] struct {
	key   K
	value T
}

// Cyclic is a limited size storage. Keys are of type K, items are of type T. Item with key n is stored to n (mod size) together with the key.
type Cyclic[K constraints.Integer, T any] struct {
	values []*cyclicItem[K, T]
	mu     *sync.RWMutex
}

// Size is the size of cyclic storage.
func (s *Cyclic[K, T]) Size() K {
	return K(len(s.values))
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

// Deprecated: NewCyclic initializes a Cyclic storage with size.
//
// Use New instead.
func NewCyclic[K constraints.Integer, T any](size int) Cyclic[K, T] {
	return Cyclic[K, T]{values: make([]*cyclicItem[K, T], size), mu: new(sync.RWMutex)}
}

// New initializes a Cyclic storage with size.
// If size is not positive, nil is returned.
func New[K constraints.Integer, T any](size int) *Cyclic[K, T] {
	if size <= 0 {
		return nil
	}
	return &Cyclic[K, T]{values: make([]*cyclicItem[K, T], size), mu: new(sync.RWMutex)}
}
