package adaptablepq

import (
	"fmt"
)

type AdaptablePQ[K, V any] struct {
	heap       []*Entry[K, V]
	comparator Comparator[K]
}

// New constructs and returns an empty adaptable pq based on a min-heap.
func New[K, V any](comparator Comparator[K]) *AdaptablePQ[K, V] {
	return &AdaptablePQ[K, V]{heap: []*Entry[K, V]{}, comparator: comparator}
}

// Enqueue adds a new entry to the pq and returns it.
func (a *AdaptablePQ[K, V]) Enqueue(key K, value V) *Entry[K, V] {
	newEntry := &Entry[K, V]{Key: key, Value: value, Index: len(a.heap)}

	a.heap = append(a.heap, newEntry)
	a.upHeapBubble(len(a.heap) - 1)

	return newEntry
}

// Dequeue removes and returns the entry with the highest priority.
func (a *AdaptablePQ[K, V]) Dequeue() *Entry[K, V] {
	if a.IsEmpty() {
		return nil
	}

	res := a.heap[0]

	a.swap(0, len(a.heap)-1)
	a.heap = a.heap[:len(a.heap)-1]
	a.downHeapBubble(0)

	return res
}

// Min returns the entry with the highest priority.
func (a *AdaptablePQ[K, V]) Min() *Entry[K, V] {
	if a.IsEmpty() {
		return nil
	}
	return a.heap[0]
}

// Remove gets an entry and removes it from the pq.
func (a *AdaptablePQ[K, V]) Remove(entry *Entry[K, V]) {
	i := entry.Index

	if i == len(a.heap)-1 {
		a.heap = a.heap[:len(a.heap)-1]
	} else {
		a.swap(i, len(a.heap)-1)
		a.heap = a.heap[:len(a.heap)-1]
		a.bubble(i)
	}
}

// Size returns the number of the elements in the pq.
func (a *AdaptablePQ[K, V]) Size() int {
	return len(a.heap)
}

// IsEmpty returns true if the pq doesn't have any elements.
func (a *AdaptablePQ[K, V]) IsEmpty() bool {
	return a.Size() == 0
}

// ReplaceKey replaces the given entry's key and relocates it in the heap to the right position.
func (a *AdaptablePQ[K, V]) ReplaceKey(entry *Entry[K, V], key K) {
	entry.Key = key
	a.bubble(entry.Index)
}

// ReplaceValue just replaces the value of the given entry.
func (a *AdaptablePQ[K, V]) ReplaceValue(entry *Entry[K, V], value V) {
	entry.Value = value
}

// String returns the string representation of the pq.
func (a *AdaptablePQ[K, V]) String() string {
	str := "[ "

	for _, e := range a.heap {
		str += fmt.Sprint(e) + " "
	}

	str += "]"

	return str
}

// swap just swaps the entires in the heap.
func (a *AdaptablePQ[K, V]) swap(i, j int) {
	a.heap[i], a.heap[j] = a.heap[j], a.heap[i]

	a.heap[i].Index = i
	a.heap[j].Index = j
}

// bubble calls the right bubbling method based on the comparator.
func (a *AdaptablePQ[K, V]) bubble(i int) {
	if i > 0 && a.comparator.Compare(a.heap[i].Key, a.heap[a.parent(i)].Key) < 0 {
		a.upHeapBubble(i)
	} else {
		a.downHeapBubble(i)
	}
}

// parent returns the index of the parent of the given index of the entry in the heap.
func (a *AdaptablePQ[K, V]) parent(i int) int {
	return (i - 1) / 2
}

// left returns the index of the left child of the given entry's index in the heap.
func (a *AdaptablePQ[K, V]) left(i int) int {
	return 2*i + 1
}

// right returns the index of the right child of the given entry's index in the heap.
func (a *AdaptablePQ[K, V]) right(i int) int {
	return 2*i + 2
}

// hasLeft returns true if the given entry has a left child in the heap.
func (a *AdaptablePQ[K, V]) hasLeft(i int) bool {
	return a.left(i) < len(a.heap)
}

// hasRight returns true if the given entry has a right child in the heap.
func (a *AdaptablePQ[K, V]) hasRight(i int) bool {
	return a.right(i) < len(a.heap)
}

// downHeapBubble swaps the node downwards until reaching the right position in the heap.
func (a *AdaptablePQ[K, V]) downHeapBubble(i int) {
	// Until reaching bottom.
	for a.hasLeft(i) {
		leftIdx := a.left(i)
		smallChildIdx := leftIdx

		if a.hasRight(i) {
			rightIdx := a.right(i)

			if a.comparator.Compare(a.heap[leftIdx].Key, a.heap[rightIdx].Key) > 0 {
				smallChildIdx = rightIdx
			}
		}

		if a.comparator.Compare(a.heap[smallChildIdx].Key, a.heap[i].Key) >= 0 {
			break
		}

		a.swap(i, smallChildIdx)
		i = smallChildIdx
	}
}

// upHeapBubble swaps the node upwards until reaching the right position in the heap.
func (a *AdaptablePQ[K, V]) upHeapBubble(i int) {
	// Until reaching root.
	for i > 0 {
		p := a.parent(i)

		if a.comparator.Compare(a.heap[i].Key, a.heap[p].Key) >= 0 {
			break
		}

		a.swap(i, p)
		i = p
	}
}
