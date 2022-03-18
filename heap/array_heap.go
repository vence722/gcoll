// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)
package heap

import (
	cheap "container/heap"
)

// A heap implementation using array
type ArrayHeap[T Comparable] struct {
	internal cheap.Interface
}

// The internal implementation
type arrayHeap[T Comparable] struct {
	heap []T
}

func (h *arrayHeap[T]) Len() int {
	return len((*h).heap)
}

func (h *arrayHeap[T]) Less(i, j int) bool {
	return (*h).heap[i].CompareTo((*h).heap[j]) < 0
}

func (h *arrayHeap[T]) Swap(i, j int) {
	(*h).heap[i], (*h).heap[j] = (*h).heap[j], (*h).heap[i]
}

func (h *arrayHeap[T]) Push(x any) {
	(*h).heap = append((*h).heap, x.(T))
}

func (h *arrayHeap[T]) Pop() any {
	n := len((*h).heap)
	x := (*h).heap[n-1]
	(*h).heap = (*h).heap[0 : n-1]
	return x
}

func NewArrayHeap[T Comparable]() *ArrayHeap[T] {
	h := &ArrayHeap[T]{}
	h.internal = &arrayHeap[T]{}
	cheap.Init(h.internal)
	return h
}

func (this *ArrayHeap[T]) Put(ele T) {
	cheap.Push(this.internal, ele)
}

func (this *ArrayHeap[T]) Top() T {
	if this.Empty() {
		var empty T
		return empty
	}
	return (*this.internal.(*arrayHeap[T])).heap[0]
}

func (this *ArrayHeap[T]) Take() T {
	return cheap.Pop(this.internal).(T)
}

func (this *ArrayHeap[T]) Empty() bool {
	h := *this.internal.(*arrayHeap[T])
	return h.Len() <= 0
}

func (this *ArrayHeap[T]) Size() int {
	return this.internal.Len()
}

// Consider ArrayHeap as Primary Queue
func (this *ArrayHeap[T]) EnQueue(ele T) bool {
	this.Put(ele)
	return true
}

func (this *ArrayHeap[T]) DeQueue() T {
	return this.Take()
}

func (this *ArrayHeap[T]) Front() T {
	return this.Top()
}
