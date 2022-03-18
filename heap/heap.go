// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)
package heap

const (
	INIT_LEN = 0
	INIT_CAP = 16
)

// The element need to be comparable in the Heap structure
type Comparable interface {
	CompareTo(x Comparable) int
}

// The generic heap interface
type Heap[T Comparable] interface {
	// Put
	Put(ele T)
	// See the top element without taking it
	Top() T
	// Take the top element
	Take() T
	// The current number of elements
	Size() int
	// Return if the heap has no element
	Empty() bool
}

// The PriorityQueue interface
type PriorityQueue[T Comparable] interface {
	EnQueue(ele T) bool
	DeQueue() T
	Front() T
	Size() int
}
