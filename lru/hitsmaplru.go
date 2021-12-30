// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package lru

import (
	"sync"

	"github.com/vence722/gcoll/maps"
)

// Data entry with hits counter
type hitsEntry[K comparable, V any] struct {
	entry[K, V]
	hits int
}

// HitsMapLRUCache LRU Cache implementation of Hits Map
type HitsMapLRUCache[K comparable, V any] struct {
	hitsMap  maps.Map[K, *hitsEntry[K, V]]
	capacity int
	mutex    sync.Mutex
}

func NewHitsMapLRUCache[K comparable, V any](capacity int) *HitsMapLRUCache[K, V] {
	return &HitsMapLRUCache[K, V]{
		hitsMap:  maps.NewHashMap[K, *hitsEntry[K, V]](),
		capacity: capacity,
	}
}

// Put Puts new element into the LRU cache.
// If key already exists, update the value.
// If the key is not found in the LRU cache,
// and the LRU cache is not full, add a new entry,
// otherwise remove the entry in the HitsMap and
// clear hits count before inserting the new element.
func (hm *HitsMapLRUCache[K, V]) Put(key K, value V) bool {
	hm.mutex.Lock()
	defer hm.mutex.Unlock()
	ele, ok := hm.hitsMap.Get(key)
	if ok {
		ele.value = value
		// clear hits count
		ele.hits = 0
	} else {
		// eliminate out-of-date entry if needed
		hm.eliminate()
		// Insert new entry
		ne := &hitsEntry[K, V]{}
		ne.key = key
		ne.value = value
		ne.hits = 1
		hm.hitsMap.Put(ne.key, ne)
	}
	return true
}

func (hm *HitsMapLRUCache[K, V]) Get(key K) (V, bool) {
	hm.mutex.Lock()
	defer hm.mutex.Unlock()
	ele, ok := hm.hitsMap.Get(key)
	if ok {
		// update hits counter
		ele.hits++
		return ele.value, ok
	}
	var zero V
	return zero, false
}

func (hm *HitsMapLRUCache[K, V]) Size() int {
	hm.mutex.Lock()
	defer hm.mutex.Unlock()
	return hm.hitsMap.Size()
}

// Cap Returns the capacity of the LRU Cache
func (hm *HitsMapLRUCache[K, V]) Cap() int {
	return hm.capacity
}

// Clear Clears the LRU Cache
func (hm *HitsMapLRUCache[K, V]) Clear() {
	hm.mutex.Lock()
	defer hm.mutex.Unlock()
	hm.hitsMap.Clear()
}

func (hm *HitsMapLRUCache[K, V]) eliminate() {
	if hm.hitsMap.Size() == hm.capacity {
		// LRU Cache is full, remove the one of the least used entry
		// from the cache
		var leastUsed *hitsEntry[K, V]
		for _, ele := range hm.hitsMap.Entries() {
			if leastUsed != nil {
				he := ele.Value
				if he.hits < leastUsed.hits {
					leastUsed = ele.Value
				}
			} else {
				leastUsed = ele.Value
			}
		}
		// Remove the least used entry
		if leastUsed != nil {
			hm.hitsMap.Remove(leastUsed.key)
		}
	}
}
