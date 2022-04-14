package gods

import (
	"github.com/MehdiEidi/gods/linkedlist/circularly"
	"github.com/MehdiEidi/gods/linkedlist/doubly"
	"github.com/MehdiEidi/gods/linkedlist/singly"
	"github.com/MehdiEidi/gods/set"
)

func NewSinglyLinkedList[T comparable]() *singly.SinglyLinkedList[T] {
	return singly.New[T]()
}

func NewDoublyLinkedList[T comparable]() *doubly.DoublyLinkedList[T] {
	return doubly.New[T]()
}

func NewCircularlyLinkedList[T comparable]() *circularly.CircularlyLinkedList[T] {
	return circularly.New[T]()
}

func NewSet[T comparable]() *set.Set[T] {
	return set.New[T]()
}
