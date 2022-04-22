package sortedpq

import (
	"github.com/MehdiEidi/gods/positionallist"
	"github.com/MehdiEidi/gods/priorityqueue"
	"golang.org/x/exp/constraints"
)

type sortedPQ[K constraints.Ordered, V any] struct {
	list       positionallist.PositionalList[*priorityqueue.Entry[K, V]]
	comparator priorityqueue.Comparator[K]
}

func New[K constraints.Ordered, V any](comparator priorityqueue.Comparator[K]) priorityqueue.PriorityQueue[K, V] {
	return &sortedPQ[K, V]{list: positionallist.NewDoublyPositionList[*priorityqueue.Entry[K, V]](), comparator: comparator}
}

func (s *sortedPQ[K, V]) Insert(key K, value V) *priorityqueue.Entry[K, V] {
	newEntry := &priorityqueue.Entry[K, V]{Key: key, Value: value}

	walk := s.list.Last()

	for walk.Element() != nil && walk != s.list.First() && s.comparator.Compare(newEntry.Key, walk.Element().Key) < 0 {
		walk = s.list.Before(walk)
	}

	if walk == nil {
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

func (s *sortedPQ[K, V]) RemoveMin() *priorityqueue.Entry[K, V] {
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
