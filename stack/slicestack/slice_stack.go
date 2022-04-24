package slicestack

import "fmt"

type SliceStack[T any] struct {
	data []T
}

// New constructs and returns an empty slice stack.
func New[T any]() *SliceStack[T] {
	return &SliceStack[T]{data: []T{}}
}

// Push adds an element on top of the stack.
func (s *SliceStack[T]) Push(data T) {
	s.data = append(s.data, data)
}

// Pop removes the top element and returns it. It returns false if the stack was empty.
func (s *SliceStack[T]) Pop() (val T, ok bool) {
	if s.IsEmpty() {
		return
	}

	val = s.data[len(s.data)-1]

	s.data = s.data[:len(s.data)-1]

	return val, true
}

// Top returns the top element of the stack. It returns false if the stack was empty.
func (s *SliceStack[T]) Top() (val T, ok bool) {
	if s.IsEmpty() {
		return
	}

	return s.data[len(s.data)-1], true
}

// Size returns the size of the stack.
func (s *SliceStack[T]) Size() int {
	return len(s.data)
}

// IsEmpty returns true if the stack doesn't have any elements.
func (s *SliceStack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// String returns the string representation of the stack.
func (s *SliceStack[T]) String() string {
	return fmt.Sprint(s.data)
}
