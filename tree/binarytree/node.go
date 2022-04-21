package binarytree

type Node[T any] struct {
	Data   T
	Parent *Node[T]
	Left   *Node[T]
	Right  *Node[T]
}
