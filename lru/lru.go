// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package lru

// Data entry
type entry[K comparable, V any] struct {
	key   K
	value V
}

// LRUCache LRU Cache interface
// All implementations should all support concurrent access
type LRUCache[K comparable, V any] interface {
	Put(key K, value V) bool
	Get(key K) (ele V, ok bool)
	Size() int
	Cap() int
	Clear()
}
