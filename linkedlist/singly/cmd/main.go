package main

import (
	"fmt"

	"github.com/MehdiEidi/gods/linkedlist/singly"
)

func main() {
	s := singly.New[int]()

	s.AddFirst(4)
	s.AddFirst(3)
	s.AddFirst(2)
	s.AddFirst(1)

	fmt.Println(s)

	s2 := singly.New[int]()

	s2.AddFirst(4)
	s2.AddFirst(3)
	s2.AddFirst(2)
	s2.AddFirst(1)

	fmt.Println(s2)

	fmt.Println(s.Equals(s2))

	s3 := s2.Clone()
	s3.RemoveLast()
	fmt.Println(s3)

	s4 := singly.New[int]()
	fmt.Println(s4.ToSlice())
}
