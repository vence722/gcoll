// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package collection

type Collection[T comparable] interface {
	Size() int
	IsEmpty() bool
	Contains(ele T) bool
	ToSlice() []T
	Iterate() Iterator[T]
	Add(ele T) bool
	Remove(ele T) bool
	AddAll(c Collection[T]) bool
	ContainsAll(c Collection[T]) bool
	RemoveAll(c Collection[T]) bool
	Clear()
}

type Iterator[T comparable] interface {
	HasNext() bool
	Next() T
}
