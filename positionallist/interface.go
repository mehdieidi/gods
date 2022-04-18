package positionallist

type Position[T comparable] interface {
	Element() T
}

type PositionalList[T comparable] interface {
	Size() int
	IsEmpty() bool
	First() Position[T]
	Last() Position[T]
	Before(Position[T]) Position[T]
	After(Position[T]) Position[T]
	AddLast(T) Position[T]
	AddFirst(T) Position[T]
	AddBefore(Position[T], T) Position[T]
	AddAfter(Position[T], T) Position[T]
	Set(Position[T], T) T
	Remove(Position[T]) T
}
