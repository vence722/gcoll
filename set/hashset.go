// gcoll
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)

package set

import (
	"fmt"

	"github.com/vence722/gcoll/collection"
	"github.com/vence722/gcoll/maps"
)

const (
	PRESENT = ""
)

// The HashSet struct
type HashSet struct {
	mmap maps.Map
}

// Return a new HashSet
func NewHashSet() *HashSet {
	mmap := maps.NewHashMap()
	return &HashSet{mmap}
}

// Return the size of this set
func (this *HashSet) Size() int {
	return this.mmap.Size()
}

// Return the set containing elements or not
func (this *HashSet) IsEmpty() bool {
	return this.mmap.Size() != 0
}

// Return the set conaining specified element or not
func (this *HashSet) Contains(ele interface{}) bool {
	return this.mmap.ContainsKey(ele)
}

// Return a slice containing all the elements in this set
func (this *HashSet) ToSlice() []interface{} {
	return this.mmap.Keys()
}

// Return a Iterator of this set
func (this *HashSet) Iterate() collection.Iterator {
	return &HashSetIterator{this.mmap.Keys(), -1}
}

// Add new element to this set
func (this *HashSet) Add(ele interface{}) bool {
	return this.mmap.Put(ele, PRESENT)
}

// Remove specified element from this set
func (this *HashSet) Remove(ele interface{}) bool {
	return this.mmap.Remove(ele)
}

// Add a collection to this set
func (this *HashSet) AddAll(c collection.Collection) bool {
	it := c.Iterate()
	for it.HasNext() {
		this.Add(it.Next())
	}
	return true
}

// Return the list containing specified elements or not
func (this *HashSet) ContainsAll(c collection.Collection) bool {
	it := c.Iterate()
	for it.HasNext() {
		if !this.Contains(it.Next()) {
			return false
		}
	}
	return true
}

// Remove all elements in specified collection from this set
// Could NOT call this method for the collection itself
func (this *HashSet) RemoveAll(c collection.Collection) bool {
	b := false
	if this != c {
		it := c.Iterate()
		for it.HasNext() {
			b = this.Remove(it.Next()) || b
		}
	}
	return b
}

// Remove all elements from this set
func (this *HashSet) Clear() {
	this.mmap.Clear()
}

// Return the string that describes the contains of this list
func (this *HashSet) String() string {
	return fmt.Sprint(this.ToSlice())
}

// The iterator struct
type HashSetIterator struct {
	values []interface{}
	index  int
}

// Does the iterator have more elements?
func (it *HashSetIterator) HasNext() bool {
	if it.index+1 < len(it.values) {
		return true
	}
	return false
}

// Return the next element of the iterator
func (it *HashSetIterator) Next() interface{} {
	if it.HasNext() {
		it.index++
		return it.values[it.index]
	}
	return nil
}
