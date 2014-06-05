// gcoll v1.0
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)

package maps

const (
	INIT_LEN = 0
	INIT_CAP = 16
)

type basemap interface {
	Size() int
	IsEmpty() bool
	ContainsValue(value interface{}) bool
	Clear()
	Values() []interface{}
}

// The Map interface
type Map interface {
	basemap
	ContainsKey(key interface{}) bool
	Get(key interface{}) interface{}
	Put(key, value interface{}) bool
	Remove(key interface{}) bool
	PutAll(amap Map) bool
	Keys() []interface{}
	Entries() []MapEntry
}

// The entry struct
type MapEntry struct {
	Key   interface{}
	Value interface{}
}

// The StringMap interface, for Maps that need to compare keys
type StringMap interface {
	basemap
	ContainsKey(key string) bool
	Get(key string) interface{}
	Put(key string, value interface{}) bool
	Remove(key string) bool
	PutAll(amap StringMap) bool
	Keys() []string
	Entries() []StringMapEntry
}

// The string entry struct
type StringMapEntry struct {
	Key   string
	Value interface{}
}
