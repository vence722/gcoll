// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package list

import (
	"fmt"

	"gcoll/collection"
)

// The LinkedList struct
type LinkedList[T comparable] struct {
	size int
	head *entry[T]
	tail *entry[T]
}

// Each element of a LinkedList is stored in an entry
type entry[T comparable] struct {
	ele  T
	next *entry[T]
}

// NewLinkedList Returns a new LinkedList
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{0, nil, nil}
}

// Size Returns the size of this list
func (list *LinkedList[T]) Size() int {
	return list.size
}

// IsEmpty Returns the list containing elements or not
func (list *LinkedList[T]) IsEmpty() bool {
	return list.size == 0
}

// Contains Returns whether the list contains the specified element or not
func (list *LinkedList[T]) Contains(ele T) bool {
	if list.head == nil {
		return false
	}
	pe := list.head
	for pe != list.tail {
		if pe.ele == ele {
			return true
		}
		pe = pe.next
	}
	return false
}

// ToSlice Returns a slice that contains all the elements in this list
func (list *LinkedList[T]) ToSlice() []T {
	slice := make([]T, InitLen, InitCap)
	pe := list.head
	for pe != nil {
		slice = append(slice, pe.ele)
		pe = pe.next
	}
	return slice
}

// Iterate Returns an Iterator of this list
func (list *LinkedList[T]) Iterate() collection.Iterator[T] {
	return &LinkedListIterator[T]{list, -1}
}

// Add Adds new element to this list
func (list *LinkedList[T]) Add(ele T) bool {
	en := entry[T]{ele, nil}
	if list.head == nil && list.tail == nil {
		list.head = &en
		list.tail = &en
	} else {
		list.tail.next = &en
		list.tail = &en
	}
	list.size++
	return true
}

// Remove Removes FIRST specified element from this list
func (list *LinkedList[T]) Remove(ele T) bool {
	if list.head == nil && list.tail == nil {
		return false
	}
	pe := list.head
	var ps *entry[T]
	ps = nil
	for pe != nil {
		if pe.ele == ele {
			if ps != nil {
				ps.next = pe.next
				if pe == list.tail {
					list.tail = ps
				}
				list.size--
			} else {
				// remove first element
				list.head = pe.next
				// if remove the only one element, tail must be change to nil
				if list.tail == pe {
					list.tail = nil
				}
				list.size--
			}
			return true
		}
		ps = pe
		pe = pe.next
	}
	return false
}

// ContainsAll Returns whether the list contains specified elements or not
func (list *LinkedList[T]) ContainsAll(c collection.Collection[T]) bool {
	it := c.Iterate()
	for it.HasNext() {
		if !list.Contains(it.Next()) {
			return false
		}
	}
	return true
}

// AddAll Adds a collection to this list
func (list *LinkedList[T]) AddAll(c collection.Collection[T]) bool {
	it := c.Iterate()
	for it.HasNext() {
		list.Add(it.Next())
	}
	return true
}

// RemoveAll Removes all elements in specified collection from this list
// Could NOT call this method for the collection itself
func (list *LinkedList[T]) RemoveAll(c collection.Collection[T]) bool {
	b := false
	if list != c {
		it := c.Iterate()
		for it.HasNext() {
			b = list.Remove(it.Next()) || b
		}
	}
	return b
}

