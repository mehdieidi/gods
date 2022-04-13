package singly

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

func (n *Node[T]) Equals(other *Node[T]) bool {
	if other == nil {
		return false
	}
	return n.Data == other.Data
}
