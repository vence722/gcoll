// gcoll
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)

package lru

import (
	"sync"

	"github.com/vence722/gcoll/maps"
)

// Data entry with hits counter
type hitsEntry struct {
	entry
	hits int
}

// LRU Cache implementation of Hits Map
type HitsMapLRUCache struct {
	HitsMap maps.Map
	Cap     int
	Mutex   sync.Mutex
}

// Constructor
func NewHitsMapLRUCache(capacity int) *HitsMapLRUCache {
	return &HitsMapLRUCache{
		HitsMap: maps.NewHashMap(),
		Cap:     capacity,
	}
}

// Put new element into the LRU cache.
// If key alrady exists, update the value.
// If the key is not found in the LRU cache,
// and the LRU cache is not full, add a new entry.
// Otherwise remove the entry in the HitsMap and
// clear hits count before inserting the new element.
func (this *HitsMapLRUCache) Put(key interface{}, value interface{}) {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()
	elem := this.HitsMap.Get(key)
	if elem != nil {
		ent := elem.(*hitsEntry)
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
		this.HitsMap.Put(ne.key, ne)
	}
}

func (this *HitsMapLRUCache) Get(key interface{}) interface{} {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()
	elem := this.HitsMap.Get(key)
	if elem != nil {
		ent := elem.(*hitsEntry)
		// update hits counter
		ent.hits++
		return ent.value
	} else {
		return nil
	}
}

func (this *HitsMapLRUCache) Size() int {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()
	return this.HitsMap.Size()
}

// Clear the LRU Cache
func (this *HitsMapLRUCache) Clear() {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()
	this.HitsMap.Clear()
}

func (this *HitsMapLRUCache) eliminate() {
	if this.HitsMap.Size() == this.Cap {
		// LRU Cache is full, remove the one of the least used entry
		// from the cache
		var leastUsed *hitsEntry
		for _, elem := range this.HitsMap.Entries() {
			if leastUsed != nil {
				he := elem.Value.(*hitsEntry)
				if he.hits < leastUsed.hits {
					leastUsed = elem.Value.(*hitsEntry)
				}
			} else {
				leastUsed = elem.Value.(*hitsEntry)
			}
		}
		// Remove the least used entry
		if leastUsed != nil {
			this.HitsMap.Remove(leastUsed.key)
		}
	}
}
