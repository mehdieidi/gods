package positionallist

type node[T comparable] struct {
	data T
	prev *node[T]
	next *node[T]
}

func (n *node[T]) Element() T {
	return n.data
}
