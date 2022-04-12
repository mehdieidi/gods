package circularly

import "strconv"

type CircularlyLinkedList struct {
	Tail *Node
	Size int
}

func New() *CircularlyLinkedList { return &CircularlyLinkedList{} }

func (c *CircularlyLinkedList) IsEmpty() bool { return c.Size == 0 }

func (c *CircularlyLinkedList) First() (int, bool) {
	if c.IsEmpty() {
		return 0, false
	}

	return c.Tail.Next.Data, true
}

func (c *CircularlyLinkedList) Last() (int, bool) {
	if c.IsEmpty() {
		return 0, false
	}

	return c.Tail.Data, true
}

func (c *CircularlyLinkedList) Rotate() {
	if c.Tail != nil {
		c.Tail = c.Tail.Next
	}
}

func (c *CircularlyLinkedList) AddFirst(data int) {
	if c.IsEmpty() {
		c.Tail = &Node{Data: data}
		c.Tail.Next = c.Tail
	} else {
		newNode := Node{Data: data, Next: c.Tail.Next}
		c.Tail.Next = &newNode
	}

	c.Size++
}

func (c *CircularlyLinkedList) AddLast(data int) {
	c.AddFirst(data)
	c.Tail = c.Tail.Next
}

func (c *CircularlyLinkedList) RemoveFirst() (int, bool) {
	if c.IsEmpty() {
		return 0, false
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

func (c *CircularlyLinkedList) String() string {
	if c.IsEmpty() {
		return "[ ]"
	}

	str := "[ "

	current := c.Tail.Next

	for ; current != c.Tail; current = current.Next {
		str += strconv.Itoa(current.Data) + " "
	}

	str += strconv.Itoa(c.Tail.Data) + " ]"

	return str
}
