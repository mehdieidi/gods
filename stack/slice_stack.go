package stack

import "fmt"

type SliceStack[T any] struct {
	Data []T
}

func NewSliceStack[T any]() Stack[T] {
	return &SliceStack[T]{Data: []T{}}
}

func (ss *SliceStack[T]) Push(data T) {
	ss.Data = append(ss.Data, data)
}

func (ss *SliceStack[T]) Pop() (val T, ok bool) {
	if ss.IsEmpty() {
		return
	}

	val = ss.Data[len(ss.Data)-1]

	ss.Data = ss.Data[:len(ss.Data)-1]

	return val, true
}

func (ss *SliceStack[T]) Top() (val T, ok bool) {
	if ss.IsEmpty() {
		return
	}

	return ss.Data[len(ss.Data)-1], true
}

func (ss *SliceStack[T]) Size() int {
	return len(ss.Data)
}

func (ss *SliceStack[T]) IsEmpty() bool {
	return ss.Size() == 0
}

func (ss *SliceStack[T]) String() string {
	return fmt.Sprint(ss.Data)
}
