package binarytree

import (
	"fmt"
	"math"

	"github.com/MehdiEidi/gods/queue"
	"github.com/MehdiEidi/gods/stack"
)

type BinaryTree[T any] struct {
	Root *Node[T]
	Size int
}

// New constructs and returns an empty binary tree.
func New[T any]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

// Parent returns the parent of the given node.
func (bt *BinaryTree[T]) Parent(n *Node[T]) *Node[T] {
	return n.Parent
}

// Children returns a slice of the non-nil children of the given node.
func (bt *BinaryTree[T]) Children(n *Node[T]) (c []*Node[T]) {
	if n.Left != nil {
		c = append(c, n.Left)
	}

	if n.Right != nil {
		c = append(c, n.Right)
	}

	return
}

// ChildrenCount counts the non-nil children of the given node and returns it.
func (bt *BinaryTree[T]) ChildrenCount(n *Node[T]) (count int) {
	if n.Left != nil {
		count++
	}

	if n.Right != nil {
		count++
	}

	return
}

// IsInternal returns true if the given node has one or two children (its an internal node).
func (bt *BinaryTree[T]) IsInternal(n *Node[T]) bool {
	return bt.ChildrenCount(n) != 0
}

// IsExternal returns true if the given node doesn't have any children (its an external node).
func (bt *BinaryTree[T]) IsExternal(n *Node[T]) bool {
	return bt.ChildrenCount(n) == 0
}

// IsRoot checks if the given node is the root of the tree.
func (bt *BinaryTree[T]) IsRoot(n *Node[T]) bool {
	return n == bt.Root
}

// IsEmpty returns true if the tree is empty (doesn't have any node).
func (bt *BinaryTree[T]) IsEmpty() bool {
	return bt.Size == 0
}

// Height recursively finds the height of the given node. Height is based on the bottom of the tree (a leaf has height 0).
func (bt *BinaryTree[T]) Height(n *Node[T]) (h int) {
	for _, c := range bt.Children(n) {
		h = int(math.Max(float64(h), float64(1+bt.Height(c))))
	}
	return
}

// Depth recursively finds the depth of the given node. Depth is based on the root of the tree (root has depth 0).
func (bt *BinaryTree[T]) Depth(n *Node[T]) int {
	if bt.IsRoot(n) {
		return 0
	}

	return 1 + bt.Depth(n.Parent)
}

// LeftChild returns the left child of the given node.
func (bt *BinaryTree[T]) LeftChild(n *Node[T]) *Node[T] {
	return n.Left
}

// RightChild returns the right child of the given node.
func (bt *BinaryTree[T]) RightChild(n *Node[T]) *Node[T] {
	return n.Right
}

// Sibling returns the sibling of the given node (the other child of its parent).
func (bt *BinaryTree[T]) Sibling(n *Node[T]) *Node[T] {
	// Its root. Root doesn't have any siblings.
	if n.Parent == nil {
		return nil
	}

	if n == n.Parent.Left {
		return n.Parent.Right
	}

	return n.Parent.Left
}

// AddRoot constructs a node out of the given data and makes it the root of the tree. It returns
// TreeNotEmptyErr if tree wasn't empty before.
func (bt *BinaryTree[T]) AddRoot(data T) (*Node[T], error) {
	if !bt.IsEmpty() {
		return nil, TreeNotEmptyErr
	}

	bt.Root = &Node[T]{Data: data}
	bt.Size++

	return bt.Root, nil
}

// AddLeft constructs a node out of the given data and makes it the left child of the given node.
// It returns LeftChildExistsErr if the given node already has a left child.
func (bt *BinaryTree[T]) AddLeft(n *Node[T], data T) (*Node[T], error) {
	if n.Left != nil {
		return nil, LeftChildExistsErr
	}

	c := &Node[T]{Data: data, Parent: n}
	n.Left = c
	bt.Size++

	return c, nil
}

// AddRight constructs a node out of the given data and makes it the right child of the given node.
// It returns RightChildExistsErr if the given node already has a right child.
func (bt *BinaryTree[T]) AddRight(n *Node[T], data T) (*Node[T], error) {
	if n.Right != nil {
		return nil, RightChildExistsErr
	}

	c := &Node[T]{Data: data, Parent: n}
	n.Right = c
	bt.Size++

	return c, nil
}

// Set changes the data of the given node. It returns the previously stored data.
func (bt *BinaryTree[T]) Set(n *Node[T], data T) T {
	val := n.Data
	n.Data = data
	return val
}

// Attach gets a leaf node and two binary trees, it attaches the root of those trees to the given node as its
// children. It returns MustBeLeafErr if the given node is not a leaf.
func (bt *BinaryTree[T]) Attach(n *Node[T], t1 BinaryTree[T], t2 BinaryTree[T]) error {
	if bt.IsInternal(n) {
		return MustBeLeafErr
	}

	if !t1.IsEmpty() {
		t1.Root.Parent = n
		n.Left = t1.Root

		t1.Root = nil
	}

	if !t2.IsEmpty() {
		t2.Root.Parent = n
		n.Right = t2.Root

		t2.Root = nil
	}

	bt.Size += t1.Size + t2.Size

	return nil
}

