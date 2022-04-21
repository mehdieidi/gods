package generaltree

type Node[T any] struct {
	Data     T
	Parent   *Node[T]
	Children []*Node[T]
}
