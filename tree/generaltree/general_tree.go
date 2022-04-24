package generaltree

import (
	"fmt"
	"math"

	"github.com/MehdiEidi/gods/queue/linkedqueue"
	"github.com/MehdiEidi/gods/stack/linkedstack"
)

type GeneralTree[T any] struct {
	Root *Node[T]
	Size int
}

// New constructs and returns an empty general tree.
func New[T any]() *GeneralTree[T] {
	return &GeneralTree[T]{}
}

// Parent returns the parent of the given node.
func (gt *GeneralTree[T]) Parent(n *Node[T]) *Node[T] {
	return n.Parent
}

// Children returns a slice of the children of the given node.
func (gt *GeneralTree[T]) Children(n *Node[T]) []*Node[T] {
	return n.Children
}

// ChildrenCount returns the count of the children of the given node.
func (gt *GeneralTree[T]) ChildrenCount(n *Node[T]) int {
	return len(n.Children)
}

// IsInternal returns true if the given node has one or more children (its an internal node).
func (gt *GeneralTree[T]) IsInternal(n *Node[T]) bool {
	return gt.ChildrenCount(n) != 0
}

// IsExternal returns true if the given node doesn't have any children (its an external node).
func (gt *GeneralTree[T]) IsExternal(n *Node[T]) bool {
	return gt.ChildrenCount(n) == 0
}

// IsRoot checks if the given node is the root of the tree.
func (gt *GeneralTree[T]) IsRoot(n *Node[T]) bool {
	return gt.Root == n
}

// IsEmpty returns true if the tree is empty (doesn't have any node).
func (gt *GeneralTree[T]) IsEmpty() bool {
	return gt.Size == 0
}

// Depth recursively finds the depth of the given node. Depth is based on the root of the tree (root has depth 0).
func (gt *GeneralTree[T]) Depth(n *Node[T]) int {
	if gt.IsRoot(n) {
		return 0
	}

	return 1 + gt.Depth(n.Parent)
}

// Height recursively finds the height of the given node. Height is based on the bottom of the tree (a leaf has height 0).
func (gt *GeneralTree[T]) Height(n *Node[T]) (h int) {
	for _, c := range n.Children {
		h = int(math.Max(float64(h), float64(1+gt.Height(c))))
	}
	return
}

// AddRoot constructs a node out of the given data and makes it the root of the tree. It returns
// TreeNotEmptyErr if tree wasn't empty before.
func (gt *GeneralTree[T]) AddRoot(data T) (*Node[T], error) {
	if !gt.IsEmpty() {
		return nil, TreeNotEmptyErr
	}

	gt.Root = &Node[T]{Data: data}
	gt.Size++

	return gt.Root, nil
}

// AddChildTo constructs a node out of the given data and adds it to the children of the given node.
// It returns the newly constructed node.
func (gt *GeneralTree[T]) AddChildTo(parent *Node[T], data T) *Node[T] {
	child := &Node[T]{Data: data, Parent: parent}
	parent.Children = append(parent.Children, child)

	gt.Size++

	return child
}

// Set changes the data of the given node. It returns the previously stored data.
func (gt *GeneralTree[T]) Set(n *Node[T], data T) T {
	val := n.Data
	n.Data = data
	return val
}

// PreOrder returns a slice of the nodes of the tree in pre-order.
func (gt *GeneralTree[T]) PreOrder() (list []*Node[T]) {
	if !gt.IsEmpty() {
		gt.preOrderUtil(gt.Root, &list)
	}
	return
}

// preOrderUtil walks the tree in pre-order recursively and appends visited nodes to the slice.
func (gt *GeneralTree[T]) preOrderUtil(n *Node[T], list *[]*Node[T]) {
	*list = append(*list, n)

	for _, c := range gt.Children(n) {
		gt.preOrderUtil(c, list)
	}
}

// String returns a string representation of the tree in pre-order.
func (gt *GeneralTree[T]) String() string {
	list := gt.PreOrder()

	str := "[ "

	for _, n := range list {
		str += fmt.Sprint(n.Data) + " "
	}

	str += "]"

	return str
}

// PostOrder returns a slice of the nodes of the tree in post-order.
func (gt *GeneralTree[T]) PostOrder() (list []*Node[T]) {
	if !gt.IsEmpty() {
		gt.postOrderUtil(gt.Root, &list)
	}
	return
}

// postOrderUtil walks the tree in post-order recursively and appends visited nodes to the slice.
func (gt *GeneralTree[T]) postOrderUtil(n *Node[T], list *[]*Node[T]) {
	for _, c := range gt.Children(n) {
		gt.postOrderUtil(c, list)
	}

	*list = append(*list, n)
}

// BFS returns a slice of the nodes of the tree in breadth-first order.
func (gt *GeneralTree[T]) BFS() (list []*Node[T]) {
	if gt.IsEmpty() {
		return
	}

	queue := linkedqueue.New[*Node[T]]()
	queue.Enqueue(gt.Root)

	for !queue.IsEmpty() {
		n, _ := queue.Dequeue()

		list = append(list, n)

		for _, c := range gt.Children(n) {
			queue.Enqueue(c)
		}
	}

	return
}

// DFS returns a slice of the nodes of the tree in depth-first order.
func (gt *GeneralTree[T]) DFS() (list []*Node[T]) {
	if gt.IsEmpty() {
		return
	}

	stack := linkedstack.New[*Node[T]]()
	stack.Push(gt.Root)

	for !stack.IsEmpty() {
		n, _ := stack.Pop()

		list = append(list, n)

		for _, c := range gt.Children(n) {
			stack.Push(c)
		}
	}

	return
}

// EulerTour returns a slice of the nodes of the tree visited in the euler tour.
func (gt *GeneralTree[T]) EulerTour() (list []*Node[T]) {
	if !gt.IsEmpty() {
		gt.eulerTourUtil(gt.Root, &list)
	}
	return
}

// eulerTourUtil walks the tree in euler-tour recursively and appends visited nodes to the slice.
func (gt *GeneralTree[T]) eulerTourUtil(n *Node[T], list *[]*Node[T]) {
	*list = append(*list, n)

	for _, c := range gt.Children(n) {
		gt.eulerTourUtil(c, list)
	}

	*list = append(*list, n)
}
