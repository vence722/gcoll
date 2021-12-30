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
	CompareTo(x any) int
}

// The generic heap interface
type Heap interface {
	// Put
	Put(ele Comparable)
	// See the top element without taking it
	Top() Comparable
	// Take the top element
	Take() Comparable
	// The current number of elements
	Size() int
	// Return if the heap has no element
	Empty() bool
}

// The PriorityQueue interface
type PriorityQueue interface {
	EnQueue(ele Comparable) bool
	DeQueue() Comparable
	Front() Comparable
	Size() int
}
