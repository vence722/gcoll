// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package gcoll

import (
	"gcoll/heap"
	"gcoll/list"
	"gcoll/lru"
	"gcoll/maps"
	"gcoll/set"
)

func NewArrayList[T comparable]() list.List[T] {
	return list.NewArrayList[T]()
}

func NewLinkedList[T comparable]() list.List[T] {
	return list.NewLinkedList[T]()
}

func NewQueue[T comparable]() list.Queue[T] {
	return list.NewLinkedList[T]()
}

func NewStack[T comparable]() list.Stack[T] {
	return list.NewLinkedList[T]()
}

func NewHashMap[K comparable, V any]() maps.Map[K, V] {
	return maps.NewHashMap[K, V]()
}

func NewSyncMap[K comparable, V any]() maps.SyncMap[K, V] {
	return maps.NewTypedSyncMap[K, V]()
}

func NewHashSet[T comparable]() set.Set[T] {
	return set.NewHashSet[T]()
}

func NewArrayHeap() heap.Heap {
	return heap.NewArrayHeap()
}

func NewFifoLRUCache[K comparable, V any](capacity int) lru.LRUCache[K, V] {
	return lru.NewFifoLRUCache[K, V](capacity)
}

func NewHitsMapLRUCache[K comparable, V any](capacity int) lru.LRUCache[K, V] {
	return lru.NewHitsMapLRUCache[K, V](capacity)
}

func NewPriorityQueue() heap.PriorityQueue {
	// Use ArrayHeap as Priority Queue
	return heap.NewArrayHeap()
}
