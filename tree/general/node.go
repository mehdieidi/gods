package tree

type Node[T comparable] struct {
	Data     T
	Parent   *Node[T]
	Children []*Node[T]
}

func (n *Node[T]) Element() T {
	return n.Data
}
