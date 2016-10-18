// gcoll
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)
package graph

import (
	"fmt"
	"math"
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

// Build the minimal spanning tree from the graph with Prim Algorithm
// Started from a specific vertex
// *** The graph should be UNDIRECTED ***
// Returns the retult tree and the total minimal weight
func (this *SimpleWeightedGraph) CreateMinimalSpanningTree(root Vertex) (WeightedGraph, float64, error) {
	var rootIndex int = this.indexVertex(root)
	var totalWeight float64 = 0
	if rootIndex == -1 {
		return nil, totalWeight, ERR_ROOT_VERTEX_NOT_EXISTS
	}
	mst := NewSimpleWeightedGraph()
	mst.AddVertex(root.Key(), root.Value())
	// vector to remember the index of the vertics that added to minimal spanning tree
	indexVec := []int{rootIndex}
	// calculate the adjacency matrix
	adjMtrx := this.calculateAdjacencyMatrix()
	// initialize distance vector(distance to MST)
	distVec := make([]float64, len(this.edges))
	for i := 0; i < len(distVec); i++ {
		distVec[i] = math.MaxFloat64
	}
	distVec[rootIndex] = 0
	// Loop until distance vector not changed
	for {
		minDistToMST := math.MaxFloat64
		minDistIndex := -1
		minDistMSTIndex := -1
		for i, _ := range this.Vertices() {
			// distance > 0 means vertex is not already in MST
			if distVec[i] > 0 {
				// determine the distance between the vertex and the MST
				distToMST := math.MaxFloat64
				distMSTIndex := -1
				for _, index := range indexVec {
					if adjMtrx[i][index] < distToMST {
						distToMST = adjMtrx[i][index]
						distMSTIndex = index
					}
				}
				// update distance vector
				distVec[i] = distToMST
				if distToMST < minDistToMST {
					minDistToMST = distToMST
					minDistIndex = i
					minDistMSTIndex = distMSTIndex
				}
			}
		}
		if minDistIndex != -1 {
			// add the nearest point to MST
			nearestVertex := this.vertics[minDistIndex]
			mst.AddVertex(nearestVertex.Key(), nearestVertex.Value())
			// add the shortest edge mutually
			shortestEdge := this.GetEdge(this.vertics[minDistIndex], this.vertics[minDistMSTIndex])
			mst.AddEdgeWithWeight(shortestEdge.From(), shortestEdge.To(), shortestEdge.Value(), minDistToMST)
			mst.AddEdgeWithWeight(shortestEdge.To(), shortestEdge.From(), shortestEdge.Value(), minDistToMST)
			// update total weight
			totalWeight = totalWeight + minDistToMST
			// mark the vertex already in MST
			distVec[minDistIndex] = 0
			// update indexVec
			indexVec = append(indexVec, minDistIndex)
		} else {
			// no new vertex needed to add, can exit the loop
			break
		}
	}
	return mst, totalWeight, nil
}

func (this *SimpleWeightedGraph) String() string {
	str := ""
	for _, edge := range this.edges {
		str += fmt.Sprintf("%s--->%s, weight:%f\n", edge.From().Key(), edge.To().Key(), edge.(*SimpleWeightedEdge).weight)
	}
	return str
}

func (this *SimpleWeightedGraph) indexVertex(v Vertex) int {
	for i, vertex := range this.vertics {
		if vertex.Key() == v.Key() {
			return i
		}
	}
	return -1
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
