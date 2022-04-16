package stack

type Stack[T any] interface {
	Push(T)
	Pop() (T, bool)
	Top() (T, bool)
	Size() int
	IsEmpty() bool
}
