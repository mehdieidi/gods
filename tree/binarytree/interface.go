package binarytree

type BinaryTree[T any] interface {
	Root() *Node[T]
	SetRoot(*Node[T])
	IsRoot(*Node[T]) bool
	Size() int
	SetSize(int)
	IsEmpty() bool
	Parent(*Node[T]) *Node[T]
	Children(*Node[T]) []*Node[T]
	ChildrenNum(*Node[T]) int
	IsInternal(*Node[T]) bool
	IsExternal(*Node[T]) bool
	Height(*Node[T]) int
	Depth(*Node[T]) int
	LeftChild(*Node[T]) *Node[T]
	RightChild(*Node[T]) *Node[T]
	Sibling(*Node[T]) *Node[T]
	AddRoot(T) (*Node[T], error)
	AddLeft(*Node[T], T) (*Node[T], error)
	AddRight(*Node[T], T) (*Node[T], error)
	Set(*Node[T], T) T
	Attach(*Node[T], BinaryTree[T], BinaryTree[T]) error
	Remove(*Node[T]) (T, error)
	String() string
	PreOrder() []*Node[T]
	PostOrder() []*Node[T]
	InOrder() []*Node[T]
	BFS() []*Node[T]
	DFS() []*Node[T]
	EulerTour() []*Node[T]
}
