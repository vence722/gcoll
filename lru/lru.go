// gcoll
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)

package lru

// Data entry
type entry struct {
	key   interface{}
	value interface{}
}

// LRU Cache interface
// All implementations should all support concurrent access
type LRUCache interface {
	Put(key interface{}, value interface{})
	Get(key interface{}) interface{}
	Size() int
	Cap() int
	Clear()
}
