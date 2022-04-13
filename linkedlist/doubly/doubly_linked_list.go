package doubly

import "fmt"

type DoublyLinkedList[T comparable] struct {
	Header  *Node[T]
	Trailer *Node[T]
	Size    int
}

func New[T comparable]() *DoublyLinkedList[T] {
	var d DoublyLinkedList[T]

	d.Header = &Node[T]{}
	d.Trailer = &Node[T]{Prev: d.Header}
	d.Header.Next = d.Trailer

	return &d
}

func (d *DoublyLinkedList[T]) IsEmpty() bool { return d.Size == 0 }

func (d *DoublyLinkedList[T]) First() (data T, ok bool) {
	if d.IsEmpty() {
		return
	}

	return d.Header.Next.Data, true
}

func (d *DoublyLinkedList[T]) Last() (data T, ok bool) {
	if d.IsEmpty() {
		return
	}

	return d.Trailer.Prev.Data, true
}

func (d *DoublyLinkedList[T]) AddBetween(data T, predecessor *Node[T], successor *Node[T]) {
	newNode := &Node[T]{Data: data, Next: successor, Prev: predecessor}

	predecessor.Next = newNode
	successor.Prev = newNode

	d.Size++
}

func (d *DoublyLinkedList[T]) AddFirst(data T) {
	d.AddBetween(data, d.Header, d.Header.Next)
}

func (d *DoublyLinkedList[T]) AddLast(data T) {
	d.AddBetween(data, d.Trailer.Prev, d.Trailer)
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

	return d.Remove(d.Header.Next), true
}

func (d *DoublyLinkedList[T]) RemoveLast() (data T, ok bool) {
	if d.IsEmpty() {
		return
	}

	return d.Remove(d.Trailer.Prev), true
}

func (d *DoublyLinkedList[T]) String() string {
	str := "[ "

	current := d.Header.Next

	for ; current != d.Trailer; current = current.Next {
		str += fmt.Sprint(current.Data) + " "
	}

	str += "]"

	return str
}
