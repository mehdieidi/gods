package doubly

import "fmt"

type DoublyLinkedList[T comparable] struct {
	header  *Node[T]
	trailer *Node[T]
	Size    int
}

func New[T comparable]() *DoublyLinkedList[T] {
	var d DoublyLinkedList[T]

	d.header = &Node[T]{}
	d.trailer = &Node[T]{Prev: d.header}
	d.header.Next = d.trailer

	return &d
}

func (d *DoublyLinkedList[T]) IsEmpty() bool {
	return d.Size == 0
}

func (d *DoublyLinkedList[T]) First() (data T, ok bool) {
	if d.IsEmpty() {
		return
	}

	return d.header.Next.Data, true
}

func (d *DoublyLinkedList[T]) Last() (data T, ok bool) {
	if d.IsEmpty() {
		return
	}

	return d.trailer.Prev.Data, true
}

func (d *DoublyLinkedList[T]) AddBetween(data T, predecessor *Node[T], successor *Node[T]) {
	newNode := &Node[T]{Data: data, Next: successor, Prev: predecessor}

	predecessor.Next = newNode
	successor.Prev = newNode

	d.Size++
}

func (d *DoublyLinkedList[T]) AddFirst(data T) {
	d.AddBetween(data, d.header, d.header.Next)
}

func (d *DoublyLinkedList[T]) AddLast(data T) {
	d.AddBetween(data, d.trailer.Prev, d.trailer)
}

func (d *DoublyLinkedList[T]) Remove(node *Node[T]) T {
	predecessor := node.Prev
	successor := node.Next

	predecessor.Next = successor
	successor.Prev = predecessor

	d.Size--

	return node.Data
}

func (d *DoublyLinkedList[T]) RemoveFirst() (data T, ok bool) {
	if d.IsEmpty() {
		return
	}

	return d.Remove(d.header.Next), true
}

func (d *DoublyLinkedList[T]) RemoveLast() (data T, ok bool) {
	if d.IsEmpty() {
		return
	}

	return d.Remove(d.trailer.Prev), true
}

func (d *DoublyLinkedList[T]) String() string {
	str := "[ "

	for current := d.header.Next; current != d.trailer; current = current.Next {
		str += fmt.Sprint(current.Data) + " "
	}

	str += "]"

	return str
}

func (d *DoublyLinkedList[T]) Equals(other *DoublyLinkedList[T]) bool {
	if other == nil {
		return false
	}

	current1 := d.header.Next
	current2 := other.header.Next

	for current1 != nil {
		if !current1.Equals(current2) {
			return false
		}

		current1 = current1.Next
		current2 = current2.Next
	}

	return true
}

func (d *DoublyLinkedList[T]) Clone() *DoublyLinkedList[T] {
	newDoubly := New[T]()

	if d.Size == 0 {
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

func (d *DoublyLinkedList[T]) ToSlice() []T {
	r := make([]T, d.Size)

	for i, cur := 0, d.header.Next; cur != d.trailer && i < len(r); i, cur = i+1, cur.Next {
		r[i] = cur.Data
	}

	return r
}
