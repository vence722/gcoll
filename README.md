gcoll --- go collection library
=====

A collection library of the most frequently used data structures in Go programing language.

Installation
-----

```
go get github.com/vence722/gcoll
```

Usage
-----


Use the constructor functions that defines in gcoll namespace to create collection classes. The naming pattern is very straightforward.For example:

```
list := gcoll.NewLinkedList()
```

It creates a linked-list collection. Then you can use the collection class by calling its methods in a much object-oriented way.

```
// Insert element
list.Add(3)
// Retrieve element
obj := list.Get(0)
// Remove element
list.Remove(obj)
// Get the size of the collection
s := list.Size()
```

Now enjoy your programming!

Supported Collection Interfaces & Implementation Classes
-----

List: ArrayList, LinkedList, Queue, Stack

Queue: LinkedList

Stack: LinkedList

Set: HashSet

Map: HashMap, BinarySortTree

Heap: ArrayHeap

LRUCache: FifoLRUCache, HitsMapLRUCache

