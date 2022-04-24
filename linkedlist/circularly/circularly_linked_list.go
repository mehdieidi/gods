package circularly

import "fmt"

type LinkedList[T any] struct {
	tail *Node[T]
	Size int
}

// New constructs and returns an empty circularly linked list.
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// IsEmpty returns true if the list doesn't have any elements.
func (c *LinkedList[T]) IsEmpty() bool {
	return c.Size == 0
}

// First returns the first element of the list. It returns false if the list is empty.
func (c *LinkedList[T]) First() (data T, ok bool) {
	if c.IsEmpty() {
		return
	}

	return c.tail.Next.Data, true
}

// Last returns the last element of the list. It returns false if the list is empty.
func (c *LinkedList[T]) Last() (data T, ok bool) {
	if c.IsEmpty() {
		return
	}

	return c.tail.Data, true
}

// Rotate rotates the list. It moves the first element to the end.
func (c *LinkedList[T]) Rotate() {
	if c.tail != nil {
		c.tail = c.tail.Next
	}
}

// AddFirst adds a new element to the first of the list.
func (c *LinkedList[T]) AddFirst(data T) {
	if c.IsEmpty() {
		c.tail = &Node[T]{Data: data}
		c.tail.Next = c.tail
	} else {
		newNode := Node[T]{Data: data, Next: c.tail.Next}
		c.tail.Next = &newNode
	}

	c.Size++
}

// AddLast adds a new element to the end of the list.
func (c *LinkedList[T]) AddLast(data T) {
	c.AddFirst(data)
	c.tail = c.tail.Next
}

// RemoveFirst removes and returns the first element of the list. It returns false if the list is empty.
func (c *LinkedList[T]) RemoveFirst() (data T, ok bool) {
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

// RemoveLast removes and returns the last element of the list. It returns false if the list empty.
func (c *LinkedList[T]) RemoveLast() (data T, ok bool) {
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

// String retruns the string representation of the list.
func (c *LinkedList[T]) String() string {
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
