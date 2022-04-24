package sortedpq

import (
	"github.com/MehdiEidi/gods/positionallist"
)

type SortedPQ[K any, V any] struct {
	list       *positionallist.PositionalList[*Entry[K, V]]
	comparator Comparator[K]
}

// New constructs and returns an empty priority queue based on a sorted positional list.
func New[K any, V any](comparator Comparator[K]) *SortedPQ[K, V] {
	return &SortedPQ[K, V]{list: positionallist.New[*Entry[K, V]](), comparator: comparator}
}

// Enqueue walks the pq, finds the correct position to add the new constructed Node. It returns the new Node.
func (s *SortedPQ[K, V]) Enqueue(key K, value V) *Entry[K, V] {
	newEntry := &Entry[K, V]{Key: key, Value: value}

	walk := s.list.Last()

	for walk.Data != nil && s.comparator.Compare(newEntry.Key, walk.Data.Key) < 0 {
		walk = s.list.Before(walk)
	}

	if walk.Data == nil {
		s.list.AddFirst(newEntry)
	} else {
		s.list.AddAfter(walk, newEntry)
	}

	return newEntry
}

// Dequeue removes and returns the entry with the highest priority (that entry is in the front of the list).
func (s *SortedPQ[K, V]) Dequeue() *Entry[K, V] {
	if s.list.IsEmpty() {
		return nil
	}
	return s.list.Remove(s.list.First())
}

// Min returns the entry with the highest priority (that entry is in the front of the list).
func (s *SortedPQ[K, V]) Min() *Entry[K, V] {
	if s.list.IsEmpty() {
		return nil
	}
	return s.list.First().Data
}

// Size returns the number of elements in the pq.
func (s *SortedPQ[K, V]) Size() int {
	return s.list.Size
}

// IsEmpty returns true if the pq doesn't have any elements.
func (s *SortedPQ[K, V]) IsEmpty() bool {
	return s.list.Size == 0
}

// String returns the string representation of the pq.
func (s *SortedPQ[K, V]) String() string {
	return s.list.String()
}
