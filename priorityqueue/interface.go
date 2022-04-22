package priorityqueue

type PriorityQueue[K any, V any] interface {
	Size() int
	IsEmpty() bool
	Enqueue(K, V) *Entry[K, V]
	Dequeue() *Entry[K, V]
	Min() *Entry[K, V]
}

type Comparator[T any] interface {
	Compare(T, T) int
}
