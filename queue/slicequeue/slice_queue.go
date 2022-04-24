package slicequeue

import "fmt"

type SliceQueue[T any] struct {
	data []T
}

// New constructs and returns an empty slice queue.
func New[T any]() *SliceQueue[T] {
	return &SliceQueue[T]{data: []T{}}
}

// Enqueue adds an element to the end of the queue.
func (q *SliceQueue[T]) Enqueue(data T) {
	q.data = append(q.data, data)
}

// Dequeue removes the front element of the queue and returns it. It returns false if the queue was empty.
func (q *SliceQueue[T]) Dequeue() (val T, ok bool) {
	if q.IsEmpty() {
		return
	}

	val = q.data[0]

	q.data = q.data[1:]

	return val, true
}

// First returns the front element of the queue. It returns false if the queue was empty.
func (q *SliceQueue[T]) First() (val T, ok bool) {
	if q.IsEmpty() {
		return
	}

	return q.data[0], true
}

// Size returns the number of elements in the queue.
func (q *SliceQueue[T]) Size() int {
	return len(q.data)
}

// IsEmpty returns true if the queue is empty.
func (q *SliceQueue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// String returns the string representation of the queue.
func (q *SliceQueue[T]) String() string {
	return fmt.Sprint(q.data)
}
