package doubly

import "strconv"

type DoublyLinkedList struct {
	Header  *Node
	Trailer *Node
	Size    int
}

func New() *DoublyLinkedList {
	var d DoublyLinkedList

	d.Header = &Node{}
	d.Trailer = &Node{Prev: d.Header}
	d.Header.Next = d.Trailer

	return &d
}

func (d *DoublyLinkedList) IsEmpty() bool { return d.Size == 0 }

func (d *DoublyLinkedList) First() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}

	return d.Header.Next.Data, true
}

func (d *DoublyLinkedList) Last() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}

	return d.Trailer.Prev.Data, true
}

func (d *DoublyLinkedList) AddBetween(data int, predecessor *Node, successor *Node) {
	newNode := &Node{Data: data, Next: successor, Prev: predecessor}

	predecessor.Next = newNode
	successor.Prev = newNode

	d.Size++
}

func (d *DoublyLinkedList) AddFirst(data int) {
	d.AddBetween(data, d.Header, d.Header.Next)
}

func (d *DoublyLinkedList) AddLast(data int) {
	d.AddBetween(data, d.Trailer.Prev, d.Trailer)
}

func (d *DoublyLinkedList) Remove(node *Node) int {
	predecessor := node.Prev
	successor := node.Next

	predecessor.Next = successor
	successor.Prev = predecessor

	d.Size--

	return node.Data
}

func (d *DoublyLinkedList) RemoveFirst() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}

	return d.Remove(d.Header.Next), true
}

func (d *DoublyLinkedList) RemoveLast() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}

	return d.Remove(d.Trailer.Prev), true
}

func (d *DoublyLinkedList) String() string {
	str := "[ "

	current := d.Header.Next

	for ; current != d.Trailer; current = current.Next {
		str += strconv.Itoa(current.Data) + " "
	}

	str += "]"

	return str
}
