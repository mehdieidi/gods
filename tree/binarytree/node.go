package binarytree

import "fmt"

type Node[T any] struct {
	Data   T
	Parent *Node[T]
	Left   *Node[T]
	Right  *Node[T]
}

// String returns a string representation of the node's data.
func (n *Node[T]) String() string {
	return fmt.Sprint(n.Data)
}
