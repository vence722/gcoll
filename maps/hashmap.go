// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package maps

import (
	"fmt"
)

// The HashMap struct
type HashMap[K comparable, V any] struct {
	hmap map[K]V
}

// NewHashMap Returns a new HashMap
func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	hmap := make(map[K]V)
	return &HashMap[K, V]{hmap}
}

// Size Returns the size of the map
func (m *HashMap[K, V]) Size() int {
	return len(m.hmap)
}

// IsEmpty Returns the map containing elements or not
func (m *HashMap[K, V]) IsEmpty() bool {
	return len(m.hmap) == 0
}

// ContainsKey Returns the map conaining specified key or not
func (m *HashMap[K, V]) ContainsKey(key K) bool {
	for k, _ := range m.hmap {
		if k == key {
			return true
		}
	}
	return false
}

// Get Returns the element with the specified key in this map
func (m *HashMap[K, V]) Get(key K) (ele V, ok bool) {
	v, ok := m.hmap[key]
	return v, ok
}

// Put Adds new key value pair to this map
func (m *HashMap[K, V]) Put(key K, value V) bool {
	m.hmap[key] = value
	return true
}

// Remove Removes value with specified key from this map
func (m *HashMap[K, V]) Remove(key K) bool {
	delete(m.hmap, key)
	return true
}

// PutAll Adds another map to this map
func (m *HashMap[K, V]) PutAll(amap Map[K, V]) bool {
	entries := amap.Entries()
	for _, entry := range entries {
		m.Put(entry.Key, entry.Value)
	}
	return true
}

// Clear Removes all elements from this map
func (m *HashMap[K, V]) Clear() {
	for k, _ := range m.hmap {
		delete(m.hmap, k)
	}
}

// Keys Returns a copy of all the keys in the map
func (m *HashMap[K, V]) Keys() []K {
	var keys []K
	for k, _ := range m.hmap {
		keys = append(keys, k)
	}
	return keys
}

// Values Returns a copy of all the values in the map
func (m *HashMap[K, V]) Values() []V {
	var values []V
	for _, v := range m.hmap {
		values = append(values, v)
	}
	return values
}

// Entries Returns a copy of all the entries in the map
func (m *HashMap[K, V]) Entries() []MapEntry[K, V] {
	var entries []MapEntry[K, V]
	for k, v := range m.hmap {
		entries = append(entries, MapEntry[K, V]{k, v})
	}
	return entries
}

// String Returns the string that describes the contents of this map
func (m *HashMap[K, V]) String() string {
	return fmt.Sprint(m.hmap)
}
