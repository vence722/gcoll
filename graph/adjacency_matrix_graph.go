package graph

import (
	"math"

	"github.com/vence722/gcoll/matrix"
)

type AdjacencyMatrixGraph struct {
	adjacencyMatrix matrix.Matrix
	vertics         []Vertex
}

func (this *AdjacencyMatrixGraph) Vertices() []Vertex {
	return this.vertics
}

func (this *AdjacencyMatrixGraph) Edges() []Edge {
	var edges []Edge
	return edges
}

func (this *AdjacencyMatrixGraph) GetVertex(key interface{}) Vertex {
	for _, v := range this.vertics {
		if v.Key() == key {
			return v
		}
	}
	return nil
}

func (this *AdjacencyMatrixGraph) AddVertex(key interface{}, value interface{}) (Vertex, error) {
	if this.GetVertex(key) != nil {
		return nil, ERR_VERTEX_KEY_EXISTS
	}
	v := &AdjacencyMatrixVertex{key: key, value: value}
	this.vertics = append(this.vertics, v)

	// Need to expand the adjacency matrix
	sizeRows, sizeCols := this.adjacencyMatrix.Size()
	this.adjacencyMatrix.Resize(sizeRows+1, sizeCols+1)
	sizeRows, sizeCols = this.adjacencyMatrix.Size()
	for i := 0; i < sizeRows; i++ {
		this.adjacencyMatrix.Set(i, sizeCols-1, math.MaxFloat64)
	}
	for i := 0; i < sizeCols; i++ {
		this.adjacencyMatrix.Set(sizeRows-1, i, math.MaxFloat64)
	}
	this.adjacencyMatrix.Set(sizeRows-1, sizeCols-1, 0)

	return v, nil
}

func (this *AdjacencyMatrixGraph) RemoveVertex(key interface{}) error {
	return nil
}

func (this *AdjacencyMatrixGraph) GetEdge(x Vertex, y Vertex) Edge {
	return nil
}

func (this *AdjacencyMatrixGraph) AddEdge(x Vertex, y Vertex, value interface{}) (Edge, error) {
	return nil, nil
}

func (this *AdjacencyMatrixGraph) RemoveEdge(x Vertex, y Vertex) error {
	return nil
}

func (this *AdjacencyMatrixGraph) IterateByBFS(startKey interface{}) GraphIterator {
	return nil
}

func (this *AdjacencyMatrixGraph) IterateByDFS(startKey interface{}) GraphIterator {
	return nil
}

func (this *AdjacencyMatrixGraph) indexVertex(key interface{}) int {
	for i, v := range this.vertics {
		if v.Key() == key {
			return i
		}
	}
	return -1
}

type AdjacencyMatrixVertex struct {
	key   interface{}
	value interface{}
}