// Clear Removes all elements from this list
func (list *LinkedList[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

// Get Returns the element at the specified position in this list
func (list *LinkedList[T]) Get(index int) (T, bool) {
	if index < 0 || index >= list.Size() {
		var zero T
		return zero, false
	}
	pe := list.head
	for i := 0; i < index; i++ {
		pe = pe.next
	}
	return pe.ele, true
}

// MustGet Returns the element at the specified position in this list
// Returns empty value if the index specified is invalid
func (list *LinkedList[T]) MustGet(index int) T {
	ele, _ := list.Get(index)
	return ele
}

// Set Modifies the element at the specified position in this list with the new one
func (list *LinkedList[T]) Set(index int, ele T) bool {
	if index < 0 || index >= list.Size() {
		return false
	}
	pe := list.head
	for i := 0; i < index; i++ {
		pe = pe.next
	}
	pe.ele = ele
	return true
}

// Insert Inserts a new element at the specified position into this list
// if index == this.Size(), the new element will insert at the end of the list
func (list *LinkedList[T]) Insert(index int, ele T) bool {
	if index < 0 || index > list.Size() {
		return false
	}
	if index == list.Size() {
		return list.Add(ele)
	}
	en := entry[T]{ele, nil}
	pe := list.head
	var ps *entry[T]
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
		list.head = &en
		en.next = pe
	}
	list.size++
	return true
}

// RemoveAt Removes the element at the specified position in this list
func (list *LinkedList[T]) RemoveAt(index int) (T, bool) {
	if index < 0 || index >= list.Size() {
		var zero T
		return zero, false
	}
	pe := list.head
	var ps *entry[T]
	var ele T
	ps = nil
	for i := 0; i < index; i++ {
		ps = pe
		pe = pe.next
	}
	if ps != nil {
		ele = pe.ele
		ps.next = pe.next
		if pe == list.tail {
			// remove the last element
			list.tail = ps
		}
	} else {
		// remove the first element
		ele = pe.ele
		list.head = list.head.next
		// if remove the only one element, tail must be change to nil
		if list.tail == pe {
			list.tail = nil
		}
	}
	list.size--
	return ele, true
}

// IndexOf Returns the index of the first occurrence of the specified element in this list,
// or -1 if this list does not contain the element.
func (list *LinkedList[T]) IndexOf(ele T) int {
	pe := list.head
	index := -1
	for pe != nil {
		index++
		if pe.ele == ele {
			return index
		}
		pe = pe.next
	}
	return -1
}

// LastIndexOf Returns the index of the last occurrence of the specified element in this list,
// or -1 if this list does not contain the element.
func (list *LinkedList[T]) LastIndexOf(ele T) int {
	lastIndex := -1
	pe := list.head
	index := -1
	for pe != nil {
		index++
		if pe.ele == ele {
			lastIndex = index
		}
		pe = pe.next
	}
	return lastIndex
}

// Returns a view of the portion of this list between the specified range
func (list *LinkedList[T]) SubList(fromIndex, toIndex int) List[T] {
	emptyList := NewLinkedList[T]()
	if fromIndex < 0 || fromIndex >= list.Size() {
		return emptyList
	}
	if toIndex < 0 || toIndex >= list.Size() {
		return emptyList
	}
	if fromIndex > toIndex {
		return emptyList
	}
	l := &LinkedList[T]{0, nil, nil}
	pe := list.head
	index := -1
	for pe != nil {
		index++
		if index >= toIndex {
			break
		}
		if index >= fromIndex {
			l.Add(pe.ele)
		}
		pe = pe.next
	}
	return l
}

// String Returns the string that describes the contents of this list
func (list *LinkedList[T]) String() string {
	return fmt.Sprint(list.ToSlice())
}

// LinkedListIterator The iterator struct
type LinkedListIterator[T comparable] struct {
	list  *LinkedList[T]
	index int
}

// HasNext Does the iterator have more elements?
func (it *LinkedListIterator[T]) HasNext() bool {
	if it.index+1 < it.list.Size() {
		return true
	}
	return false
}

// TODO optimize
// Next Returns the next element of the iterator
func (it *LinkedListIterator[T]) Next() T {
	if it.HasNext() {
		it.index++
		return it.list.MustGet(it.index)
	}
	var zero T
	return zero
}

// EnQueue Puts an element at the tail of the Queue
func (list *LinkedList[T]) EnQueue(ele T) bool {
	return list.Add(ele)
}

// DeQueue Removes the first element of the Queue and return it
func (list *LinkedList[T]) DeQueue() (T, bool) {
	if list.head == nil {
		var zero T
		return zero, false
	}
	return list.RemoveAt(0)
}

// Front Returns the front of the Queue
func (list *LinkedList[T]) Front() (T, bool) {
	if list.head == nil {
		var zero T
		return zero, false
	}
	return list.head.ele, true
}

// Tail Returns the tail of the Queue
func (list *LinkedList[T]) Tail() (T, bool) {
	if list.tail == nil {
		var zero T
		return zero, false
	}
	return list.tail.ele, true
}

// Push Pushs an element into the Stack
func (list *LinkedList[T]) Push(ele T) bool {
	return list.Insert(0, ele)
}

// Pop Pops the last pushed element from the Stack
func (list *LinkedList[T]) Pop() (T, bool) {
	if list.head == nil {
		var zero T
		return zero, false
	}
	return list.RemoveAt(0)
}

// Top Returns the last pushed element from the Stack
func (list *LinkedList[T]) Top() (T, bool) {
	if list.head == nil {
		var zero T
		return zero, false
	}
	return list.head.ele, true
}
