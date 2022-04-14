package circularly

import "fmt"

type CircularlyLinkedList[T comparable] struct {
	Tail *Node[T]
	Size int
}

func New[T comparable]() *CircularlyLinkedList[T] {
	return &CircularlyLinkedList[T]{}
}

func (c *CircularlyLinkedList[T]) IsEmpty() bool { return c.Size == 0 }

func (c *CircularlyLinkedList[T]) First() (data T, ok bool) {
	if c.IsEmpty() {
		return
	}

	return c.Tail.Next.Data, true
}

func (c *CircularlyLinkedList[T]) Last() (data T, ok bool) {
	if c.IsEmpty() {
		return
	}

	return c.Tail.Data, true
}

func (c *CircularlyLinkedList[T]) Rotate() {
	if c.Tail != nil {
		c.Tail = c.Tail.Next
	}
}

func (c *CircularlyLinkedList[T]) AddFirst(data T) {
	if c.IsEmpty() {
		c.Tail = &Node[T]{Data: data}
		c.Tail.Next = c.Tail
	} else {
		newNode := Node[T]{Data: data, Next: c.Tail.Next}
		c.Tail.Next = &newNode
	}

	c.Size++
}

func (c *CircularlyLinkedList[T]) AddLast(data T) {
	c.AddFirst(data)
	c.Tail = c.Tail.Next
}

func (c *CircularlyLinkedList[T]) RemoveFirst() (data T, ok bool) {
	if c.IsEmpty() {
		return
	}

	head := c.Tail.Next

	val := head.Data

	if head == c.Tail {
		c.Tail = nil
	} else {
		c.Tail.Next = head.Next
	}

	c.Size--

	return val, true
}

func (c *CircularlyLinkedList[T]) String() string {
	if c.IsEmpty() {
		return "[ ]"
	}

	str := "[ "

	current := c.Tail.Next

	for ; current != c.Tail; current = current.Next {
		str += fmt.Sprint(current.Data) + " "
	}

	str += fmt.Sprint(c.Tail.Data) + " ]"

	return str
}
