package main

import (
	"fmt"

	"github.com/MehdiEidi/gods/linkedlist/doubly"
)

func main() {
	d := doubly.New()

	d.AddLast(1)
	d.AddLast(2)
	d.AddLast(3)
	d.AddLast(4)

	fmt.Println(d.Size)
	fmt.Println(d.First())
	fmt.Println(d.Last())

	fmt.Println(d)

	d.RemoveFirst()
	d.RemoveFirst()

	fmt.Println(d.Size)
	fmt.Println(d.First())
	fmt.Println(d.Last())

	fmt.Println(d)
}
