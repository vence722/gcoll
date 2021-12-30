// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package maps

const (
	INIT_LEN = 0
	INIT_CAP = 16
)

// The Map interface
type Map[K comparable, V any] interface {
	Size() int
	IsEmpty() bool
	Clear()
	Keys() []K
	Values() []V
	Entries() []MapEntry[K, V]
	ContainsKey(key K) bool
	Get(key K) (ele V, ok bool)
	Put(key K, value V) bool
	Remove(key K) bool
	PutAll(m Map[K, V]) bool
}

// The MapEntry struct
type MapEntry[K comparable, V any] struct {
	Key   K
	Value V
}

type SyncMap[K comparable, V any] interface {
	Store(key K, value V)
	Load(key K) (V, bool)
	LoadOrStore(key K, value V) (V, bool)
	LoadAndDelete(key K) (V, bool)
	Delete(key K)
	Range(f func(key K, value V) bool)
}
