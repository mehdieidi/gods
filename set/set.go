// Package set implements a set data structure for strings. It uses map as its backing data structure.
package set

import "sync"

// Set contains a map as its backing data structure and a mutex to ensure thread safety.
type Set struct {
	Lock sync.RWMutex
	Data map[string]struct{}
}

// New initializes and returns a new set.
func New() *Set {
	return &Set{sync.RWMutex{}, make(map[string]struct{})}
}

// Add just adds a new string to the set.
func (s *Set) Add(str string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.Data[str] = struct{}{}
}

// Has returns true if the set has str.
func (s *Set) Has(str string) bool {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	_, ok := s.Data[str]
	return ok
}

// Delete just deletes str from the set.
func (s *Set) Delete(str string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	delete(s.Data, str)
}
