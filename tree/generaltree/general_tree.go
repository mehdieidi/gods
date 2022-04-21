package generaltree

import (
	"fmt"
	"math"

	"github.com/MehdiEidi/gods/queue"
	"github.com/MehdiEidi/gods/stack"
)

type generalTree[T any] struct {
	root *Node[T]
	size int
}

func New[T any]() GeneralTree[T] {
	return &generalTree[T]{}
}

func (gt *generalTree[T]) Root() *Node[T] {
	return gt.root
}

func (gt *generalTree[T]) SetRoot(n *Node[T]) {
	gt.root = n
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
	return gt.root == n
}

func (gt *generalTree[T]) Size() int {
	return gt.size
}

func (gt *generalTree[T]) SetSize(s int) {
	gt.size = s
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

func (gt *generalTree[T]) AddRoot(data T) (*Node[T], error) {
	if !gt.IsEmpty() {
		return nil, TreeNotEmptyErr
	}

	gt.root = &Node[T]{Data: data}
	gt.size++

	return gt.root, nil
}

func (gt *generalTree[T]) AddChildTo(parent *Node[T], data T) *Node[T] {
	child := &Node[T]{Data: data, Parent: parent}
	parent.Children = append(parent.Children, child)

	gt.size++

	return child
}

func (gt *generalTree[T]) Set(n *Node[T], data T) T {
	val := n.Data
	n.Data = data
	return val
}

func (gt *generalTree[T]) PreOrder() (list []*Node[T]) {
	if !gt.IsEmpty() {
		gt.preOrderUtil(gt.root, &list)
	}
	return
}

func (gt *generalTree[T]) preOrderUtil(n *Node[T], list *[]*Node[T]) {
	*list = append(*list, n)

	for _, c := range gt.Children(n) {
		gt.preOrderUtil(c, list)
	}
}

func (gt *generalTree[T]) String() string {
	list := gt.PreOrder()

	str := "[ "

	for _, n := range list {
		str += fmt.Sprint(n.Data) + " "
	}

	str += "]"

	return str
}

func (gt *generalTree[T]) PostOrder() (list []*Node[T]) {
	if !gt.IsEmpty() {
		gt.postOrderUtil(gt.root, &list)
	}
	return
}

func (gt *generalTree[T]) postOrderUtil(n *Node[T], list *[]*Node[T]) {
	for _, c := range gt.Children(n) {
		gt.postOrderUtil(c, list)
	}

	*list = append(*list, n)
}

func (gt *generalTree[T]) BFS() (list []*Node[T]) {
	if gt.IsEmpty() {
		return
	}

	queue := queue.NewLinkedQueue[*Node[T]]()
	queue.Enqueue(gt.root)

	for !queue.IsEmpty() {
		n, _ := queue.Dequeue()

		list = append(list, n)

		for _, c := range gt.Children(n) {
			queue.Enqueue(c)
		}
	}

	return
}

func (gt *generalTree[T]) DFS() (list []*Node[T]) {
	if gt.IsEmpty() {
		return
	}

	stack := stack.NewLinkedStack[*Node[T]]()
	stack.Push(gt.root)

	for !stack.IsEmpty() {
		n, _ := stack.Pop()

		list = append(list, n)

		for _, c := range gt.Children(n) {
			stack.Push(c)
		}
	}

	return
}

func (gt *generalTree[T]) EulerTour() (list []*Node[T]) {
	if !gt.IsEmpty() {
		gt.eulerTourUtil(gt.root, &list)
	}
	return
}

func (gt *generalTree[T]) eulerTourUtil(n *Node[T], list *[]*Node[T]) {
	*list = append(*list, n)

	for _, c := range gt.Children(n) {
		gt.eulerTourUtil(c, list)
	}

	*list = append(*list, n)
}
