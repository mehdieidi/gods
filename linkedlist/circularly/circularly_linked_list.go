package circularly

import "fmt"

type CircularlyLinkedList[T comparable] struct {
	tail *Node[T]
	Size int
}

func New[T comparable]() *CircularlyLinkedList[T] {
	return &CircularlyLinkedList[T]{}
}

func (c *CircularlyLinkedList[T]) IsEmpty() bool {
	return c.Size == 0
}

func (c *CircularlyLinkedList[T]) First() (data T, ok bool) {
	if c.IsEmpty() {
		return
	}

	return c.tail.Next.Data, true
}

func (c *CircularlyLinkedList[T]) Last() (data T, ok bool) {
	if c.IsEmpty() {
		return
	}

	return c.tail.Data, true
}

func (c *CircularlyLinkedList[T]) Rotate() {
	if c.tail != nil {
		c.tail = c.tail.Next
	}
}

func (c *CircularlyLinkedList[T]) AddFirst(data T) {
	if c.IsEmpty() {
		c.tail = &Node[T]{Data: data}
		c.tail.Next = c.tail
	} else {
		newNode := Node[T]{Data: data, Next: c.tail.Next}
		c.tail.Next = &newNode
	}

	c.Size++
}

func (c *CircularlyLinkedList[T]) AddLast(data T) {
	c.AddFirst(data)
	c.tail = c.tail.Next
}

func (c *CircularlyLinkedList[T]) RemoveFirst() (data T, ok bool) {
	if c.IsEmpty() {
		return
	}

	head := c.tail.Next

	val := head.Data

	if head == c.tail {
		c.tail = nil
	} else {
		c.tail.Next = head.Next
	}

	c.Size--

	return val, true
}

func (c *CircularlyLinkedList[T]) RemoveLast() (data T, ok bool) {
	if c.IsEmpty() {
		return
	}

	data = c.tail.Data

	current := c.tail.Next
	for ; current.Next != c.tail; current = current.Next {
	}

	current.Next = c.tail.Next
	c.tail = current

	c.Size--

	return data, true
}

func (c *CircularlyLinkedList[T]) String() string {
	if c.IsEmpty() {
		return "[ ]"
	}

	str := "[ "

	current := c.tail.Next

	for ; current != c.tail; current = current.Next {
		str += fmt.Sprint(current.Data) + " "
	}

	str += fmt.Sprint(c.tail.Data) + " ]"

	return str
}

func (c *CircularlyLinkedList[T]) Equals(other *CircularlyLinkedList[T]) bool {
	if other == nil || c.Size != other.Size {
		return false
	}

	current1 := c.tail.Next
	current2 := other.tail.Next

	for current1 != c.tail {
		if !current1.Equals(current2) {
			return false
		}

		current1 = current1.Next
		current2 = current2.Next
	}

	return true
}
