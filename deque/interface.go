package deque

type Deque[T comparable] interface {
	Size() int
	IsEmpty() bool
	First() (T, bool)
	Last() (T, bool)
	AddFirst(T)
	AddLast(T)
	RemoveFirst() (T, bool)
	RemoveLast() (T, bool)
	String() string
}
