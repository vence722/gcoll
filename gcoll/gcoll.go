// gcoll v1.0
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)
package gcoll

import (
	"gcoll/list"
	"gcoll/maps"
	"gcoll/set"
)

func NewArrayList() list.List {
	return list.NewArrayList()
}

func NewLinkedList() list.List {
	return list.NewLinkedList()
}

func NewQueue() list.Queue {
	return list.NewLinkedList()
}

func NewStack() list.Stack {
	return list.NewLinkedList()
}

func NewHashMap() maps.Map {
	return maps.NewHashMap()
}

func NewHashSet() set.Set {
	return set.NewHashSet()
}
