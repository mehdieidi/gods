package singly

type Node[T any] struct {
	Data T
	Next *Node
}
