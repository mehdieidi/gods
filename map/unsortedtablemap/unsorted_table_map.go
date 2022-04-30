package unsortedtablemap

import "fmt"

type UnsortedTableMap[K comparable, V any] struct {
	Table []Entry[K, V]
}

func New[K comparable, V any]() *UnsortedTableMap[K, V] {
	return &UnsortedTableMap[K, V]{Table: []Entry[K, V]{}}
}

func (m *UnsortedTableMap[K, V]) Size() int {
	return len(m.Table)
}

func (m *UnsortedTableMap[K, V]) Get(k K) (v V, ok bool) {
	i := m.findIndex(k)
	if i == -1 {
		return
	}

	return m.Table[i].Value, true
}

func (m *UnsortedTableMap[K, V]) Put(k K, v V) (prev V) {
	i := m.findIndex(k)
	if i == -1 {
		m.Table = append(m.Table, Entry[K, V]{Key: k, Value: v})
		return
	}

	prev = m.Table[i].Value

	m.Table[i].Value = v

	return
}

func (m *UnsortedTableMap[K, V]) Remove(k K) (v V, ok bool) {
	i := m.findIndex(k)
	if i == -1 {
		return
	}

	v = m.Table[i].Value

	m.Table[i] = m.Table[len(m.Table)-1] // Move the last element to the position of the element to be removed.
	m.Table = m.Table[:len(m.Table)-1]   // Remove the last element.

	return
}

func (m *UnsortedTableMap[K, V]) String() string {
	return fmt.Sprint(m.Table)
}

func (m *UnsortedTableMap[K, V]) findIndex(k K) int {
	for i := 0; i < len(m.Table); i++ {
		if m.Table[i].Key == k {
			return i
		}
	}
	return -1
}
