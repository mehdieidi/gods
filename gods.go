package gods

import (
	"github.com/MehdiEidi/gods/bitset"
	"github.com/MehdiEidi/gods/deque"
	"github.com/MehdiEidi/gods/linkedlist/circularly"
	"github.com/MehdiEidi/gods/linkedlist/doubly"
	"github.com/MehdiEidi/gods/linkedlist/singly"
	"github.com/MehdiEidi/gods/queue/circularqueue"
	"github.com/MehdiEidi/gods/queue/linkedqueue"
	"github.com/MehdiEidi/gods/queue/slicequeue"
	"github.com/MehdiEidi/gods/set"
	"github.com/MehdiEidi/gods/stack/linkedstack"
	"github.com/MehdiEidi/gods/stack/slicestack"
	"github.com/MehdiEidi/gods/tree/binarytree"
	"github.com/MehdiEidi/gods/tree/generaltree"
)

func NewSinglyLinkedList[T any]() *singly.LinkedList[T] {
	return singly.New[T]()
}

func NewDoublyLinkedList[T any]() *doubly.LinkedList[T] {
	return doubly.New[T]()
}

func NewCircularlyLinkedList[T any]() *circularly.LinkedList[T] {
	return circularly.New[T]()
}

func NewSet[T comparable]() *set.Set[T] {
	return set.New[T]()
}

func NewDeque[T any]() *deque.Deque[T] {
	return deque.New[T]()
}

func NewLinkedQueue[T any]() *linkedqueue.LinkedQueue[T] {
	return linkedqueue.New[T]()
}

func NewCircularQueue[T any]() *circularqueue.CircularQueue[T] {
	return circularqueue.New[T]()
}

func NewSliceQueue[T any]() *slicequeue.SliceQueue[T] {
	return slicequeue.New[T]()
}

func NewLinkedStack[T any]() *linkedstack.LinkedStack[T] {
	return linkedstack.New[T]()
}

func NewSliceStack[T any]() *slicestack.SliceStack[T] {
	return slicestack.New[T]()
}

func NewBitset8() *bitset.Bitset8 {
	return bitset.NewBitset8()
}

func NewBinaryTree[T any]() *binarytree.BinaryTree[T] {
	return binarytree.New[T]()
}

func NewGeneralTree[T any]() *generaltree.GeneralTree[T] {
	return generaltree.New[T]()
}
