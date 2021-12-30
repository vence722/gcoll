// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package gcoll

import (
	"gcoll/graph"
	"gcoll/heap"
	"gcoll/list"
	"gcoll/lru"
	"gcoll/maps"
	"gcoll/set"
	"gcoll/tree"
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

func NewSimpleGraph() graph.Graph {
	return graph.NewSimpleGraph()
}

func NewWeightedSimpleGraph() graph.WeightedGraph {
	return graph.NewSimpleWeightedGraph()
}

func NewAdjacencyMatrixGraph() graph.WeightedGraph {
	return graph.NewAdjacencyMatrixGraph()
}
