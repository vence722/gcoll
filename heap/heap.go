// gcoll
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)

package heap

const (
	INIT_LEN = 0
	INIT_CAP = 16
)

// The element need to be comparable in the Heap structure
type Comparable interface {
	CompareTo(x interface{}) int
}

// The generic heap interface
type Heap interface {
	// Put
	Put(elem Comparable)
	// See the top element without taking it
	Top() interface{}
	// Take the top element
	Take() interface{}
	// The current number of elements
	Size() int
}
