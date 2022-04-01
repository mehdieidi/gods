// Package set implements a set data structure for strings. It uses map as its backing data structure.
package set

import "sync"

// Set contains a map as its backing data structure and a mutex to ensure thread safety.
type Set[T any] struct {
	Lock sync.RWMutex
	Data map[any]struct{}
}

// New initializes and returns a new set.
func New[T any]() *Set[T] {
	return &Set{sync.RWMutex{}, make(map[T]struct{})}
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
