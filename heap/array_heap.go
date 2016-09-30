// gcoll
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)
package heap

import (
	cheap "container/heap"
)

// A heap implementation using array
type ArrayHeap struct {
	internal cheap.Interface
}

// The internal implementation
type arrayHeap []Comparable

func (h *arrayHeap) Len() int {
	return len(*h)
}

func (h *arrayHeap) Less(i, j int) bool {
	return (*h)[i].CompareTo((*h)[j]) < 0
}

func (h *arrayHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *arrayHeap) Push(x interface{}) {
	*h = append(*h, x.(Comparable))
}

func (h *arrayHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func NewArrayHeap() *ArrayHeap {
	h := &ArrayHeap{}
	h.internal = &arrayHeap{}
	cheap.Init(h.internal)
	return h
}

func (this *ArrayHeap) Put(elem Comparable) {
	cheap.Push(this.internal, elem)
}

func (this *ArrayHeap) Top() Comparable {
	if this.Empty() {
		return nil
	}
	return (*this.internal.(*arrayHeap))[0]
}

func (this *ArrayHeap) Take() Comparable {
	return cheap.Pop(this.internal).(Comparable)
}

func (this *ArrayHeap) Empty() bool {
	h := *this.internal.(*arrayHeap)
	return h.Len() <= 0
}

func (this *ArrayHeap) Size() int {
	return this.internal.Len()
}

// Consider ArrayHeap as Primary Queue
func (this *ArrayHeap) EnQueue(ele Comparable) bool {
	this.Put(ele)
	return true
}

func (this *ArrayHeap) DeQueue() Comparable {
	return this.Take()
}

func (this *ArrayHeap) Front() Comparable {
	return this.Top()
}
