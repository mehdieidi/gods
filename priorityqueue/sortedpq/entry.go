package sortedpq

import "fmt"

type Entry[K any, V any] struct {
	Key   K
	Value V
}

// String returns the string representation of the entry.
func (e *Entry[K, V]) String() string {
	return fmt.Sprintf("(%v, %v)", e.Key, e.Value)
}
