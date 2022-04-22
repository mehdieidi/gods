package priorityqueue

type Entry[K any, V any] struct {
	Key   K
	Value V
}
