// gcoll v1.0
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)
package gcoll

import (
	"github.com/vence722/gcoll/list"
	"github.com/vence722/gcoll/maps"
	"github.com/vence722/gcoll/set"
	"github.com/vence722/gcoll/tree"
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

func NewBSTreeMap() maps.StringMap {
	return tree.NewBinarySortTree()
}

func NewHashSet() set.Set {
	return set.NewHashSet()
}
