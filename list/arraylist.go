// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package list

import (
	"fmt"

	"github.com/vence722/gcoll/collection"
)

// The ArrayList sturct
type ArrayList[T comparable] struct {
	elems []T
}

// NewArrayList Returns a new ArrayList
func NewArrayList[T comparable]() *ArrayList[T] {
	return &ArrayList[T]{make([]T, InitLen, InitCap)}
}

// Size Returns the size of this list
func (list *ArrayList[T]) Size() int {
	return len(list.elems)
}

// IsEmpty Returns the list containing elements or not
func (list *ArrayList[T]) IsEmpty() bool {
	return len(list.elems) == 0
}

// Contains Returns whether the list contains the specified element or not
func (list *ArrayList[T]) Contains(ele T) bool {
	for _, o := range list.elems {
		if o == ele {
			return true
		}
	}
	return false
}

// ToSlice Returns a slice containing all the elements in this list
func (list *ArrayList[T]) ToSlice() []T {
	slice := make([]T, InitLen, InitCap)
	for _, ele := range list.elems {
		slice = append(slice, ele)
	}
	return slice
}

// Iterate Returns an Iterator of this list
func (list *ArrayList[T]) Iterate() collection.Iterator[T] {
	return &ArrayListIterator[T]{list, -1}
}

// Add new element to this list
func (list *ArrayList[T]) Add(ele T) bool {
	list.elems = append(list.elems, ele)
	return true
}

// Remove FIRST specified element from this list
func (list *ArrayList[T]) Remove(ele T) bool {
	for i := 0; i < len(list.elems); i++ {
		o := list.elems[i]
		if o == ele {
			list.elems = append(list.elems[:i], list.elems[i+1:]...)
			return true
		}
	}
	return false
}

// ContainsAll Returns whether the list contains a specified collection of elements or not
func (list *ArrayList[T]) ContainsAll(c collection.Collection[T]) bool {
	it := c.Iterate()
	for it.HasNext() {
		if !list.Contains(it.Next()) {
			return false
		}
	}
	return true
}

// AddAll Adds a collection to this list
func (list *ArrayList[T]) AddAll(c collection.Collection[T]) bool {
	list.elems = append(list.elems, c.ToSlice()...)
	return true
}

// RemoveAll Removes all elements in specified collection from this list
// Could NOT call this method for the collection itself
func (list *ArrayList[T]) RemoveAll(c collection.Collection[T]) bool {
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
func (list *ArrayList[T]) Clear() {
	list.elems = list.elems[:0]
}

// Get Returns the element at the specified position in this list
// Returns ele = empty value, ok = false if the index specified is invalid
func (list *ArrayList[T]) Get(index int) (T, bool) {
	if index < 0 || index >= list.Size() {
		var zero T
		return zero, false
	}
	return list.elems[index], true
}

// MustGet Returns the element at the specified position in this list
// Returns empty value if the index specified is invalid
func (list *ArrayList[T]) MustGet(index int) T {
	ele, _ := list.Get(index)
	return ele
}

// Set Modifies the element at the specified position in this list with the new one
func (list *ArrayList[T]) Set(index int, ele T) bool {
	if index < 0 || index >= list.Size() {
		return false
	}
	list.elems[index] = ele
	return true
}

// Insert Inserts a new element at the specified position into this list
// if index == this.Size(), the new element will insert at the end of the list
func (list *ArrayList[T]) Insert(index int, ele T) bool {
	if index < 0 || index > list.Size() {
		return false
	}
	if index == list.Size() {
		return list.Add(ele)
	}
	temp := append(list.elems[:index], ele)
	list.elems = append(temp, list.elems[index+1:]...)
	return true
}

// RemoveAt Removes the element at the specified position in this list
func (list *ArrayList[T]) RemoveAt(index int) (T, bool) {
	if index < 0 || index >= list.Size() {
		var zero T
		return zero, false
	}
	ele := list.elems[index]
	list.elems = append(list.elems[:index], list.elems[index+1:]...)
	return ele, true
}

// IndexOf Returns the index of the first occurrence of the specified element in this list,
// or -1 if this list does not contain the element.
func (list *ArrayList[T]) IndexOf(ele T) int {
	for i := 0; i < list.Size(); i++ {
		if ele == list.elems[i] {
			return i
		}
	}
	return -1
}

// LastIndexOf Returns the index of the last occurrence of the specified element in this list,
// or -1 if this list does not contain the element.
func (list *ArrayList[T]) LastIndexOf(ele T) int {
	for i := list.Size() - 1; i >= 0; i-- {
		if ele == list.elems[i] {
			return i
		}
	}
	return -1
}

// SubList Returns a view of the portion of this list between the specified range
func (list *ArrayList[T]) SubList(fromIndex, toIndex int) List[T] {
	emptyList := NewArrayList[T]()
	if fromIndex < 0 || fromIndex >= list.Size() {
		return emptyList
	}
	if toIndex < 0 || toIndex >= list.Size() {
		return emptyList
	}
	if fromIndex > toIndex {
		return emptyList
	}
	return &ArrayList[T]{list.elems[fromIndex:toIndex]}
}

// String Returns the string that describes the contents of this list
func (list *ArrayList[T]) String() string {
	return fmt.Sprint(list.elems)
}

// ArrayListIterator The iterator struct
type ArrayListIterator[T comparable] struct {
	list  *ArrayList[T]
	index int
}

// HasNext Does the iterator have more elements?
func (it *ArrayListIterator[T]) HasNext() bool {
	if it.index+1 < it.list.Size() {
		return true
	}
	return false
}

// Next Return the next element of the iterator
func (it *ArrayListIterator[T]) Next() T {
	if it.HasNext() {
		it.index++
		return it.list.elems[it.index]
	}
	var zero T
	return zero
}
