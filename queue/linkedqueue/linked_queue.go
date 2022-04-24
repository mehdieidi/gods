package linkedqueue

import (
	"github.com/MehdiEidi/gods/linkedlist/singly"
)

type LinkedQueue[T any] struct {
	data singly.LinkedList[T]
}

// New constructs and returns an empty linked queue.
func New[T any]() *LinkedQueue[T] {
	return &LinkedQueue[T]{}
}

// Enqueue adds an element to the end of the queue.
func (q *LinkedQueue[T]) Enqueue(data T) {
	q.data.AddLast(data)
}

// Dequeue removes the front element of the queue and returns it. It returns false if the queue was empty.
func (q *LinkedQueue[T]) Dequeue() (val T, ok bool) {
	return q.data.RemoveFirst()
}

// First returns the front element of the queue. It returns false if the queue was empty.
func (q *LinkedQueue[T]) First() (val T, ok bool) {
	return q.data.First()
}

// Size returns the number of elements in the queue.
func (q *LinkedQueue[T]) Size() int {
	return q.data.Size
}

// IsEmpty returns true if the queue is empty.
func (q *LinkedQueue[T]) IsEmpty() bool {
	return q.data.IsEmpty()
}

// String returns the string representation of the queue.
func (q *LinkedQueue[T]) String() string {
	return q.data.String()
}
