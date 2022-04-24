package doubly

import "fmt"

type LinkedList[T any] struct {
	header  *Node[T] // header is a sentinel node. header.Next is the first element in the list.
	trailer *Node[T] // trailer is a sentinel node. trailer.Prev is the last element in the list.
	Size    int
}

// New constructs and returns an empty doubly linked list.
func New[T any]() *LinkedList[T] {
	var d LinkedList[T]

	d.header = &Node[T]{}
	d.trailer = &Node[T]{Prev: d.header}
	d.header.Next = d.trailer

	return &d
}

// IsEmpty returns true if the list doesn't have any elements.
func (d *LinkedList[T]) IsEmpty() bool {
	return d.Size == 0
}

// First returns the first element of the list. It returns false if the list is empty.
func (d *LinkedList[T]) First() (data T, ok bool) {
	if d.IsEmpty() {
		return
	}

	return d.header.Next.Data, true
}

// Last returns the last element of the list. It returns false if the list is empty.
func (d *LinkedList[T]) Last() (data T, ok bool) {
	if d.IsEmpty() {
		return
	}

	return d.trailer.Prev.Data, true
}

// AddBetween gets two nodes, constructs a new node out of the given data and adds that node in between them.
func (d *LinkedList[T]) AddBetween(data T, predecessor *Node[T], successor *Node[T]) {
	newNode := &Node[T]{Data: data, Next: successor, Prev: predecessor}

	predecessor.Next = newNode
	successor.Prev = newNode

	d.Size++
}

// AddFirst adds a new element to the first of the list.
func (d *LinkedList[T]) AddFirst(data T) {
	d.AddBetween(data, d.header, d.header.Next)
}

// AddLast adds a new element to the end of the list.
func (d *LinkedList[T]) AddLast(data T) {
	d.AddBetween(data, d.trailer.Prev, d.trailer)
}

// Remove removes the given node from the list. It returns the removed node's data.
func (d *LinkedList[T]) Remove(node *Node[T]) T {
	predecessor := node.Prev
	successor := node.Next

	predecessor.Next = successor
	successor.Prev = predecessor

	d.Size--

	return node.Data
}

// RemoveFirst removes and returns the first element of the list. It returns false if the list is empty.
func (d *LinkedList[T]) RemoveFirst() (data T, ok bool) {
	if d.IsEmpty() {
		return
	}

	return d.Remove(d.header.Next), true
}

// RemoveLast removes and returns the last element of the list. It returns false if the list empty.
func (d *LinkedList[T]) RemoveLast() (data T, ok bool) {
	if d.IsEmpty() {
		return
	}

	return d.Remove(d.trailer.Prev), true
}

// String retruns the string representation of the list.
func (d *LinkedList[T]) String() string {
	str := "[ "

	for current := d.header.Next; current != d.trailer; current = current.Next {
		str += fmt.Sprint(current.Data) + " "
	}

	str += "]"

	return str
}

// Clone constructs and returns a clone of the list.
func (d *LinkedList[T]) Clone() *LinkedList[T] {
	newDoubly := New[T]()

	if d.IsEmpty() {
		return newDoubly
	}

	newDoubly.header.Next = &Node[T]{Data: d.header.Next.Data}
	newDoubly.Size++

	newDoublyTail := newDoubly.header.Next

	walk := d.header.Next.Next
	var n *Node[T]

	for walk != d.trailer {
		n = &Node[T]{Data: walk.Data, Prev: newDoublyTail}

		newDoublyTail.Next = n
		newDoubly.Size++

		newDoublyTail = n
		walk = walk.Next
	}

	n.Next = newDoubly.trailer
	newDoubly.trailer.Prev = n

	return newDoubly
}

// ToSlice returns a slice of the list's elements.
func (d *LinkedList[T]) ToSlice() []T {
	r := make([]T, d.Size)

	for i, cur := 0, d.header.Next; cur != d.trailer && i < len(r); i, cur = i+1, cur.Next {
		r[i] = cur.Data
	}

	return r
}
