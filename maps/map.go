// gcoll v1.0
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)

package maps

const (
	INIT_LEN = 0
	INIT_CAP = 16
)

// The Map interface
type Map interface {
	Size() int
	IsEmpty() bool
	ContainsKey(key interface{}) bool
	ContainsValue(value interface{}) bool
	Get(key interface{}) interface{}
	Put(key, value interface{}) bool
	Remove(key interface{}) bool
	PutAll(amap Map) bool
	Clear()
	Keys() []interface{}
	Values() []interface{}
	Entries() []MapEntry
}

// The entry struct
type MapEntry struct {
	Key   interface{}
	Value interface{}
}
