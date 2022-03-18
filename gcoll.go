// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package gcoll

import (
	"github.com/vence722/gcoll/heap"
	"github.com/vence722/gcoll/list"
	"github.com/vence722/gcoll/lru"
	"github.com/vence722/gcoll/maps"
	"github.com/vence722/gcoll/set"
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

func NewArrayHeap[T heap.Comparable]() heap.Heap[T] {
	return heap.NewArrayHeap[T]()
}

func NewFifoLRUCache[K comparable, V any](capacity int) lru.LRUCache[K, V] {
	return lru.NewFifoLRUCache[K, V](capacity)
}

func NewHitsMapLRUCache[K comparable, V any](capacity int) lru.LRUCache[K, V] {
	return lru.NewHitsMapLRUCache[K, V](capacity)
}

func NewPriorityQueue[T heap.Comparable]() heap.PriorityQueue[T] {
	// Use ArrayHeap as Priority Queue
	return heap.NewArrayHeap[T]()
}
