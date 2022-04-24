package heappq

import (
	"fmt"
)

type HeapPQ[K any, V any] struct {
	heap       []*Entry[K, V]
	comparator Comparator[K]
}

// New constructs and returns an empty priority queue based on a min-heap.
func New[K any, V any](comparator Comparator[K]) *HeapPQ[K, V] {
	return &HeapPQ[K, V]{heap: []*Entry[K, V]{}, comparator: comparator}
}

// NewWithValues gets slices of keys and values and constructs a priority queue using heapify buttom-up construction.
func NewWithValues[K any, V any](keys []K, values []V, comparator Comparator[K]) (*HeapPQ[K, V], error) {
	if len(keys) != len(values) {
		return nil, KeysValuesNotSameLenErr
	}

	h := New[K, V](comparator)

	for i := 0; i < len(keys); i++ {
		h.heap = append(h.heap, &Entry[K, V]{Key: keys[i], Value: values[i]})
	}

	h.heapify()

	return h, nil
}

// Enqueue constructs an entry, adds it to the pq, and returns it.
func (h *HeapPQ[K, V]) Enqueue(key K, value V) *Entry[K, V] {
	newEntry := &Entry[K, V]{Key: key, Value: value}

	h.heap = append(h.heap, newEntry)
	h.upHeapBubble(h.Size() - 1)

	return newEntry
}

// Dequeue removes the entry with the highest priority and returns it.
func (h *HeapPQ[K, V]) Dequeue() *Entry[K, V] {
	if h.IsEmpty() {
		return nil
	}

	res := h.heap[0]

	h.swap(0, len(h.heap)-1)
	h.heap = h.heap[:len(h.heap)-1]
	h.downHeapBubble(0)

	return res
}

// Min returns the entry with the highest priority (that entry is the first element in the heap).
func (h *HeapPQ[K, V]) Min() *Entry[K, V] {
	if h.IsEmpty() {
		return nil
	}
	return h.heap[0]
}

// Size returns the number of elements in the pq.
func (h *HeapPQ[K, V]) Size() int {
	return len(h.heap)
}

// IsEmpty returns true if the pq doesn't have any elements.
func (h *HeapPQ[K, V]) IsEmpty() bool {
	return h.Size() == 0
}

// parent returns the index of the parent of the given index of the entry in the heap.
func (h *HeapPQ[K, V]) parent(i int) int {
	return (i - 1) / 2
}

// left returns the index of the left child of the given entry's index in the heap.
func (h *HeapPQ[K, V]) left(i int) int {
	return 2*i + 1
}

// right returns the index of the right child of the given entry's index in the heap.
func (h *HeapPQ[K, V]) right(i int) int {
	return 2*i + 2
}

// hasLeft returns true if the given entry has a left child in the heap.
func (h *HeapPQ[K, V]) hasLeft(i int) bool {
	return h.left(i) < len(h.heap)
}

// hasRight returns true if the given entry has a right child in the heap.
func (h *HeapPQ[K, V]) hasRight(i int) bool {
	return h.right(i) < len(h.heap)
}

// swap just swaps two entries in the heap.
func (h *HeapPQ[K, V]) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

// upHeapBubble swaps the node upwards until reaching the right position in the heap.
func (h *HeapPQ[K, V]) upHeapBubble(i int) {
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

// downHeapBubble swaps the node downwards until reaching the right position in the heap.
func (h *HeapPQ[K, V]) downHeapBubble(i int) {
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

// String returns the string representation of the pq.
func (h *HeapPQ[K, V]) String() string {
	str := "[ "

	for _, e := range h.heap {
		str += fmt.Sprint(e) + " "
	}

	str += "]"

	return str
}

// heapify calls downHeapBubble on the nodes to relocate them to the right position. It is used in
// button-up heap construction with the given slices of keys and values.
func (h *HeapPQ[K, V]) heapify() {
	for i := h.parent(h.Size() - 1); i >= 0; i-- {
		h.downHeapBubble(i)
	}
}
