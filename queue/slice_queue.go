package queue

import "fmt"

type SliceQueue[T any] struct {
	Data []T
}

func NewSliceQueue[T any]() Queue[T] {
	return &SliceQueue[T]{Data: []T{}}
}

func (sq *SliceQueue[T]) Enqueue(data T) {
	sq.Data = append(sq.Data, data)
}

func (sq *SliceQueue[T]) Dequeue() (val T, ok bool) {
	if sq.IsEmpty() {
		return
	}

	val = sq.Data[0]

	sq.Data = sq.Data[1:]

	return val, true
}

func (sq *SliceQueue[T]) First() (val T, ok bool) {
	if sq.IsEmpty() {
		return
	}

	return sq.Data[0], true
}

func (sq *SliceQueue[T]) Size() int {
	return len(sq.Data)
}

func (sq *SliceQueue[T]) IsEmpty() bool {
	return len(sq.Data) == 0
}

func (sq *SliceQueue[T]) String() string {
	return fmt.Sprint(sq.Data)
}
