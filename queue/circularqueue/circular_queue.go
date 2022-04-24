package circularqueue

import "github.com/MehdiEidi/gods/linkedlist/circularly"

type CircularQueue[T comparable] struct {
	data circularly.LinkedList[T]
}

// New constructs and returns an empty circular queue.
func New[T comparable]() *CircularQueue[T] {
	return &CircularQueue[T]{}
}

// Enqueue adds an element to the end of the queue.
func (q *CircularQueue[T]) Enqueue(data T) {
	q.data.AddLast(data)
}

// Dequeue removes the front element of the queue and returns it. It returns false if the queue was empty.
func (q *CircularQueue[T]) Dequeue() (val T, ok bool) {
	return q.data.RemoveFirst()
}

// First returns the front element of the queue. It returns false if the queue was empty.
func (q *CircularQueue[T]) First() (val T, ok bool) {
	return q.data.First()
}

// Size returns the number of elements in the queue.
func (q *CircularQueue[T]) Size() int {
	return q.data.Size
}

// IsEmpty returns true if the queue is empty.
func (q *CircularQueue[T]) IsEmpty() bool {
	return q.data.IsEmpty()
}

// String returns the string representation of the queue.
func (q *CircularQueue[T]) String() string {
	return q.data.String()
}

// Rotate rotates the queue. It moves the first element to the end of the queue.
func (q *CircularQueue[T]) Rotate() {
	q.data.Rotate()
}
