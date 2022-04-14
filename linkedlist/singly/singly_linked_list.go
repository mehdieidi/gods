package singly

import (
	"errors"
	"fmt"
)

// TODO: reverse, find, remove by value

type SinglyLinkedList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

func New[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (s *SinglyLinkedList[T]) IsEmpty() bool {
	return s.Size == 0
}

func (s *SinglyLinkedList[T]) First() (data T, ok bool) {
	if s.IsEmpty() {
		return
	}

	return s.Head.Data, true
}

func (s *SinglyLinkedList[T]) Last() (data T, ok bool) {
	if s.IsEmpty() {
		return
	}

	return s.Tail.Data, true
}

func (s *SinglyLinkedList[T]) AddFirst(data T) {
	s.Head = &Node[T]{Data: data, Next: s.Head}

	if s.Size == 0 {
		s.Tail = s.Head
	}

	s.Size++
}

func (s *SinglyLinkedList[T]) AddLast(data T) {
	newNode := Node[T]{Data: data}

	if s.IsEmpty() {
		s.Head = &newNode
	} else {
		s.Tail.Next = &newNode
	}

	s.Tail = &newNode

	s.Size++
}

func (s *SinglyLinkedList[T]) Add(data T, index int) error {
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

	newNode := Node[T]{Data: data, Next: current.Next}
	current.Next = &newNode

	s.Size++

	return nil
}

func (s *SinglyLinkedList[T]) RemoveFirst() (val T, ok bool) {
	if s.IsEmpty() {
		return
	}

	val = s.Head.Data

	s.Head = s.Head.Next
	s.Size--

	if s.IsEmpty() {
		s.Tail = nil
	}

	return val, true
}

func (s *SinglyLinkedList[T]) RemoveLast() (val T, ok bool) {
	if s.IsEmpty() {
		return
	}

	if s.Size == 1 {
		val = s.Tail.Data

		s.Tail = nil
		s.Head = nil
		s.Size--

		return val, true
	}

	val = s.Tail.Data

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

	return val, true
}

func (s *SinglyLinkedList[T]) Remove(index int) (val T, ok bool, err error) {
	if index < 0 || index >= s.Size {
		return val, false, errors.New("invalid index")
	}

	if index == 0 {
		val, ok = s.RemoveFirst()
		return val, ok, nil
	}

	if index == s.Size-1 {
		val, ok = s.RemoveLast()
		return val, ok, nil
	}

	var count int

	current := s.Head
	for ; current != nil; current = current.Next {
		count++
		if count == index {
			break
		}
	}

	val = current.Next.Data

	current.Next = current.Next.Next
	s.Size--

	return val, true, nil
}

func (s *SinglyLinkedList[T]) String() string {
	str := "[ "

	current := s.Head
	for ; current != nil; current = current.Next {
		str += fmt.Sprint(current.Data) + " "
	}

	str += "]"

	return str
}

func (s *SinglyLinkedList[T]) Equals(other *SinglyLinkedList[T]) bool {
	if other == nil || s.Size != other.Size {
		return false
	}

	current1 := s.Head
	current2 := other.Head

	for current1 != nil {
		if !current1.Equals(current2) {
			return false
		}

		current1 = current1.Next
		current2 = current2.Next
	}

	return true
}

func (s *SinglyLinkedList[T]) Clone() *SinglyLinkedList[T] {
	var newSingly SinglyLinkedList[T]

	if s.Size == 0 {
		return &newSingly
	}

	newSingly.Head = &Node[T]{Data: s.Head.Data}
	newSingly.Size++

	newSinglyTail := newSingly.Head

	walk := s.Head.Next

	for walk != nil {
		n := &Node[T]{Data: walk.Data}

		newSinglyTail.Next = n
		newSingly.Size++

		newSinglyTail = n
		walk = walk.Next
	}

	return &newSingly
}

func (s *SinglyLinkedList[T]) ToSlice() []T {
	r := make([]T, s.Size)

	for i, cur := 0, s.Head; cur != nil && i < len(r); i, cur = i+1, cur.Next {
		r[i] = cur.Data
	}

	return r
}
