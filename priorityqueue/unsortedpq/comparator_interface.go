package unsortedpq

type Comparator[T any] interface {
	Compare(T, T) int
}
