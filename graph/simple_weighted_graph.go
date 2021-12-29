// gcoll
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)
package graph

import (
	"fmt"

	"gcoll/matrix"
)

// Implementation of Simple Weighted Graph
type SimpleWeightedGraph struct {
	*SimpleGraph
}

func (this *SimpleWeightedGraph) WeightedEdges() []WeightedEdge {
	wEdges := make([]WeightedEdge, len(this.vertics))
	for i, edge := range this.edges {
		wEdges[i] = edge.(WeightedEdge)
	}
	return wEdges
}

func (this *SimpleWeightedGraph) GetWeightedEdge(x Vertex, y Vertex) WeightedEdge {
	v := this.GetEdge(x, y)
	if v != nil {
		return v.(WeightedEdge)
	} else {
		return nil
	}
}

func (this *SimpleWeightedGraph) AddEdgeWithWeight(x Vertex, y Vertex, value interface{}, weight float64) (WeightedEdge, error) {
	if this.GetEdge(x, y) != nil {
		return nil, ERR_EDGE_EXISTS
	}
	e := &SimpleWeightedEdge{SimpleEdge: &SimpleEdge{from: x, to: y, value: value}, weight: weight}
	this.addEdgeInternal(e)
	return e, nil
}

func (this *SimpleWeightedGraph) String() string {
	str := ""
	for _, edge := range this.edges {
		str += fmt.Sprintf("%s--->%s, weight:%f\n", edge.From().Key(), edge.To().Key(), edge.(*SimpleWeightedEdge).weight)
	}
	return str
}

func (this *SimpleWeightedGraph) GetAdjacencyMatrix() matrix.Matrix {
	size := len(this.vertics)
	// initialize adjacency matrix
	adjMtrx := make([][]*SimpleWeightedEdge, size)
	for i := 0; i < size; i++ {
		adjMtrx[i] = make([]*SimpleWeightedEdge, size)
		for j := 0; j < len(adjMtrx[i]); j++ {
			adjMtrx[i][j] = nil
		}
	}
	// for each edge, update the adjacency matrix
	for _, edge := range this.edges {
		wEdge := edge.(*SimpleWeightedEdge)
		fromKey := wEdge.From().Key()
		toKey := wEdge.To().Key()
		fromIndex := -1
		toIndex := -1
		for i, vertex := range this.vertics {
			if fromKey == vertex.Key() {
				fromIndex = i
				continue
			}
			if toKey == vertex.Key() {
				toIndex = i
				continue
			}
			if fromIndex != -1 && toIndex != -1 {
				break
			}
		}
		adjMtrx[fromIndex][toIndex] = wEdge
	}

	mtrx := matrix.NewLinkedMatrix(size, size, nil)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			mtrx.Set(i, j, adjMtrx[i][j])
		}
	}

	return mtrx
}

func NewSimpleWeightedGraph() *SimpleWeightedGraph {
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
