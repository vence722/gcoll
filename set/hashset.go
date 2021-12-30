// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package set

import (
	"fmt"

	"gcoll/collection"
	"gcoll/maps"
)

const (
	Present = ""
)

// The HashSet struct
type HashSet[T comparable] struct {
	mmap maps.Map[T, any]
}

// NewHashSet Return a new HashSet
func NewHashSet[T comparable]() *HashSet[T] {
	mmap := maps.NewHashMap[T, any]()
	return &HashSet[T]{mmap}
}

// Size Returns the size of this set
func (s *HashSet[T]) Size() int {
	return s.mmap.Size()
}

// IsEmpty Returns the set containing elements or not
func (s *HashSet[T]) IsEmpty() bool {
	return s.mmap.Size() != 0
}

// Contains Returns whether the set contains specified element or not
func (s *HashSet[T]) Contains(ele T) bool {
	return s.mmap.ContainsKey(ele)
}

// ToSlice Returns a slice containing all the elements in this set
func (s *HashSet[T]) ToSlice() []T {
	return s.mmap.Keys()
}

// Iterate Returns a Iterator of this set
func (s *HashSet[T]) Iterate() collection.Iterator[T] {
	return &HashSetIterator[T]{s.mmap.Keys(), -1}
}

// Add new element to this set
func (s *HashSet[T]) Add(ele T) bool {
	return s.mmap.Put(ele, Present)
}

// Remove Removes specified element from this set
func (s *HashSet[T]) Remove(ele T) bool {
	return s.mmap.Remove(ele)
}

// AddAll Adds a collection to this set
func (s *HashSet[T]) AddAll(c collection.Collection[T]) bool {
	it := c.Iterate()
	for it.HasNext() {
		s.Add(it.Next())
	}
	return true
}

// ContainsAll Returns whether the list contains the specified elements or not
func (s *HashSet[T]) ContainsAll(c collection.Collection[T]) bool {
	it := c.Iterate()
	for it.HasNext() {
		if !s.Contains(it.Next()) {
			return false
		}
	}
	return true
}

// Remove all elements in specified collection from this set
// Could NOT call this method for the collection itself
func (s *HashSet[T]) RemoveAll(c collection.Collection[T]) bool {
	b := false
	if s != c {
		it := c.Iterate()
		for it.HasNext() {
			b = s.Remove(it.Next()) || b
		}
	}
	return b
}

// Remove all elements from this set
func (s *HashSet[T]) Clear() {
	s.mmap.Clear()
}

// Return the string that describes the contains of this list
func (s *HashSet[T]) String() string {
	return fmt.Sprint(s.ToSlice())
}

// HashSetIterator The iterator struct
type HashSetIterator[T comparable] struct {
	values []T
	index  int
}

// HasNext Does the iterator have more elements?
func (it *HashSetIterator[T]) HasNext() bool {
	if it.index+1 < len(it.values) {
		return true
	}
	return false
}

// Next Returns the next element of the iterator
func (it *HashSetIterator[T]) Next() T {
	if it.HasNext() {
		it.index++
		return it.values[it.index]
	}
	var zero T
	return zero
}
