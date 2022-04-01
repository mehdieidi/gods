// Package set implements a set data structure for strings. It uses map as its backing data structure.
package set

import "sync"

// Set contains a map as its backing data structure and a mutex to ensure thread safety.
type Set struct {
	Lock sync.Mutex
	Data map[string]struct{}
}

// New initializes and returns a new set.
func New() *Set {
	return &Set{sync.Mutex{}, make(map[string]struct{})}
}

// Add just adds a new string to the set.
func (s *Set) Add(str string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.Data[str] = struct{}{}
}

// Has returns true if the set has str.
func (s *Set) Has(str string) bool {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	_, ok := s.Data[str]
	return ok
}

// Delete just deletes str from the set.
func (s *Set) Delete(str string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	delete(s.Data, str)
}
