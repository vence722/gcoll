// gcoll
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)

package list

import "github.com/vence722/gcoll/collection"

const (
	INIT_LEN = 0
	INIT_CAP = 16
)

// The list interface
type List interface {
	collection.Collection
	Get(index int) interface{}
	Set(index int, ele interface{}) bool
	Insert(index int, ele interface{}) bool
	RemoveAt(index int) interface{}
	IndexOf(ele interface{}) int
	LastIndexOf(ele interface{}) int
	SubList(fromIndex, toIndex int) List
}

// The Queue interface
type Queue interface {
	List
	EnQueue(ele interface{}) bool
	DeQueue() interface{}
	Front() interface{}
	Tail() interface{}
}

// The Stack interface
type Stack interface {
	List
	Push(ele interface{}) bool
	Pop() interface{}
	Top() interface{}
}
