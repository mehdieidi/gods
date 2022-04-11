package main

import (
	"fmt"

	singlylinkedlist "github.com/MehdiEidi/gods/linkedlist/singly"
)

func main() {
	s := singlylinkedlist.New()

	fmt.Println(s.Head)
	fmt.Println(s.Tail)
	fmt.Println(s.Size)
	fmt.Println(s.IsEmpty())

	s.AddFirst(4)
	s.AddFirst(3)
	s.AddFirst(2)
	s.AddFirst(1)

	fmt.Println(s)

	fmt.Println(s.Head.Data)
	fmt.Println(s.Tail.Data)
	fmt.Println(s.Size)
	fmt.Println(s.IsEmpty())

	s.AddLast(6)
	s.AddLast(7)

	fmt.Println(s)

	fmt.Println(s.Head.Data)
	fmt.Println(s.Tail.Data)
	fmt.Println(s.Size)
	fmt.Println(s.IsEmpty())

	s.Add(99, 1)

	fmt.Println(s)

	s.Remove(0)

	fmt.Println(s)

	fmt.Println(s.Head.Data)
	fmt.Println(s.Tail.Data)
	fmt.Println(s.Size)
	fmt.Println(s.IsEmpty())

	// s.RemoveLast()
	// s.RemoveLast()
	// s.RemoveLast()
	// s.RemoveLast()
	// s.RemoveLast()
	// s.RemoveLast()

	// fmt.Println(s)

	// fmt.Println(s.Head.Data)
	// fmt.Println(s.Tail.Data)
	// fmt.Println(s.Size)
	// fmt.Println(s.IsEmpty())
}
