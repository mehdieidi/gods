package circularly

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

func (n *Node[T]) Equals(other *Node[T]) bool {
	return n.Data == other.Data
}
