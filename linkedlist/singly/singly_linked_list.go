package singlylinkedlist

import "strconv"

type SinglyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func New() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (s *SinglyLinkedList) IsEmpty() bool { return s.Size == 0 }

func (s *SinglyLinkedList) AddFirst(data int) {
	s.Head = &Node{Data: data, Next: s.Head}

	if s.Size == 0 {
		s.Tail = s.Head
	}

	s.Size++
}

func (s *SinglyLinkedList) AddLast(data int) {
	newNode := Node{Data: data}

	if s.IsEmpty() {
		s.Head = &newNode
	} else {
		s.Tail.Next = &newNode
	}

	s.Tail = &newNode

	s.Size++
}

func (s *SinglyLinkedList) RemoveFirst() {
	s.Head = s.Head.Next
	s.Size--

	if s.IsEmpty() {
		s.Tail = nil
	}
}

func (s *SinglyLinkedList) RemoveLast() {
	if s.Size == 1 {
		s.Tail = nil
		s.Head = nil
		s.Size--

		return
	}

	current := s.Head

	for ; current.Next.Next != nil; current = current.Next {
	}

	current.Next = nil
	s.Tail = current
	s.Size--

	if s.IsEmpty() {
		s.Tail = nil
		s.Head = nil
	}
}

func (s *SinglyLinkedList) String() string {
	str := "[ "

	current := s.Head

	for ; current != nil; current = current.Next {
		str += strconv.Itoa(current.Data) + " "
	}

	str += "]"

	return str
}
