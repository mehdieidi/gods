package queue

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() (T, bool)
	First() (T, bool)
	Size() int
	IsEmpty() bool
	String() string
}
