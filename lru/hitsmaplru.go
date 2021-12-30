// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package lru

import (
	"sync"

	"gcoll/maps"
)

// Data entry with hits counter
type hitsEntry struct {
	entry
	hits int
}

// LRU Cache implementation of Hits Map
type HitsMapLRUCache struct {
	hitsMap  maps.Map
	capacity int
	mutex    sync.Mutex
}

// Constructor
func NewHitsMapLRUCache(capacity int) *HitsMapLRUCache {
	return &HitsMapLRUCache{
		hitsMap:  maps.NewHashMap(),
		capacity: capacity,
	}
}

// Put new element into the LRU cache.
// If key alrady exists, update the value.
// If the key is not found in the LRU cache,
// and the LRU cache is not full, add a new entry.
// Otherwise remove the entry in the HitsMap and
// clear hits count before inserting the new element.
func (this *HitsMapLRUCache) Put(key any, value any) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	ele := this.hitsMap.Get(key)
	if ele != nil {
		ent := ele.(*hitsEntry)
		ent.value = value
		// clear hits count
		ent.hits = 0
	} else {
		// eliminate out-of-date entry if needed
		this.eliminate()
		// Insert new entry
		ne := &hitsEntry{}
		ne.key = key
		ne.value = value
		ne.hits = 1
		this.hitsMap.Put(ne.key, ne)
	}
}

func (this *HitsMapLRUCache) Get(key any) any {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	ele := this.hitsMap.Get(key)
	if ele != nil {
		ent := ele.(*hitsEntry)
		// update hits counter
		ent.hits++
		return ent.value
	} else {
		return nil
	}
}

func (this *HitsMapLRUCache) Size() int {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return this.hitsMap.Size()
}

// Returns the capacity of the LRU Cache
func (this *HitsMapLRUCache) Cap() int {
	return this.capacity
}

// Clear the LRU Cache
func (this *HitsMapLRUCache) Clear() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.hitsMap.Clear()
}

func (this *HitsMapLRUCache) eliminate() {
	if this.hitsMap.Size() == this.capacity {
		// LRU Cache is full, remove the one of the least used entry
		// from the cache
		var leastUsed *hitsEntry
		for _, ele := range this.hitsMap.Entries() {
			if leastUsed != nil {
				he := ele.Value.(*hitsEntry)
				if he.hits < leastUsed.hits {
					leastUsed = ele.Value.(*hitsEntry)
				}
			} else {
				leastUsed = ele.Value.(*hitsEntry)
			}
		}
		// Remove the least used entry
		if leastUsed != nil {
			this.hitsMap.Remove(leastUsed.key)
		}
	}
}
