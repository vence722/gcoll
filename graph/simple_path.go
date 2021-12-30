package graph

import (
	"fmt"
	"reflect"
)

type SimplePath struct {
	nodes []Node
	src   Node
	dest  Node
}

func NewSimplePath() *SimplePath {
	return &SimplePath{}
}

func (this *SimplePath) Src() Node {
	return this.src
}

func (this *SimplePath) Dest() Node {
	return this.dest
}

func (this *SimplePath) Nodes() []Node {
	return this.nodes
}

func (this *SimplePath) AddNode(key any, value any, weightToPrev float64) error {
	if this.GetNode(key) != nil {
		return ERR_NODE_EXISTS
	}
	node := &SimpleNode{key: key, value: value, weightToPrev: weightToPrev}
	this.nodes = append(this.nodes, node)
	if this.src == nil {
		this.src = node
		this.dest = node
	} else {
		dest := this.dest.(*SimpleNode)
		dest.next = node
		this.dest = node
	}
	return nil
}

func (this *SimplePath) GetNode(key any) Node {
	for _, node := range this.Nodes() {
		if node.Key() == key {
			return node
		}
	}
	return nil
}

func (this *SimplePath) String() string {
	var str string = ""
	node := this.Src().(*SimpleNode)
	next := node.Next().(*SimpleNode)
	for node != nil && next != nil {
		str += fmt.Sprintf("%s--->%s, weight:%f\n", node.Key(), next.Key(), next.WeightToPrev())
		node = next
		next = node.Next().(*SimpleNode)
	}
	return str
}

func (this *SimplePath) TotalWeight() float64 {
	var totalW float64 = 0
	curr := this.src
	for !reflect.ValueOf(curr).IsNil() {
		totalW += curr.WeightToPrev()
		curr = curr.Next()
	}
	return totalW
}

type SimpleNode struct {
	key          any
	value        any
	next         *SimpleNode
	weightToPrev float64
}

func (this *SimpleNode) Key() any {
	return this.key
}

func (this *SimpleNode) Value() any {
	return this.value
}

func (this *SimpleNode) Next() Node {
	return this.next
}

func (this *SimpleNode) WeightToPrev() float64 {
	return this.weightToPrev
}
