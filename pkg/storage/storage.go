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
func (s Cyclic[K, T]) Size() K {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return K(len(s.values))
}

// Store stores value with key to key (mod size).
func (s Cyclic[K, T]) Store(key K, value T) {
	keyMod := key % s.Size()

	// make sure  0 <= keyMod < Size
	if keyMod < 0 {
		keyMod += s.Size()
	}

	storedItem := &cyclicItem[K, T]{key: key, value: value}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[keyMod] = storedItem
}

// Get retrieves element from key (mod size) if the stored element has key key.
func (s Cyclic[K, T]) Get(key K) (T, bool) {
	var k T
	keyMod := key % s.Size()

	// make sure  0 <= keyMod < Size
	if keyMod < 0 {
		keyMod += s.Size()
	}

	s.mu.RLock()
	defer s.mu.RUnlock()
	storedItem := s.values[keyMod]

	if storedItem == nil {
		return k, false
	}

	storedKey := storedItem.key

	if storedKey != key {
		return k, false
	}

	k = storedItem.value

	return k, true
}

// Deprecated: NewCyclic initializes a Cyclic storage with size.
//
// Use New instead.
func NewCyclic[K constraints.Integer, T any](size int) Cyclic[K, T] {
	return Cyclic[K, T]{values: make([]*cyclicItem[K, T], size), mu: new(sync.RWMutex)}
}

// New initializes a Cyclic storage with size.
// If size is not positive, nil pointer is returned.
func New[K constraints.Integer, T any](size int) *Cyclic[K, T] {
	if size <= 0 {
		return nil
	}
	return &Cyclic[K, T]{values: make([]*cyclicItem[K, T], size), mu: new(sync.RWMutex)}
}
