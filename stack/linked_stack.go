package stack

import "github.com/MehdiEidi/gods/linkedlist/singly"

type LinkedStack[T comparable] struct {
	Data singly.SinglyLinkedList[T]
}

func NewLinkedStack[T comparable]() Stack[T] {
	return &LinkedStack[T]{}
}

func (ls *LinkedStack[T]) Push(data T) {
	ls.Data.AddFirst(data)
}

func (ls *LinkedStack[T]) Pop() (T, bool) {
	return ls.Data.RemoveFirst()
}

func (ls *LinkedStack[T]) Top() (T, bool) {
	return ls.Data.First()
}

func (ls *LinkedStack[T]) Size() int {
	return ls.Data.Size
}

func (ls *LinkedStack[T]) IsEmpty() bool {
	return ls.Data.IsEmpty()
}

func (ls *LinkedStack[T]) String() string {
	return ls.Data.String()
}
