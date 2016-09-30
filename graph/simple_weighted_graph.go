// gcoll
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)
package graph

import (
	"fmt"
)

// Implementation of Simple Weighted Graph
type SimpleWeightedGraph struct {
	*SimpleGraph
}

func (this *SimpleWeightedGraph) AddEdgeWithWeight(x Vertex, y Vertex, value interface{}, weight float64) (WeightedEdge, error) {
	if this.GetEdge(x, y) != nil {
		return nil, ERR_EDGE_EXISTS
	}
	e := &SimpleWeightedEdge{SimpleEdge: &SimpleEdge{from: x, to: y, value: value}, weight: weight}
	this.addEdgeInternal(e)
	return e, nil
}

// Build the minimal spanning tree from the graph
// Returns the retult tree and the total minimal weight
func (this *SimpleWeightedGraph) CreateMinimalSpanningTree() (WeightedGraph, float64) {
	return nil, 0
}

func (this *SimpleWeightedGraph) String() string {
	str := ""
	for _, edge := range this.edges {
		str += fmt.Sprintf("%s--->%s, weight:%f\n", edge.From().Key(), edge.To().Key(), edge.(*SimpleWeightedEdge).weight)
	}
	return str
}

func NewSimpleWeightedGraph() WeightedGraph {
	return &SimpleWeightedGraph{&SimpleGraph{}}
}

// Implementation of Simple Weighted Edge
type SimpleWeightedEdge struct {
	*SimpleEdge
	weight float64
}

func (this *SimpleWeightedEdge) Weight() float64 {
	return this.weight
}
