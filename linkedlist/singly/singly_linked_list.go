package singlylinkedlist

import (
	"errors"
	"strconv"
)

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

func (s *SinglyLinkedList) Add(data int, index int) error {
	if index < 0 || index > s.Size {
		return errors.New("invalid index")
	}

	if index == 0 {
		s.AddFirst(data)
		return nil
	}
	if index == s.Size {
		s.AddLast(data)
		return nil
	}

	var count int

	current := s.Head
	for ; current != nil; current = current.Next {
		count++
		if count == index {
			break
		}
	}

	newNode := Node{Data: data, Next: current.Next}
	current.Next = &newNode

	s.Size++

	return nil
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

func (s *SinglyLinkedList) Remove(index int) (int, error) {
	if index < 0 || index > s.Size {
		return 0, errors.New("invalid index")
	}

	if index == 0 {
		s.RemoveFirst()
		return 0, nil
	}
	if index == s.Size-1 {
		s.RemoveLast()
		return 0, nil
	}

	var count int

	current := s.Head
	for ; current != nil; current = current.Next {
		count++
		if count == index {
			break
		}
	}

	val := current.Next.Data

	current.Next = current.Next.Next
	s.Size--

	return val, nil
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
