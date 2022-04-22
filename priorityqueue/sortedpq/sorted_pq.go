package sortedpq

import (
	"github.com/MehdiEidi/gods/positionallist"
	"github.com/MehdiEidi/gods/priorityqueue"
)

type sortedPQ[K any, V any] struct {
	list       positionallist.PositionalList[*priorityqueue.Entry[K, V]]
	comparator priorityqueue.Comparator[K]
}

func New[K any, V any](comparator priorityqueue.Comparator[K]) priorityqueue.PriorityQueue[K, V] {
	return &sortedPQ[K, V]{list: positionallist.NewDoublyPositionList[*priorityqueue.Entry[K, V]](), comparator: comparator}
}

func (s *sortedPQ[K, V]) Enqueue(key K, value V) *priorityqueue.Entry[K, V] {
	newEntry := &priorityqueue.Entry[K, V]{Key: key, Value: value}

	walk := s.list.Last()

	for walk.Element() != nil && s.comparator.Compare(newEntry.Key, walk.Element().Key) < 0 {
		walk = s.list.Before(walk)
	}

	if walk.Element() == nil {
		s.list.AddFirst(newEntry)
	} else {
		s.list.AddAfter(walk, newEntry)
	}

	return newEntry
}

func (s *sortedPQ[K, V]) Min() *priorityqueue.Entry[K, V] {
	if s.list.IsEmpty() {
		return nil
	}
	return s.list.First().Element()
}

func (s *sortedPQ[K, V]) Dequeue() *priorityqueue.Entry[K, V] {
	if s.list.IsEmpty() {
		return nil
	}
	return s.list.Remove(s.list.First())
}

func (s *sortedPQ[K, V]) Size() int {
	return s.list.Size()
}

func (s *sortedPQ[K, V]) IsEmpty() bool {
	return s.list.Size() == 0
}

func (s *sortedPQ[K, V]) String() string {
	return s.list.String()
}
