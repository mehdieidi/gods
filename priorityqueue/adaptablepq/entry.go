package adaptablepq

type Entry[K, V any] struct {
	Index int
	Key   K
	Value V
}
