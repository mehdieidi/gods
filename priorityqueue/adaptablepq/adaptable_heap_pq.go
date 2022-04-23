package adaptablepq

import (
	"fmt"

	"github.com/MehdiEidi/gods/priorityqueue"
)

type AdaptablePQ[K, V any] struct {
	heap       []*Entry[K, V]
	comparator priorityqueue.Comparator[K]
}

func New[K, V any](comparator priorityqueue.Comparator[K]) *AdaptablePQ[K, V] {
	return &AdaptablePQ[K, V]{heap: []*Entry[K, V]{}, comparator: comparator}
}

func (a *AdaptablePQ[K, V]) Enqueue(key K, value V) *Entry[K, V] {
	newEntry := &Entry[K, V]{Key: key, Value: value, Index: len(a.heap)}

	a.heap = append(a.heap, newEntry)
	a.upHeapBubble(len(a.heap) - 1)

	return newEntry
}

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

func (a *AdaptablePQ[K, V]) Size() int {
	return len(a.heap)
}

func (a *AdaptablePQ[K, V]) IsEmpty() bool {
	return a.Size() == 0
}

func (a *AdaptablePQ[K, V]) Min() *Entry[K, V] {
	if a.IsEmpty() {
		return nil
	}

	return a.heap[0]
}

func (a *AdaptablePQ[K, V]) ReplaceKey(entry *Entry[K, V], key K) {
	entry.Key = key
	a.bubble(entry.Index)
}

func (a *AdaptablePQ[K, V]) ReplaceValue(entry *Entry[K, V], value V) {
	entry.Value = value
}

func (a *AdaptablePQ[K, V]) String() string {
	str := "[ "

	for _, e := range a.heap {
		str += fmt.Sprint(e) + " "
	}

	str += "]"

	return str
}

func (a *AdaptablePQ[K, V]) swap(i, j int) {
	a.heap[i], a.heap[j] = a.heap[j], a.heap[i]

	a.heap[i].Index = i
	a.heap[j].Index = j
}

func (a *AdaptablePQ[K, V]) bubble(i int) {
	if i > 0 && a.comparator.Compare(a.heap[i].Key, a.heap[a.parent(i)].Key) < 0 {
		a.upHeapBubble(i)
	} else {
		a.downHeapBubble(i)
	}
}

func (a *AdaptablePQ[K, V]) parent(i int) int {
	return (i - 1) / 2
}

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

func (a *AdaptablePQ[K, V]) left(i int) int {
	return 2*i + 1
}

func (a *AdaptablePQ[K, V]) right(i int) int {
	return 2*i + 2
}

func (a *AdaptablePQ[K, V]) hasLeft(i int) bool {
	return a.left(i) < len(a.heap)
}

func (a *AdaptablePQ[K, V]) hasRight(i int) bool {
	return a.right(i) < len(a.heap)
}
