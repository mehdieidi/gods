package adaptablepq

type Comparator[T any] interface {
	Compare(T, T) int
}
