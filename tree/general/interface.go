package tree

type Tree[T comparable] interface {
	Root() *Node[T]
	SetRoot(*Node[T])
	Size() int
	SetSize(int)
	Parent(*Node[T]) *Node[T]
	Children(*Node[T]) []*Node[T]
	ChildrenNum(*Node[T]) int
	IsInternal(*Node[T]) bool
	IsExternal(*Node[T]) bool
	IsRoot(*Node[T]) bool
	IsEmpty() bool
	Height(*Node[T]) int
	Depth(*Node[T]) int
	AddRoot(T) (*Node[T], error)
	AddChildTo(*Node[T], T) *Node[T]
	Set(*Node[T], T) T
	String() string
	PreOrder() []*Node[T]
	PostOrder() []*Node[T]
	BFS() []*Node[T]
	DFS() []*Node[T]
	EulerTour() []*Node[T]
	Parenthesize(*Node[T])
}
