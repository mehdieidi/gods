package queue

import (
	"github.com/MehdiEidi/gods/linkedlist/singly"
)

type LinkedQueue[T comparable] struct {
	Data singly.SinglyLinkedList[T]
}

func NewLinkedQueue[T comparable]() Queue[T] {
	return &LinkedQueue[T]{}
}

func (lq *LinkedQueue[T]) Enqueue(data T) {
	lq.Data.AddLast(data)
}

func (lq *LinkedQueue[T]) Dequeue() (val T, ok bool) {
	return lq.Data.RemoveFirst()
}

func (lq *LinkedQueue[T]) First() (val T, ok bool) {
	return lq.Data.First()
}

func (lq *LinkedQueue[T]) Size() int {
	return lq.Data.Size
}

func (lq *LinkedQueue[T]) IsEmpty() bool {
	return lq.Data.IsEmpty()
}

func (lq *LinkedQueue[T]) String() string {
	return lq.Data.String()
}
