// gcoll
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)

package collection

type Collection interface {
	Size() int
	IsEmpty() bool
	Contains(ele interface{}) bool
	ToSlice() []interface{}
	Iterate() Iterator
	Add(ele interface{}) bool
	Remove(ele interface{}) bool
	AddAll(c Collection) bool
	ContainsAll(c Collection) bool
	RemoveAll(c Collection) bool
	Clear()
}

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
