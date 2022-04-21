package main

import (
	"fmt"
	"log"

	binarytree "github.com/MehdiEidi/gods/tree/binary"
)

func main() {
	b := binarytree.New[int]()

	r, err := b.AddRoot(1)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(b.Root().Data)

	b.AddLeft(r, 2)
	b.AddRight(r, 3)

	fmt.Println(b.Root().Data, b.Root().Left.Data, b.Root().Right.Data)

	fmt.Println(b.PostOrder())

	fmt.Println(b)
}
