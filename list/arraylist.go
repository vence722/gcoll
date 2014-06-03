// gcoll v1.0
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)

package list

import (
	"fmt"
	"gcoll/collection"
)

// The ArrayList sturct
type ArrayList struct {
	elems []interface{}
}

// Return a new ArrayList
func NewArrayList() *ArrayList {
	return &ArrayList{make([]interface{}, INIT_LEN, INIT_CAP)}
}

// Return the size of this list
func (this *ArrayList) Size() int {
	return len(this.elems)
}

// Return the list containing elements or not
func (this *ArrayList) IsEmpty() bool {
	return len(this.elems) == 0
}

// Return the list conaining specified element or not
func (this *ArrayList) Contains(ele interface{}) bool {
	for _, o := range this.elems {
		if o == ele {
			return true
		}
	}
	return false
}

// Return a slice containing all the elements in this list
func (this *ArrayList) ToSlice() []interface{} {
	slice := make([]interface{}, INIT_LEN, INIT_CAP)
	for _, elem := range this.elems {
		slice = append(slice, elem)
	}
	return slice
}

// Return a Iterator of this list
func (this *ArrayList) Iterate() collection.Iterator {
	return &ArrayListIterator{this, -1}
}

// Add new element to this list
func (this *ArrayList) Add(ele interface{}) bool {
	this.elems = append(this.elems, ele)
	return true
}

// Remove FIRST specified element from this list
func (this *ArrayList) Remove(ele interface{}) bool {
	for i := 0; i < len(this.elems); i++ {
		o := this.elems[i]
		if o == ele {
			this.elems = append(this.elems[:i], this.elems[i+1:]...)
			return true
		}
	}
	return false
}

// Return the list containing specified elements or not
func (this *ArrayList) ContainsAll(c collection.Collection) bool {
	it := c.Iterate()
	for it.HasNext() {
		if !this.Contains(it.Next()) {
			return false
		}
	}
	return true
}

// Add a collection to this list
func (this *ArrayList) AddAll(c collection.Collection) bool {
	this.elems = append(this.elems, c.ToSlice()...)
	return true
}

// Remove all elements in specified collection from this list
// Could NOT call this method for the collection itself
func (this *ArrayList) RemoveAll(c collection.Collection) bool {
	b := false
	if this != c {
		it := c.Iterate()
		for it.HasNext() {
			b = this.Remove(it.Next()) || b
		}
	}
	return b
}

// Remove all elements from this list
func (this *ArrayList) Clear() {
	this.elems = this.elems[:0]
}

// Return the element at the specified position in this list
func (this *ArrayList) Get(index int) interface{} {
	if index < 0 || index >= this.Size() {
		panic("index out of bound")
	}
	return this.elems[index]
}

// Modify the element at the specified position in this list with the new one
func (this *ArrayList) Set(index int, ele interface{}) bool {
	if index < 0 || index >= this.Size() {
		panic("index out of bound")
	}
	this.elems[index] = ele
	return true
}

// Insert a new element at the specified position into this list
// if index == this.Size(), the new element will insert at the end of the list
func (this *ArrayList) Insert(index int, ele interface{}) bool {
	if index < 0 || index > this.Size() {
		panic("index out of bound")
	}
	if index == this.Size() {
		return this.Add(ele)
	}
	temp := append(this.elems[:index], ele)
	this.elems = append(temp, this.elems[index+1:]...)
	return true
}

// Remove the element at the specified position in this list
func (this *ArrayList) RemoveAt(index int) interface{} {
	if index < 0 || index >= this.Size() {
		panic("index out of bound")
	}
	ele := this.elems[index]
	this.elems = append(this.elems[:index], this.elems[index+1:]...)
	return ele
}

// Returns the index of the first occurrence of the specified element in this list,
// or -1 if this list does not contain the element.
func (this *ArrayList) IndexOf(ele interface{}) int {
	for i := 0; i < this.Size(); i++ {
		if ele == this.elems[i] {
			return i
		}
	}
	return -1
}

// Returns the index of the last occurrence of the specified element in this list,
// or -1 if this list does not contain the element.
func (this *ArrayList) LastIndexOf(ele interface{}) int {
	for i := this.Size() - 1; i >= 0; i-- {
		if ele == this.elems[i] {
			return i
		}
	}
	return -1
}

// Returns a view of the portion of this list between the specified range
func (this *ArrayList) SubList(fromIndex, toIndex int) List {
	if fromIndex < 0 || fromIndex >= this.Size() {
		panic("fromIndex out of bound")
	}
	if toIndex < 0 || toIndex >= this.Size() {
		panic("toIndex out of bound")
	}
	if fromIndex > toIndex {
		panic("fromIndex can't larger than toIndex")
	}
	return &ArrayList{this.elems[fromIndex:toIndex]}
}

// Return the string that describes the contains of this list
func (this *ArrayList) String() string {
	return fmt.Sprint(this.elems)
}

// The iterator struct
type ArrayListIterator struct {
	list  *ArrayList
	index int
}

// Does the iterator have more elements?
func (it *ArrayListIterator) HasNext() bool {
	if it.index+1 < it.list.Size() {
		return true
	}
	return false
}

// Return the next element of the iterator
func (it *ArrayListIterator) Next() interface{} {
	if it.HasNext() {
		it.index++
		return it.list.elems[it.index]
	}
	return nil
}
