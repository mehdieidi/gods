package binarytree

import (
	"fmt"
	"math"

	"github.com/MehdiEidi/gods/queue"
	"github.com/MehdiEidi/gods/stack"
)

type binaryTree[T any] struct {
	root *Node[T]
	size int
}

func New[T any]() BinaryTree[T] {
	return &binaryTree[T]{}
}

func (bt *binaryTree[T]) Parent(n *Node[T]) *Node[T] {
	return n.Parent
}

func (bt *binaryTree[T]) Children(n *Node[T]) (c []*Node[T]) {
	if n.Left != nil {
		c = append(c, n.Left)
	}

	if n.Right != nil {
		c = append(c, n.Right)
	}

	return
}

func (bt *binaryTree[T]) ChildrenNum(n *Node[T]) (count int) {
	if n.Left != nil {
		count++
	}

	if n.Right != nil {
		count++
	}

	return
}

func (bt *binaryTree[T]) IsInternal(n *Node[T]) bool {
	return bt.ChildrenNum(n) != 0
}

func (bt *binaryTree[T]) IsExternal(n *Node[T]) bool {
	return bt.ChildrenNum(n) == 0
}

func (bt *binaryTree[T]) Root() *Node[T] {
	return bt.root
}

func (bt *binaryTree[T]) SetRoot(n *Node[T]) {
	bt.root = n
}

func (bt *binaryTree[T]) Size() int {
	return bt.size
}

func (bt *binaryTree[T]) SetSize(s int) {
	bt.size = s
}

func (bt *binaryTree[T]) IsRoot(n *Node[T]) bool {
	return n == bt.root
}

func (bt *binaryTree[T]) IsEmpty() bool {
	return bt.size == 0
}

func (bt *binaryTree[T]) Height(n *Node[T]) (h int) {
	for _, c := range bt.Children(n) {
		h = int(math.Max(float64(h), 1.0+float64(bt.Height(c))))
	}
	return
}

func (bt *binaryTree[T]) Depth(n *Node[T]) int {
	if bt.IsRoot(n) {
		return 0
	}

	return 1 + bt.Depth(n.Parent)
}

func (bt *binaryTree[T]) LeftChild(n *Node[T]) *Node[T] {
	return n.Left
}

func (bt *binaryTree[T]) RightChild(n *Node[T]) *Node[T] {
	return n.Right
}

func (bt *binaryTree[T]) Sibling(n *Node[T]) *Node[T] {
	// Its root.
	if n.Parent == nil {
		return nil
	}

	if n == n.Parent.Left {
		return n.Parent.Right
	}

	return n.Parent.Left
}

func (bt *binaryTree[T]) AddRoot(data T) (*Node[T], error) {
	if !bt.IsEmpty() {
		return nil, TreeNotEmptyErr
	}

	bt.root = &Node[T]{Data: data}
	bt.size++

	return bt.root, nil
}

func (bt *binaryTree[T]) AddLeft(n *Node[T], data T) (*Node[T], error) {
	if n.Left != nil {
		return nil, LeftChildExistsErr
	}

	c := &Node[T]{Data: data, Parent: n}
	n.Left = c
	bt.size++

	return c, nil
}

func (bt *binaryTree[T]) AddRight(n *Node[T], data T) (*Node[T], error) {
	if n.Right != nil {
		return nil, RightChildExistsErr
	}

	c := &Node[T]{Data: data, Parent: n}
	n.Right = c
	bt.size++

	return c, nil
}

func (bt *binaryTree[T]) Set(n *Node[T], data T) T {
	val := n.Data
	n.Data = data
	return val
}

func (bt *binaryTree[T]) Attach(n *Node[T], t1 BinaryTree[T], t2 BinaryTree[T]) error {
	if bt.IsInternal(n) {
		return MustBeLeafErr
	}

	bt.size += t1.Size() + t2.Size()

	if !t1.IsEmpty() {
		t1.Root().Parent = n
		n.Left = t1.Root()

		t1.SetRoot(nil)
	}

	if !t2.IsEmpty() {
		t2.Root().Parent = n
		n.Right = t2.Root()

		t2.SetRoot(nil)
	}

	return nil
}

func (bt *binaryTree[T]) Remove(n *Node[T]) (val T, err error) {
	if bt.ChildrenNum(n) == 2 {
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

	if n == bt.root {
		bt.root = child
	} else {
		if n == n.Parent.Left {
			n.Parent.Left = child
		} else {
			n.Parent.Right = child
		}
	}

	bt.size--

	val = n.Data

	// Help GC.
	var empty T
	n.Data = empty
	n.Left = nil
	n.Right = nil
	n.Parent = nil

	return
}

func (bt *binaryTree[T]) String() string {
	list := bt.PreOrder()

	str := "[ "

	for _, n := range list {
		str += fmt.Sprint(n.Data) + " "
	}

	str += "]"

	return str
}

func (bt *binaryTree[T]) PreOrder() (list []*Node[T]) {
	if !bt.IsEmpty() {
		bt.preOrderUtil(bt.root, &list)
	}
	return
}

func (bt *binaryTree[T]) preOrderUtil(n *Node[T], list *[]*Node[T]) {
	*list = append(*list, n)

	for _, c := range bt.Children(n) {
		bt.preOrderUtil(c, list)
	}
}

func (bt *binaryTree[T]) PostOrder() (list []*Node[T]) {
	if !bt.IsEmpty() {
		bt.postOrderUtil(bt.root, &list)
	}
	return
}

func (bt *binaryTree[T]) postOrderUtil(n *Node[T], list *[]*Node[T]) {
	for _, c := range bt.Children(n) {
		bt.postOrderUtil(c, list)
	}

	*list = append(*list, n)
}

func (bt *binaryTree[T]) InOrder() (list []*Node[T]) {
	if !bt.IsEmpty() {
		bt.inOrderUtil(bt.root, &list)
	}
	return
}

func (bt *binaryTree[T]) inOrderUtil(n *Node[T], list *[]*Node[T]) {
	if n.Left != nil {
		bt.inOrderUtil(n.Left, list)
	}

	*list = append(*list, n)

	if n.Right != nil {
		bt.inOrderUtil(n.Right, list)
	}
}

func (bt *binaryTree[T]) BFS() (list []*Node[T]) {
	if bt.IsEmpty() {
		return
	}

	queue := queue.NewLinkedQueue[*Node[T]]()
	queue.Enqueue(bt.root)

	for !queue.IsEmpty() {
		n, _ := queue.Dequeue()

		list = append(list, n)

		for _, c := range bt.Children(n) {
			queue.Enqueue(c)
		}
	}

	return
}

func (bt *binaryTree[T]) DFS() (list []*Node[T]) {
	if bt.IsEmpty() {
		return
	}

	stack := stack.NewLinkedStack[*Node[T]]()
	stack.Push(bt.root)

	for !stack.IsEmpty() {
		n, _ := stack.Pop()

		list = append(list, n)

		for _, c := range bt.Children(n) {
			stack.Push(c)
		}
	}

	return
}

func (bt *binaryTree[T]) EulerTour() (list []*Node[T]) {
	if !bt.IsEmpty() {
		bt.eulerTourUtil(bt.root, &list)
	}
	return
}

func (bt *binaryTree[T]) eulerTourUtil(n *Node[T], list *[]*Node[T]) {
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
