package main

import (
	"fmt"

	"github.com/MehdiEidi/gods/priorityqueue/sortedpq"
	"golang.org/x/exp/constraints"
)

type comparator[k constraints.Ordered] struct {
}

func (c *comparator[k]) Compare(k1 k, k2 k) int {
	if k1 < k2 {
		return -1
	}
	if k1 > k2 {
		return 1
	}
	return 0
}
func main() {
	pq := sortedpq.New[int, string](&comparator[int]{})

	pq.Insert(5, "hello")
	pq.Insert(19, "hi")
	pq.Insert(1, "dude")

	fmt.Println(pq.Min())
	fmt.Println(pq)
}
