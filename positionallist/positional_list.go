package positionallist

import (
	"fmt"
)

type PositionalList[T any] struct {
	header  *Node[T] // header is a sentinel node. header.Next is the first element in the list.
	trailer *Node[T] // trailer is a sentinel node. trailer.Prev is the last element in the list.
	Size    int
}

// New constructs and returns an empty positional list.
func New[T any]() *PositionalList[T] {
	var d PositionalList[T]

	d.header = &Node[T]{}
	d.trailer = &Node[T]{Prev: d.header}
	d.header.Next = d.trailer

	return &d
}

// IsEmpty returns true if the list doesn't have any elements.
func (p *PositionalList[T]) IsEmpty() bool {
	return p.Size == 0
}

// First returns the first Node in the list.
func (p *PositionalList[T]) First() *Node[T] {
	return p.header.Next
}

// Last returns the last Node in the list.
func (p *PositionalList[T]) Last() *Node[T] {
	return p.trailer.Prev
}

// Before returns the Node before the given Node.
func (p *PositionalList[T]) Before(n *Node[T]) *Node[T] {
	return n.Prev
}

// After returns the Node after the given Node.
func (p *PositionalList[T]) After(n *Node[T]) *Node[T] {
	return p.validate(n.Next)
}

// AddFirst gets data, constructs a Node out of it, adds it to the first of the list, and returns it.
func (p *PositionalList[T]) AddFirst(data T) *Node[T] {
	return p.addBetween(data, p.header, p.header.Next)
}

// AddLast gets data, constructs a Node out of it, adds it to the end of the list, and returns it.
func (p *PositionalList[T]) AddLast(data T) *Node[T] {
	return p.addBetween(data, p.trailer.Prev, p.trailer)
}

// AddBefore gets data, constructs a Node out of it, adds it before the given Node, and returns it.
func (p *PositionalList[T]) AddBefore(n *Node[T], data T) *Node[T] {
	return p.addBetween(data, n.Prev, n)
}

// AddAfter gets data, constructs a Node out of it, adds it after the given Node, and returns it.
func (p *PositionalList[T]) AddAfter(n *Node[T], data T) *Node[T] {
	return p.addBetween(data, n, n.Next)
}

// Set changes the given Node's value to the given data. It returns the previous data.
func (p *PositionalList[T]) Set(n *Node[T], data T) T {
	val := n.Data
	n.Data = data
	return val
}

// Remove removes the given Node from the list and returns it.
func (p *PositionalList[T]) Remove(n *Node[T]) T {
	predecessor := n.Prev
	successor := n.Next

	predecessor.Next = successor
	successor.Prev = predecessor

	p.Size--

	val := n.Data

	var empty T
	n.Data = empty
	n.Next = nil
	n.Prev = nil

	return val
}

// String returns the string representation of the list.
func (d *PositionalList[T]) String() string {
	str := "[ "

	for current := d.header.Next; current != d.trailer; current = current.Next {
		str += fmt.Sprint(current.Data) + " "
	}

	str += "]"

	return str
}

// addBetween constructs a new Node out of the given data, adds it between the given two Nodes, and returns it.
func (d *PositionalList[T]) addBetween(data T, predecessor *Node[T], successor *Node[T]) *Node[T] {
	newNode := &Node[T]{Data: data, Next: successor, Prev: predecessor}

	predecessor.Next = newNode
	successor.Prev = newNode

	d.Size++

	return newNode
}

// validate returns nil if the given Node is a sentinel Node.
func (p *PositionalList[T]) validate(n *Node[T]) *Node[T] {
	if n == p.header || n == p.trailer {
		return nil
	}

	return n
}
