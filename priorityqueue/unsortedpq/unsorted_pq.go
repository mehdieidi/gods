package unsortedpq

import (
	"github.com/MehdiEidi/gods/positionallist"
)

type UnsortedPQ[K any, V any] struct {
	list       *positionallist.PositionalList[*Entry[K, V]]
	comparator Comparator[K]
}

// New constructs and returns an empty priority queue based on an unsorted positional list.
func New[K any, V any](comparator Comparator[K]) *UnsortedPQ[K, V] {
	return &UnsortedPQ[K, V]{list: positionallist.New[*Entry[K, V]](), comparator: comparator}
}

// Enqueue constructs a new entry, adds it to the priority queue, and returns it.
func (u *UnsortedPQ[K, V]) Enqueue(key K, value V) *Entry[K, V] {
	newEntry := &Entry[K, V]{Key: key, Value: value}
	u.list.AddLast(newEntry)
	return newEntry
}

// Dequeue removes and returns the entry with heighest priority (lower numbers indicate higher priority).
func (u *UnsortedPQ[K, V]) Dequeue() *Entry[K, V] {
	if u.list.IsEmpty() {
		return nil
	}
	return u.list.Remove(u.findMin())
}

// Min returns the entry which has the highest priority (lower numbers indicate higher priority).
func (u *UnsortedPQ[K, V]) Min() *Entry[K, V] {
	if u.list.IsEmpty() {
		return nil
	}
	return u.findMin().Data
}

// Size returns the number of the elements in the priority queue.
func (u *UnsortedPQ[K, V]) Size() int {
	return u.list.Size
}

// IsEmpty returns true if the pq doesn't have any elements.
func (u *UnsortedPQ[K, V]) IsEmpty() bool {
	return u.Size() == 0
}

// String returns the string representation of the pq.
func (u *UnsortedPQ[K, V]) String() string {
	return u.list.String()
}

// findMin walks the list and finds the entry with the highest priority (lower numbers indicate higher priority).
func (u *UnsortedPQ[K, V]) findMin() *positionallist.Node[*Entry[K, V]] {
	min := u.list.First()

	for walk := u.list.After(min); walk != nil; walk = u.list.After(walk) {
		if u.comparator.Compare(walk.Data.Key, min.Data.Key) < 0 {
			min = walk
		}
	}

	return min
}
