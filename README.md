gcoll --- go collection library
=====

A collection library of the most frequently used data structures in Go programing language.

Requirements
-----
```
go version >= 1.8
```

gcoll version 2 is released, which supports generic types in Go 1.8.

Installation
-----

```
go get github.com/vence722/gcoll
```

Usage
-----


Use the constructor functions that defines in gcoll namespace to create collection objects. The names of the constructors are straightforward. For example:

```
list := gcoll.NewLinkedList[int]()
```

It creates a linked-list of integer type. Then you can use it in a very object-oriented way.

```
// Insert element
list.Add(3)
// Retrieve element
obj, ok := list.Get(0)
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

Map: HashMap

Heap: ArrayHeap

LRUCache: FifoLRUCache, HitsMapLRUCache

