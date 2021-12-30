// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package lru

// Data entry
type entry struct {
	key   any
	value any
}

// LRU Cache interface
// All implementations should all support concurrent access
type LRUCache interface {
	Put(key any, value any)
	Get(key any) any
	Size() int
	Cap() int
	Clear()
}
