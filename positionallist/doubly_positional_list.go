package positionallist

import (
	"fmt"
)

type doublyPositionalList[T comparable] struct {
	header  *node[T]
	trailer *node[T]
	size    int
}

func NewDoublyPositionList[T comparable]() PositionalList[T] {
	var d doublyPositionalList[T]

	d.header = &node[T]{}
	d.trailer = &node[T]{prev: d.header}
	d.header.next = d.trailer

	return &d
}

func (d *doublyPositionalList[T]) node(p Position[T]) *node[T] {
	return p.(*node[T])
}

func (d *doublyPositionalList[T]) position(n *node[T]) Position[T] {
	if n == d.header || n == d.trailer {
		return nil
	}

	return n
}

func (d *doublyPositionalList[T]) Size() int {
	return d.size
}

func (d *doublyPositionalList[T]) IsEmpty() bool {
	return d.size == 0
}

func (d *doublyPositionalList[T]) First() Position[T] {
	return d.header.next
}

func (d *doublyPositionalList[T]) Last() Position[T] {
	return d.trailer.prev
}

func (d *doublyPositionalList[T]) Before(p Position[T]) Position[T] {
	n := d.node(p)
	return n.prev
}

func (d *doublyPositionalList[T]) After(p Position[T]) Position[T] {
	n := d.node(p)
	return d.position(n.next)
}

func (d *doublyPositionalList[T]) addBetween(data T, predecessor *node[T], successor *node[T]) Position[T] {
	newNode := &node[T]{data: data, next: successor, prev: predecessor}

	predecessor.next = newNode
	successor.prev = newNode

	d.size++

	return newNode
}

func (d *doublyPositionalList[T]) AddFirst(data T) Position[T] {
	return d.addBetween(data, d.header, d.header.next)
}

func (d *doublyPositionalList[T]) AddLast(data T) Position[T] {
	return d.addBetween(data, d.trailer.prev, d.trailer)
}

func (d *doublyPositionalList[T]) AddBefore(p Position[T], data T) Position[T] {
	n := d.node(p)
	return d.addBetween(data, n.prev, n)
}

func (d *doublyPositionalList[T]) AddAfter(p Position[T], data T) Position[T] {
	n := d.node(p)
	return d.addBetween(data, n, n.next)
}

func (d *doublyPositionalList[T]) Set(p Position[T], data T) T {
	n := d.node(p)

	val := n.data
	n.data = data

	return val
}

func (d *doublyPositionalList[T]) Remove(p Position[T]) T {
	node := d.node(p)

	predecessor := node.prev
	successor := node.next

	predecessor.next = successor
	successor.prev = predecessor

	d.size--

	val := node.data

	var empty T
	node.data = empty
	node.next = nil
	node.prev = nil

	return val
}

func (d *doublyPositionalList[T]) String() string {
	str := "[ "

	for current := d.header.next; current != d.trailer; current = current.next {
		str += fmt.Sprint(current.data) + " "
	}

	str += "]"

	return str
}
