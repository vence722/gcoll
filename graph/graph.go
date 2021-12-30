// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)
package graph

import (
	"gcoll/matrix"
)

// Basic Graph
type Graph interface {
	Vertices() []Vertex
	Edges() []Edge
	GetVertex(key any) Vertex
	AddVertex(key any, value any) (Vertex, error)
	RemoveVertex(key any) error
	GetEdge(x Vertex, y Vertex) Edge
	AddEdge(x Vertex, y Vertex, value any) (Edge, error)
	RemoveEdge(x Vertex, y Vertex) error
	IterateByBFS(startKey any) GraphIterator
	IterateByDFS(startKey any) GraphIterator
}

type Vertex interface {
	Key() any
	Value() any
	Neighbors() []Vertex
	AddNeighbor(neighbor Vertex) error
	RemoveNeighbor(neighbor Vertex) error
}

type Edge interface {
	Value() any
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
	AddEdgeWithWeight(x Vertex, y Vertex, value any, weight float64) (WeightedEdge, error)
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
	AddNode(key any, value any, weight float64) error
	GetNode(key any) Node
	TotalWeight() float64
}

type Node interface {
	Key() any
	Value() any
	Next() Node
	WeightToPrev() float64
}
