package circularly

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}
