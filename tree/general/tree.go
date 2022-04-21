package tree

import "math"

type generalTree[T comparable] struct {
	root *Node[T]
	size int
}

func New[T comparable]() Tree[T] {
	return &generalTree[T]{}
}

func (gt *generalTree[T]) Root() *Node[T] {
	return gt.root
}

func (gt *generalTree[T]) Parent(n *Node[T]) *Node[T] {
	return n.Parent
}

func (gt *generalTree[T]) Children(n *Node[T]) []*Node[T] {
	return n.Children
}

func (gt *generalTree[T]) ChildrenNum(n *Node[T]) int {
	return len(n.Children)
}

func (gt *generalTree[T]) IsInternal(n *Node[T]) bool {
	return gt.ChildrenNum(n) != 0
}

func (gt *generalTree[T]) IsExternal(n *Node[T]) bool {
	return gt.ChildrenNum(n) == 0
}

func (gt *generalTree[T]) IsRoot(n *Node[T]) bool {
	return gt.Root() == n
}

func (gt *generalTree[T]) Size() int {
	return gt.size
}

func (gt *generalTree[T]) IsEmpty() bool {
	return gt.size == 0
}

func (gt *generalTree[T]) Depth(n *Node[T]) int {
	if gt.IsRoot(n) {
		return 0
	}

	return 1 + gt.Depth(n.Parent)
}

func (gt *generalTree[T]) Height(n *Node[T]) (h int) {
	for _, c := range n.Children {
		h = int(math.Max(float64(h), 1.0+float64(gt.Height(c))))
	}
	return
}
