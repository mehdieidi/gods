package deque

import "github.com/MehdiEidi/gods/linkedlist/doubly"

type LinkedDeque[T comparable] struct {
	Data doubly.LinkedList[T]
}

func NewLinkedDeque[T comparable]() Deque[T] {
	return &LinkedDeque[T]{Data: *doubly.New[T]()}
}

func (ld *LinkedDeque[T]) Size() int {
	return ld.Data.Size
}

func (ld *LinkedDeque[T]) IsEmpty() bool {
	return ld.Size() == 0
}

func (ld *LinkedDeque[T]) First() (T, bool) {
	return ld.Data.First()
}

func (ld *LinkedDeque[T]) Last() (T, bool) {
	return ld.Data.Last()
}

func (ld *LinkedDeque[T]) AddFirst(data T) {
	ld.Data.AddFirst(data)
}

func (ld *LinkedDeque[T]) AddLast(data T) {
	ld.Data.AddLast(data)
}

func (ld *LinkedDeque[T]) RemoveFirst() (T, bool) {
	return ld.Data.RemoveFirst()
}

func (ld *LinkedDeque[T]) RemoveLast() (T, bool) {
	return ld.Data.RemoveLast()
}

func (ld *LinkedDeque[T]) String() string {
	return ld.Data.String()
}
