package singly

import "fmt"

type Node[T any] struct {
	Data T
	Next *Node[T]
}

// String returns the string representation of the node's data.
func (n *Node[T]) String() string {
	return fmt.Sprint(n.Data)
}
