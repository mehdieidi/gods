package heappq

import (
	"fmt"

	"github.com/MehdiEidi/gods/priorityqueue"
)

type heapPQ[K any, V any] struct {
	heap       []*priorityqueue.Entry[K, V]
	comparator priorityqueue.Comparator[K]
}

func New[K any, V any](comparator priorityqueue.Comparator[K]) priorityqueue.PriorityQueue[K, V] {
	return &heapPQ[K, V]{heap: []*priorityqueue.Entry[K, V]{}, comparator: comparator}
}

func (h *heapPQ[K, V]) Enqueue(key K, value V) *priorityqueue.Entry[K, V] {
	newEntry := &priorityqueue.Entry[K, V]{Key: key, Value: value}

	h.heap = append(h.heap, newEntry)
	h.upHeapBubble(h.Size() - 1)

	return newEntry
}

func (h *heapPQ[K, V]) Dequeue() *priorityqueue.Entry[K, V] {
	if h.IsEmpty() {
		return nil
	}

	res := h.heap[0]

	h.swap(0, len(h.heap)-1)
	h.heap = h.heap[:len(h.heap)-1]
	h.downHeapBubble(0)

	return res
}

func (h *heapPQ[K, V]) Size() int {
	return len(h.heap)
}

func (h *heapPQ[K, V]) IsEmpty() bool {
	return h.Size() == 0
}

func (h *heapPQ[K, V]) Min() *priorityqueue.Entry[K, V] {
	if h.IsEmpty() {
		return nil
	}
	return h.heap[0]
}

func (h *heapPQ[K, V]) parent(i int) int {
	return (i - 1) / 2
}

func (h *heapPQ[K, V]) left(i int) int {
	return 2*i + 1
}

func (h *heapPQ[K, V]) right(i int) int {
	return 2*i + 2
}

func (h *heapPQ[K, V]) hasLeft(i int) bool {
	return h.left(i) < len(h.heap)
}

func (h *heapPQ[K, V]) hasRight(i int) bool {
	return h.right(i) < len(h.heap)
}

func (h *heapPQ[K, V]) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *heapPQ[K, V]) upHeapBubble(i int) {
	// Until reaching root.
	for i > 0 {
		p := h.parent(i)

		if h.comparator.Compare(h.heap[i].Key, h.heap[p].Key) >= 0 {
			break
		}

		h.swap(i, p)
		i = p
	}
}

func (h *heapPQ[K, V]) downHeapBubble(i int) {
	// Until reaching bottom.
	for h.hasLeft(i) {
		leftIdx := h.left(i)
		smallChildIdx := leftIdx

		if h.hasRight(i) {
			rightIdx := h.right(i)

			if h.comparator.Compare(h.heap[leftIdx].Key, h.heap[rightIdx].Key) > 0 {
				smallChildIdx = rightIdx
			}
		}

		if h.comparator.Compare(h.heap[smallChildIdx].Key, h.heap[i].Key) >= 0 {
			break
		}

		h.swap(i, smallChildIdx)
		i = smallChildIdx
	}
}

func (h *heapPQ[K, V]) String() string {
	str := "[ "

	for _, e := range h.heap {
		str += fmt.Sprint(e) + " "
	}

	str += "]"

	return str
}
