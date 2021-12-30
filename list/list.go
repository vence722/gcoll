// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package list

import (
	"github.com/vence722/gcoll/collection"
)

const (
	InitLen = 0
	InitCap = 16
)

// List The list interface
type List[T comparable] interface {
	collection.Collection[T]
	Get(index int) (ele T, ok bool)
	MustGet(index int) T
	Set(index int, ele T) bool
	Insert(index int, ele T) bool
	RemoveAt(index int) (ele T, ok bool)
	IndexOf(ele T) int
	LastIndexOf(ele T) int
	SubList(fromIndex, toIndex int) List[T]
}

// The Queue interface
type Queue[T comparable] interface {
	List[T]
	EnQueue(ele T) bool
	DeQueue() (ele T, ok bool)
	Front() (ele T, ok bool)
	Tail() (ele T, ok bool)
}

// The Stack interface
type Stack[T comparable] interface {
	List[T]
	Push(ele T) bool
	Pop() (ele T, ok bool)
	Top() (ele T, ok bool)
}
