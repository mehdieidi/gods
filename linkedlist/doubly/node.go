package doubly

type Node[T comparable] struct {
	Data T
	Prev *Node[T]
	Next *Node[T]
}
