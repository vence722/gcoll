// gcoll
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)
package graph

import (
	"errors"
	"fmt"
	"math"
)

var (
	ERR_ROOT_VERTEX_NOT_EXISTS error = errors.New("Root vertex not exists")
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

// Build the minimal spanning tree from the graph, started from a specific vertex
// Returns the retult tree and the total minimal weight
func (this *SimpleWeightedGraph) CreateMinimalSpanningTree(root Vertex) (WeightedGraph, float64, error) {
	if !this.checkVertex(root) {
		return nil, 0, ERR_ROOT_VERTEX_NOT_EXISTS
	}
	mst := NewSimpleWeightedGraph()
	mst.AddVertex(root.Key(), root.Value())

	adjMtrx := this.calculateAdjacencyMatrix()
	fmt.Println(adjMtrx)

	return nil, 0, nil
}

func (this *SimpleWeightedGraph) String() string {
	str := ""
	for _, edge := range this.edges {
		str += fmt.Sprintf("%s--->%s, weight:%f\n", edge.From().Key(), edge.To().Key(), edge.(*SimpleWeightedEdge).weight)
	}
	return str
}

func (this *SimpleWeightedGraph) checkVertex(v Vertex) bool {
	for _, vertex := range this.vertics {
		if vertex.Key() == v.Key() {
			return true
		}
	}
	return false
}

func (this *SimpleWeightedGraph) calculateAdjacencyMatrix() [][]float64 {
	size := len(this.vertics)
	// initialize adjacency matrix
	adjMtrx := make([][]float64, size)
	for i := 0; i < size; i++ {
		adjMtrx[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			adjMtrx[i][j] = math.MaxFloat64
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
		adjMtrx[fromIndex][toIndex] = wEdge.weight
	}
	return adjMtrx
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
