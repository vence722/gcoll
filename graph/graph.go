// gcoll
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)
package graph

import (
	"github.com/vence722/gcoll/matrix"
)

// Basic Graph
type Graph interface {
	Vertices() []Vertex
	Edges() []Edge
	GetVertex(key interface{}) Vertex
	AddVertex(key interface{}, value interface{}) (Vertex, error)
	RemoveVertex(key interface{}) error
	GetEdge(x Vertex, y Vertex) Edge
	AddEdge(x Vertex, y Vertex, value interface{}) (Edge, error)
	RemoveEdge(x Vertex, y Vertex) error
	IterateByBFS(startKey interface{}) GraphIterator
	IterateByDFS(startKey interface{}) GraphIterator
}

type Vertex interface {
	Key() interface{}
	Value() interface{}
	Neighbors() []Vertex
	AddNeighbor(neighbor Vertex) error
	RemoveNeighbor(neighbor Vertex) error
}

type Edge interface {
	Value() interface{}
	From() Vertex
	To() Vertex
}

type GraphIterator interface {
	HasNext() bool
	Next() Vertex
}

// Weighted Graph
type WeightedGraph interface {
	Graph
	WeightedEdges() []WeightedEdge
	GetWeightedEdge(x Vertex, y Vertex) WeightedEdge
	AddEdgeWithWeight(x Vertex, y Vertex, value interface{}, weight float64) (WeightedEdge, error)
	GetAdjacencyMatrix() matrix.Matrix
}

type WeightedEdge interface {
	Edge
	Weight() float64
}

// Path (Can be considered as graph with N vertics and N-1 edges)
type Path interface {
	Src() Node
	Dest() Node
	Nodes() []Node
	AddNode(key interface{}, value interface{}, weight float64) error
	GetNode(key interface{}) Node
	TotalWeight() float64
}

type Node interface {
	Key() interface{}
	Value() interface{}
	Next() Node
	WeightToPrev() float64
}
