package positionallist

type Node[T any] struct {
	Data T
	Prev *Node[T]
	Next *Node[T]
}
