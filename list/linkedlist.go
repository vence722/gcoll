// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package list

import (
	"fmt"

	"gcoll/collection"
)

// The LinkedList struct
type LinkedList struct {
	size int
	head *entry
	tail *entry
}

// Each element of a LinkedList is stored in an entry
type entry struct {
	elem any
	next *entry
}

// Return a new LinkedList
func NewLinkedList() *LinkedList {
	return &LinkedList{0, nil, nil}
}

// Return the size of this list
func (this *LinkedList) Size() int {
	return this.size
}

// Return the list containing elements or not
func (this *LinkedList) IsEmpty() bool {
	return this.size == 0
}

// Return the list conaining specified element or not
func (this *LinkedList) Contains(ele any) bool {
	if this.head == nil {
		return false
	}
	pe := this.head
	for pe != this.tail {
		if pe.elem == ele {
			return true
		}
		pe = pe.next
	}
	return false
}

// Return a slice containing all the elements in this list
func (this *LinkedList) ToSlice() []any {
	slice := make([]any, InitLen, InitCap)
	pe := this.head
	for pe != nil {
		slice = append(slice, pe.elem)
		pe = pe.next
	}
	return slice
}

// Return a Iterator of this list
func (this *LinkedList) Iterate() collection.Iterator {
	return &LinkedListIterator{this, -1}
}

// Add new element to this list
func (this *LinkedList) Add(ele any) bool {
	en := entry{ele, nil}
	if this.head == nil && this.tail == nil {
		this.head = &en
		this.tail = &en
	} else {
		this.tail.next = &en
		this.tail = &en
	}
	this.size++
	return true
}

// Remove FIRST specified element from this list
func (this *LinkedList) Remove(ele any) bool {
	if this.head == nil && this.tail == nil {
		return false
	}
	pe := this.head
	var ps *entry
	ps = nil
	for pe != nil {
		if pe.elem == ele {
			if ps != nil {
				ps.next = pe.next
				if pe == this.tail {
					this.tail = ps
				}
				this.size--
			} else {
				// remove first element
				this.head = pe.next
				// if remove the only one element, tail must be change to nil
				if this.tail == pe {
					this.tail = nil
				}
				this.size--
			}
			return true
		}
		ps = pe
		pe = pe.next
	}
	return false
}

// Return the list containing specified elements or not
func (this *LinkedList) ContainsAll(c collection.Collection) bool {
	it := c.Iterate()
	for it.HasNext() {
		if !this.Contains(it.Next()) {
			return false
		}
	}
	return true
}

// Add a collection to this list
func (this *LinkedList) AddAll(c collection.Collection) bool {
	it := c.Iterate()
	for it.HasNext() {
		this.Add(it.Next())
	}
	return true
}

// Remove all elements in specified collection from this list
// Could NOT call this method for the collection itself
func (this *LinkedList) RemoveAll(c collection.Collection) bool {
	b := false
	it := c.Iterate()
	for it.HasNext() {
		b = b || this.Remove(it.Next())
	}
	return true
}

// Remove all elements from this list
func (this *LinkedList) Clear() {
	this.head = nil
	this.tail = nil
	this.size = 0
}

// Return the element at the specified position in this list
func (this *LinkedList) Get(index int) any {
	if index < 0 || index >= this.Size() {
		panic("index out of bound")
	}
	pe := this.head
	for i := 0; i < index; i++ {
		pe = pe.next
	}
	return pe.elem
}

// Modify the element at the specified position in this list with the new one
func (this *LinkedList) Set(index int, ele any) bool {
	if index < 0 || index >= this.Size() {
		panic("index out of bound")
	}
	pe := this.head
	for i := 0; i < index; i++ {
		pe = pe.next
	}
	pe.elem = ele
	return true
}

// Insert a new element at the specified position into this list
// if index == this.Size(), the new element will insert at the end of the list
func (this *LinkedList) Insert(index int, ele any) bool {
	if index < 0 || index > this.Size() {
		panic("index out of bound")
	}
	if index == this.Size() {
		return this.Add(ele)
	}
	en := entry{ele, nil}
	pe := this.head
	var ps *entry
	ps = nil
	for i := 0; i < index; i++ {
		ps = pe
		pe = pe.next
	}
	if ps != nil {
		ps.next = &en
		en.next = pe
	} else {
		// insert in the first place
		this.head = &en
		en.next = pe
	}
	this.size++
	return true
}

