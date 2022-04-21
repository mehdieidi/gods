package tree

type Tree[T comparable] interface {
	Root() *Node[T]
	Parent(*Node[T]) *Node[T]
	Children(*Node[T]) []*Node[T]
	ChildrenNum(*Node[T]) int
	IsInternal(*Node[T]) bool
	IsExternal(*Node[T]) bool
	IsRoot(*Node[T]) bool
	Size() int
	IsEmpty() bool
	Height(*Node[T]) int
	Depth(*Node[T]) int
}