// Remove gets a node that only has one child (left or right not both) and removes it from the tree.
// It attaches its child to its parent. It returns the removed node's data also HasTwoChildrenErr if the
// given node has two children.
func (bt *BinaryTree[T]) Remove(n *Node[T]) (val T, err error) {
	if bt.ChildrenCount(n) == 2 {
		return val, HasTwoChildrenErr
	}

	var child *Node[T]

	if n.Left != nil {
		child = n.Left
	} else {
		child = n.Right
	}

	if child != nil {
		child.Parent = n.Parent
	}

	if n == bt.Root {
		bt.Root = child
	} else {
		if n == n.Parent.Left {
			n.Parent.Left = child
		} else {
			n.Parent.Right = child
		}
	}

	bt.Size--

	val = n.Data

	// Help GC.
	var empty T
	n.Data = empty
	n.Left = nil
	n.Right = nil
	n.Parent = nil

	return
}

// String returns a string representation of the tree in pre-order.
func (bt *BinaryTree[T]) String() string {
	list := bt.PreOrder()

	str := "[ "

	for _, n := range list {
		str += fmt.Sprint(n.Data) + " "
	}

	str += "]"

	return str
}

// PreOrder returns a slice of the nodes of the tree in pre-order.
func (bt *BinaryTree[T]) PreOrder() (list []*Node[T]) {
	if !bt.IsEmpty() {
		bt.preOrderUtil(bt.Root, &list)
	}
	return
}

// preOrderUtil walks the tree in pre-order recursively and appends visited nodes to the slice.
func (bt *BinaryTree[T]) preOrderUtil(n *Node[T], list *[]*Node[T]) {
	*list = append(*list, n)

	for _, c := range bt.Children(n) {
		bt.preOrderUtil(c, list)
	}
}

// PostOrder returns a slice of the nodes of the tree in post-order.
func (bt *BinaryTree[T]) PostOrder() (list []*Node[T]) {
	if !bt.IsEmpty() {
		bt.postOrderUtil(bt.Root, &list)
	}
	return
}

// postOrderUtil walks the tree in post-order recursively and appends visited nodes to the slice.
func (bt *BinaryTree[T]) postOrderUtil(n *Node[T], list *[]*Node[T]) {
	for _, c := range bt.Children(n) {
		bt.postOrderUtil(c, list)
	}

	*list = append(*list, n)
}

// InOrder returns a slice of the nodes of the tree in in-order.
func (bt *BinaryTree[T]) InOrder() (list []*Node[T]) {
	if !bt.IsEmpty() {
		bt.inOrderUtil(bt.Root, &list)
	}
	return
}

// inOrderUtil walks the tree in in-order recursively and appends visited nodes to the slice.
func (bt *BinaryTree[T]) inOrderUtil(n *Node[T], list *[]*Node[T]) {
	if n.Left != nil {
		bt.inOrderUtil(n.Left, list)
	}

	*list = append(*list, n)

	if n.Right != nil {
		bt.inOrderUtil(n.Right, list)
	}
}

// BFS returns a slice of the nodes of the tree in breadth-first order.
func (bt *BinaryTree[T]) BFS() (list []*Node[T]) {
	if bt.IsEmpty() {
		return
	}

	queue := queue.NewLinkedQueue[*Node[T]]()
	queue.Enqueue(bt.Root)

	for !queue.IsEmpty() {
		n, _ := queue.Dequeue()

		list = append(list, n)

		for _, c := range bt.Children(n) {
			queue.Enqueue(c)
		}
	}

	return
}

// DFS returns a slice of the nodes of the tree in depth-first order.
func (bt *BinaryTree[T]) DFS() (list []*Node[T]) {
	if bt.IsEmpty() {
		return
	}

	stack := stack.NewLinkedStack[*Node[T]]()
	stack.Push(bt.Root)

	for !stack.IsEmpty() {
		n, _ := stack.Pop()

		list = append(list, n)

		for _, c := range bt.Children(n) {
			stack.Push(c)
		}
	}

	return
}

// EulerTour returns a slice of the nodes of the tree visited in the euler tour.
func (bt *BinaryTree[T]) EulerTour() (list []*Node[T]) {
	if !bt.IsEmpty() {
		bt.eulerTourUtil(bt.Root, &list)
	}
	return
}

// eulerTourUtil walks the tree in euler-tour recursively and appends visited nodes to the slice.
func (bt *BinaryTree[T]) eulerTourUtil(n *Node[T], list *[]*Node[T]) {
	*list = append(*list, n)

	if n.Left != nil {
		bt.eulerTourUtil(n.Left, list)
	}

	*list = append(*list, n)

	if n.Right != nil {
		bt.eulerTourUtil(n.Right, list)
	}

	*list = append(*list, n)
}
