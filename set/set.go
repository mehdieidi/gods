// Package set implements a set data structure for comparables. It uses map as its backing data structure.
package set

import "sync"

// Set contains a map as its backing data structure and a mutex to ensure thread safety.
type Set[T comparable] struct {
	Lock sync.RWMutex
	Data map[T]struct{}
}

// New initializes and returns a new set.
func New[T comparable]() *Set[T] {
	return &Set[T]{sync.RWMutex{}, make(map[T]struct{})}
}

// Add just adds a new value to the set.
func (s *Set[T]) Add(val T) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.Data[val] = struct{}{}
}

// Has returns true if the set has val.
func (s *Set[T]) Has(val T) bool {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	_, ok := s.Data[val]
	return ok
}

// Delete just deletes val from the set.
func (s *Set) Delete(val T) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	delete(s.Data, val)
}
