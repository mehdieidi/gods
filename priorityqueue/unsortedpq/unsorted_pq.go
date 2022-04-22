package unsortedpq

import (
	"github.com/MehdiEidi/gods/positionallist"
	"github.com/MehdiEidi/gods/priorityqueue"
)

type unsortedPQ[K any, V any] struct {
	list       positionallist.PositionalList[*priorityqueue.Entry[K, V]]
	comparator priorityqueue.Comparator[K]
}

func New[K any, V any](comparator priorityqueue.Comparator[K]) priorityqueue.PriorityQueue[K, V] {
	return &unsortedPQ[K, V]{list: positionallist.NewDoublyPositionList[*priorityqueue.Entry[K, V]](), comparator: comparator}
}

func (u *unsortedPQ[K, V]) findMin() positionallist.Position[*priorityqueue.Entry[K, V]] {
	min := u.list.First()

	for walk := u.list.After(u.list.First()); walk != nil; walk = u.list.After(walk) {
		if u.comparator.Compare(walk.Element().Key, min.Element().Key) < 0 {
			min = walk
		}
	}

	return min
}

func (u *unsortedPQ[K, V]) Enqueue(key K, value V) *priorityqueue.Entry[K, V] {
	newEntry := &priorityqueue.Entry[K, V]{Key: key, Value: value}
	u.list.AddLast(newEntry)
	return newEntry
}

func (u *unsortedPQ[K, V]) Min() *priorityqueue.Entry[K, V] {
	if u.list.IsEmpty() {
		return nil
	}
	return u.findMin().Element()
}

func (u *unsortedPQ[K, V]) Dequeue() *priorityqueue.Entry[K, V] {
	if u.list.IsEmpty() {
		return nil
	}
	return u.list.Remove(u.findMin())
}

func (u *unsortedPQ[K, V]) Size() int {
	return u.list.Size()
}

func (u *unsortedPQ[K, V]) IsEmpty() bool {
	return u.Size() == 0
}

func (u *unsortedPQ[K, V]) String() string {
	return u.list.String()
}
