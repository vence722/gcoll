// gcoll
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)
package gcoll

import (
	"github.com/vence722/gcoll/heap"
	"github.com/vence722/gcoll/list"
	"github.com/vence722/gcoll/lru"
	"github.com/vence722/gcoll/maps"
	"github.com/vence722/gcoll/set"
	"github.com/vence722/gcoll/tree"
)

func NewArrayList() list.List {
	return list.NewArrayList()
}

func NewLinkedList() list.List {
	return list.NewLinkedList()
}

func NewQueue() list.Queue {
	return list.NewLinkedList()
}

func NewStack() list.Stack {
	return list.NewLinkedList()
}

func NewHashMap() maps.Map {
	return maps.NewHashMap()
}

func NewBSTreeMap() maps.StringMap {
	return tree.NewBinarySortTree()
}

func NewHashSet() set.Set {
	return set.NewHashSet()
}

func NewArrayHeap() heap.Heap {
	return heap.NewArrayHeap()
}

func NewFifoLRUCache(capacity int) lru.LRUCache {
	return lru.NewFifoLRUCache(capacity)
}

func NewHitsMapLRUCache(capacity int) lru.LRUCache {
	return lru.NewHitsMapLRUCache(capacity)
}

func NewPriorityQueue() heap.PriorityQueue {
	// Use ArrayHeap as Priority Queue
	return heap.NewArrayHeap()
}
