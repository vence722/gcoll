// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package lru

import (
	"sync"

	"github.com/vence722/gcoll/list"
)

// FifoLRUCache LRU Cache implementation of Fifo Queue
type FifoLRUCache[K comparable, V any] struct {
	fifoQueue list.Queue[*entry[K, V]]
	capacity  int
	mutex     sync.Mutex
}

func NewFifoLRUCache[K comparable, V any](capacity int) *FifoLRUCache[K, V] {
	return &FifoLRUCache[K, V]{
		fifoQueue: list.NewLinkedList[*entry[K, V]](),
		capacity:  capacity,
	}
}

// Put Puts new element into the LRU cache.
// If key alrady exists, update the value.
// If the key is not found in the LRU cache,
// and the LRU cache is not full, add a new entry.
// Otherwise remove the tail of the FIFO queue
// before inserting the new element.
func (lru *FifoLRUCache[K, V]) Put(key K, value V) bool {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	if e, found := lru.findByKey(key); found {
		// update the value
		e.value = value
		return true
	} else {
		// eliminate out-of-date entry if needed
		lru.eliminate()
		// insert a new entry
		ne := &entry[K, V]{}
		ne.key = key
		ne.value = value
		return lru.fifoQueue.EnQueue(ne)
	}
}

// Get Finds the element from the LRU Cache by the key.
// Returns nil if no such key is found.
func (lru *FifoLRUCache[K, V]) Get(key K) (ele V, ok bool) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	if entry, found := lru.findByKey(key); found {
		return entry.value, found
	}
	var zero V
	return zero, false
}

// Size Returns the size of the LRU Cache
func (lru *FifoLRUCache[K, V]) Size() int {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	return lru.fifoQueue.Size()
}

// Cap Returns the capacity of the LRU Cache
func (lru *FifoLRUCache[K, V]) Cap() int {
	return lru.capacity
}

// Clear Clears the LRU Cache
func (lru *FifoLRUCache[K, V]) Clear() {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	lru.fifoQueue.Clear()
}

// eliminate For FifoLRUCache, remove the last element
func (lru *FifoLRUCache[K, V]) eliminate() {
	// check the size of LRU Cache
	if lru.fifoQueue.Size() == lru.capacity {
		lru.fifoQueue.DeQueue()
	}
}

// findByKey Loop the Fifo queue and find the entry by key
func (lru *FifoLRUCache[K, V]) findByKey(key K) (*entry[K, V], bool) {
	found := false
	var ele *entry[K, V]
	it := lru.fifoQueue.Iterate()
	for it.HasNext() {
		e := it.Next()
		if e.key == key {
			found = true
			ele = e
			break
		}
	}
	return ele, found
}
