package heappq

type Comparator[T any] interface {
	Compare(T, T) int
}
