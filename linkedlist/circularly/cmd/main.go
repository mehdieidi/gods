package main

import (
	"fmt"

	"github.com/MehdiEidi/gods/linkedlist/circularly"
)

func main() {
	c := circularly.New()

	fmt.Println(c.Size)
	fmt.Println(c.First())
	fmt.Println(c.Last())
	fmt.Println(c.IsEmpty())

	fmt.Println(c)

	c.AddFirst(4)
	c.AddFirst(3)
	c.AddFirst(2)
	c.AddFirst(1)
	c.AddFirst(0)

	fmt.Println(c)

	fmt.Println(c.Size)
	fmt.Println(c.First())
	fmt.Println(c.Last())
	fmt.Println(c.IsEmpty())

	c.Rotate()

	fmt.Println(c)

	fmt.Println(c.Size)
	fmt.Println(c.First())
	fmt.Println(c.Last())
	fmt.Println(c.IsEmpty())

	c.AddLast(0)
	c.AddLast(1)
	c.AddLast(2)
	c.AddLast(3)
	c.AddLast(4)

	fmt.Println(c)

	fmt.Println(c.Size)
	fmt.Println(c.First())
	fmt.Println(c.Last())
	fmt.Println(c.IsEmpty())

	c.RemoveFirst()
	c.RemoveFirst()

	fmt.Println(c)

	fmt.Println(c.Size)
	fmt.Println(c.First())
	fmt.Println(c.Last())
	fmt.Println(c.IsEmpty())
}