// Remove the element at the specified position in this list
func (this *LinkedList) RemoveAt(index int) any {
	if index < 0 || index >= this.Size() {
		panic("index out of bound")
	}
	pe := this.head
	var ps *entry
	var ele any
	ps = nil
	for i := 0; i < index; i++ {
		ps = pe
		pe = pe.next
	}
	if ps != nil {
		ele = pe.elem
		ps.next = pe.next
		if pe == this.tail {
			// remove the last element
			this.tail = ps
		}
	} else {
		// remove the first element
		ele = pe.elem
		this.head = this.head.next
		// if remove the only one element, tail must be change to nil
		if this.tail == pe {
			this.tail = nil
		}
	}
	this.size--
	return ele
}

// Returns the index of the first occurrence of the specified element in this list,
// or -1 if this list does not contain the element.
func (this *LinkedList) IndexOf(ele any) int {
	pe := this.head
	index := -1
	for pe != nil {
		index++
		if pe.elem == ele {
			return index
		}
		pe = pe.next
	}
	return -1
}

// Returns the index of the last occurrence of the specified element in this list,
// or -1 if this list does not contain the element.
func (this *LinkedList) LastIndexOf(ele any) int {
	lastIndex := -1
	pe := this.head
	index := -1
	for pe != nil {
		index++
		if pe.elem == ele {
			lastIndex = index
		}
		pe = pe.next
	}
	return lastIndex
}

// Returns a view of the portion of this list between the specified range
func (this *LinkedList) SubList(fromIndex, toIndex int) List {
	if fromIndex < 0 || fromIndex >= this.Size() {
		panic("fromIndex out of bound")
	}
	if toIndex < 0 || toIndex >= this.Size() {
		panic("toIndex out of bound")
	}
	if fromIndex > toIndex {
		panic("fromIndex can't larger than toIndex")
	}
	l := LinkedList{0, nil, nil}
	pe := this.head
	index := -1
	for pe != nil {
		index++
		if index >= toIndex {
			break
		}
		if index >= fromIndex {
			l.Add(pe.elem)
		}
		pe = pe.next
	}
	return &l
}

// Return the string that describes the contains of this list
func (this *LinkedList) String() string {
	return fmt.Sprint(this.ToSlice())
}

// The iterator struct
type LinkedListIterator struct {
	list  *LinkedList
	index int
}

// Does the iterator have more elements?
func (it *LinkedListIterator) HasNext() bool {
	if it.index+1 < it.list.Size() {
		return true
	}
	return false
}

// Return the next element of the iterator
func (it *LinkedListIterator) Next() any {
	if it.HasNext() {
		it.index++
		return it.list.Get(it.index)
	}
	return nil
}

// Put an element at the tail of the Queue
func (this *LinkedList) EnQueue(ele any) bool {
	return this.Add(ele)
}

// Remove the first element of the Queue and return it
// If there're no elements, return a nil value
func (this *LinkedList) DeQueue() any {
	if this.head == nil {
		return nil
	}
	return this.RemoveAt(0)
}

// Return the front of the Queue
// If there's no element, return a nil value
func (this *LinkedList) Front() any {
	if this.head == nil {
		return nil
	}
	return this.head.elem
}

// Return the tail of the Queue
// If there's no element, return a nil value
func (this *LinkedList) Tail() any {
	if this.tail == nil {
		return nil
	}
	return this.tail.elem
}

// Push an element into the Stack
func (this *LinkedList) Push(ele any) bool {
	return this.Insert(0, ele)
}

// Pop the last pushed element from the Stack
// If there's no element, return a nil value
func (this *LinkedList) Pop() any {
	if this.head == nil {
		return nil
	}
	return this.RemoveAt(0)
}

// Return the last pushed element from the Stack
// If there's no element, return a nil value
func (this *LinkedList) Top() any {
	if this.head == nil {
		return nil
	}
	return this.head.elem
}
