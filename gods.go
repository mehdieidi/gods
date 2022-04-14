package gods

import (
	"github.com/MehdiEidi/gods/linkedlist/circularly"
	"github.com/MehdiEidi/gods/linkedlist/doubly"
	"github.com/MehdiEidi/gods/linkedlist/singly"
	"github.com/MehdiEidi/gods/set"
)

type linked[T comparable] struct{}

func (l linked[T]) Singly() *singly.SinglyLinkedList[T] {
	return singly.New[T]()
}

func (l linked[T]) Doubly() *doubly.DoublyLinkedList[T] {
	return doubly.New[T]()
}

func (l linked[T]) Circularly() *circularly.CircularlyLinkedList[T] {
	return circularly.New[T]()
}

func Linked[T comparable]() linked[T] {
	return linked[T]{}
}

func Set[T comparable]() *set.Set[T] {
	return set.New[T]()
}
