// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package maps

const (
	INIT_LEN = 0
	INIT_CAP = 16
)

type basemap interface {
	Size() int
	IsEmpty() bool
	ContainsValue(value any) bool
	Clear()
	Values() []any
}

// The Map interface
type Map interface {
	basemap
	ContainsKey(key any) bool
	Get(key any) any
	Put(key, value any) bool
	Remove(key any) bool
	PutAll(amap Map) bool
	Keys() []any
	Entries() []MapEntry
}

// The entry struct
type MapEntry struct {
	Key   any
	Value any
}

// The StringMap interface, for Maps that need to compare keys
type StringMap interface {
	basemap
	ContainsKey(key string) bool
	Get(key string) any
	Put(key string, value any) bool
	Remove(key string) bool
	PutAll(amap StringMap) bool
	Keys() []string
	Entries() []StringMapEntry
}

// The string entry struct
type StringMapEntry struct {
	Key   string
	Value any
}
