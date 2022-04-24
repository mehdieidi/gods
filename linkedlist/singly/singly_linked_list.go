package singly

import (
	"fmt"
)

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

// New constructs and returns an empty singlt linked list.
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// IsEmpty returns true if the linked list doesn't have any nodes.
func (s *LinkedList[T]) IsEmpty() bool {
	return s.Size == 0
}

// First returns the first element of the list. It returns false if the list is empty.
func (s *LinkedList[T]) First() (data T, ok bool) {
	if s.IsEmpty() {
		return
	}

	return s.Head.Data, true
}

// Last returns the last element of the list. It returns false if the list is empty.
func (s *LinkedList[T]) Last() (data T, ok bool) {
	if s.IsEmpty() {
		return
	}

	return s.Tail.Data, true
}

// AddFirst adds a new element to the first of the list.
func (s *LinkedList[T]) AddFirst(data T) {
	s.Head = &Node[T]{Data: data, Next: s.Head}

	if s.Size == 0 {
		s.Tail = s.Head
	}

	s.Size++
}

// AddLast adds a new element to the end of the list.
func (s *LinkedList[T]) AddLast(data T) {
	newNode := &Node[T]{Data: data}

	if s.IsEmpty() {
		s.Head = newNode
	} else {
		s.Tail.Next = newNode
	}

	s.Tail = newNode

	s.Size++
}

// Add adds an element to the given index in the list. It returns InvalidIndexErr if the given index is
// out of bound.
func (s *LinkedList[T]) Add(data T, index int) error {
	if index < 0 || index > s.Size {
		return InvalidIndexErr
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

// RemoveFirst removes and returns the first element of the list. It returns false if the list is empty.
func (s *LinkedList[T]) RemoveFirst() (val T, ok bool) {
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

// RemoveLast removes and returns the last element of the list. It returns false if the list empty.
func (s *LinkedList[T]) RemoveLast() (val T, ok bool) {
	if s.IsEmpty() {
		return
	}

	val = s.Tail.Data

	if s.Size == 1 {
		s.Tail = nil
		s.Head = nil
		s.Size--

		return val, true
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

	return val, true
}

// Remove removes the element in the given index. It returns false if the list is empty also it
// returns InvalidIndexErr if the given index is out of bound.
func (s *LinkedList[T]) Remove(index int) (val T, ok bool, err error) {
	if index < 0 || index >= s.Size {
		return val, false, InvalidIndexErr
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

// String retruns the string representation of the list.
func (s *LinkedList[T]) String() string {
	str := "[ "

	current := s.Head
	for ; current != nil; current = current.Next {
		str += fmt.Sprint(current.Data) + " "
	}

	str += "]"

	return str
}

// Clone constructs and returns a clone of the list.
func (s *LinkedList[T]) Clone() *LinkedList[T] {
	var newSingly LinkedList[T]

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

// ToSlice returns a slice of the list's elements.
func (s *LinkedList[T]) ToSlice() []T {
	r := make([]T, s.Size)

	for i, cur := 0, s.Head; cur != nil && i < len(r); i, cur = i+1, cur.Next {
		r[i] = cur.Data
	}

	return r
}
