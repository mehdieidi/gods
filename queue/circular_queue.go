package queue

import "github.com/MehdiEidi/gods/linkedlist/circularly"

type CircularQueue[T comparable] struct {
	data circularly.CircularlyLinkedList[T]
}

func NewCircularQueue[T comparable]() *CircularQueue[T] {
	return &CircularQueue[T]{}
}

func (cq *CircularQueue[T]) Enqueue(data T) {
	cq.data.AddLast(data)
}

func (cq *CircularQueue[T]) Dequeue() (val T, ok bool) {
	return cq.data.RemoveFirst()
}

func (cq *CircularQueue[T]) First() (val T, ok bool) {
	return cq.data.First()
}

func (cq *CircularQueue[T]) Size() int {
	return cq.data.Size
}

func (cq *CircularQueue[T]) IsEmpty() bool {
	return cq.data.IsEmpty()
}

func (cq *CircularQueue[T]) String() string {
	return cq.data.String()
}

func (cq *CircularQueue[T]) Rotate() {
	cq.data.Rotate()
}
