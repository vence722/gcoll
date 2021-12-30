// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)
package lru

import (
	"sync"

	"gcoll/list"
)

// LRU Cache implementation of Fifo Queue
type FifoLRUCache struct {
	fifoQueue list.Queue
	capacity  int
	mutex     sync.Mutex
}

// Constructor
func NewFifoLRUCache(capacity int) *FifoLRUCache {
	return &FifoLRUCache{
		fifoQueue: list.NewLinkedList(),
		capacity:  capacity,
	}
}

// Put new element into the LRU cache.
// If key alrady exists, update the value.
// If the key is not found in the LRU cache,
// and the LRU cache is not full, add a new entry.
// Otherwise remove the tail of the FIFO queue
// before inserting the new element.
func (this *FifoLRUCache) Put(key any, value any) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if found, e := this.findByKey(key); found {
		// update the value
		e.value = value
	} else {
		// eliminate out-of-date entry if needed
		this.eliminate()
		// insert a new entry
		ne := &entry{}
		ne.key = key
		ne.value = value
		this.fifoQueue.EnQueue(ne)
	}
}

// Find the element from the LRU Cache by the key.
// Returns nil if no such key is found.
func (this *FifoLRUCache) Get(key any) any {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if found, entry := this.findByKey(key); found {
		return entry.value
	} else {
		return nil
	}
}

// Returns the size of the LRU Cache
func (this *FifoLRUCache) Size() int {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return this.fifoQueue.Size()
}

// Returns the capacity of the LRU Cache
func (this *FifoLRUCache) Cap() int {
	return this.capacity
}

// Clear the LRU Cache
func (this *FifoLRUCache) Clear() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.fifoQueue.Clear()
}

// For FifoLRUCache, remove the last element
func (this *FifoLRUCache) eliminate() {
	// check the size of LRU Cache
	if this.fifoQueue.Size() == this.capacity {
		this.fifoQueue.DeQueue()
	}
}

// Loop the Fifo queue and find the element by key
func (this *FifoLRUCache) findByKey(key any) (bool, *entry) {
	found := false
	var ele *entry
	it := this.fifoQueue.Iterate()
	for it.HasNext() {
		e := it.Next().(*entry)
		if e.key == key {
			found = true
			ele = e
			break
		}
	}
	return found, ele
}
