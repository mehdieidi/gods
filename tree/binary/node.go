package binarytree

type Node[T comparable] struct {
	Data   T
	Parent *Node[T]
	Left   *Node[T]
	Right  *Node[T]
}

func (n *Node[T]) Element() T {
	return n.Data
}
