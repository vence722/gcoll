// gcoll
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)
package graph

import (
	"fmt"

	"github.com/vence722/gcoll/list"
	"github.com/vence722/gcoll/set"
)

// Implementation of Simple Graph
// It is flexible for other implementation to "extend"
type SimpleGraph struct {
	vertics []Vertex
	edges   []Edge
}

func (this *SimpleGraph) Vertices() []Vertex {
	return this.vertics
}

func (this *SimpleGraph) Edges() []Edge {
	return this.edges
}

func (this *SimpleGraph) GetVertex(key interface{}) Vertex {
	for _, vertex := range this.vertics {
		if vertex.Key() == key {
			return vertex
		}
	}
	return nil
}

func (this *SimpleGraph) AddVertex(key interface{}, value interface{}) (Vertex, error) {
	if this.GetVertex(key) != nil {
		return nil, ERR_VERTEX_KEY_EXISTS
	}
	v := &SimpleVertex{key: key, value: value, neighbors: []Vertex{}}
	this.vertics = append(this.vertics, v)
	return v, nil
}

// Remove vertex with specified key, and delete the edges related to it
func (this *SimpleGraph) RemoveVertex(key interface{}) error {
	var vertex Vertex = nil
	for i, v := range this.vertics {
		if v.Key() == key {
			vertex = v
			this.vertics = append(this.vertics[:i], this.vertics[i+1:]...)
			break
		}
	}
	if vertex == nil {
		return ERR_VERTEX_KEY_NOT_EXISTS
	}
	for _, n := range vertex.Neighbors() {
		// delete edges for both directions
		this.RemoveEdge(vertex, n)
		this.RemoveEdge(n, vertex)
	}
	return nil
}

func (this *SimpleGraph) GetEdge(x Vertex, y Vertex) Edge {
	for _, e := range this.edges {
		if e.From().Key() == x.Key() && e.To().Key() == y.Key() {
			return e
		}
	}
	return nil
}

func (this *SimpleGraph) AddEdge(x Vertex, y Vertex, value interface{}) (Edge, error) {
	if this.GetEdge(x, y) != nil {
		return nil, ERR_EDGE_EXISTS
	}
	e := &SimpleEdge{from: x, to: y, value: value}
	this.addEdgeInternal(e)
	return e, nil
}

func (this *SimpleGraph) RemoveEdge(x Vertex, y Vertex) error {
	// delete edge
	if err := this.delEdge(x, y); err != nil {
		return err
	}
	// if edges for both directions are deleted
	// delete neighbors relationship
	if this.GetEdge(y, x) == nil {
		this.delNeighbors(x, y)
	}
	return nil
}

// Breadth-first Search Iteration
func (this *SimpleGraph) IterateByBFS(startKey interface{}) GraphIterator {
	var startVertex *SimpleVertex
	if startKey != nil {
		v := this.GetVertex(startKey)
		if v != nil {
			startVertex = v.(*SimpleVertex)
		}
	}
	// If start key not specified or start vertex not found, use first vertex in the list
	if startKey == nil && len(this.vertics) > 0 {
		startVertex = this.vertics[0].(*SimpleVertex)
	}
	return newSimpleGraphIterator(this, ITERATE_METHOD_BFS, startVertex)
}

// Depth-First Search Iteration
func (this *SimpleGraph) IterateByDFS(startKey interface{}) GraphIterator {
	var startVertex *SimpleVertex
	if startKey != nil {
		v := this.GetVertex(startKey)
		if v != nil {
			startVertex = v.(*SimpleVertex)
		}
	}
	// If start key not specified or start vertex not found, use first vertex in the list
	if startKey == nil && len(this.vertics) > 0 {
		startVertex = this.vertics[0].(*SimpleVertex)
	}
	return newSimpleGraphIterator(this, ITERATE_METHOD_DFS, startVertex)
}

func (this *SimpleGraph) String() string {
	str := ""
	for _, edge := range this.edges {
		str += fmt.Sprintf("%s--->%s\n", edge.From().Key(), edge.To().Key())
	}
	return str
}

func (this *SimpleGraph) addEdgeInternal(edge Edge) {
	this.edges = append(this.edges, edge)
	svx := edge.From()
	svy := edge.To()
	// Add neighbour if not exists
	var exists = false
	for _, n := range svx.Neighbors() {
		if n.Key() == svy.Key() {
			exists = true
		}
	}
	if !exists {
		svx.AddNeighbor(svy)
	}

	exists = false
	for _, n := range svy.Neighbors() {
		if n.Key() == svx.Key() {
			exists = true
		}
	}
	if !exists {
		svy.AddNeighbor(svx)
	}
}

func (this *SimpleGraph) delEdge(x Vertex, y Vertex) error {
	if this.GetEdge(x, y) == nil {
		return ERR_EDGE_NOT_EXISTS
	}
	for i, e := range this.edges {
		if e.From().Key() == x.Key() && e.To().Key() == y.Key() {
			this.edges = append(this.edges[:i], this.edges[i+1:]...)
			break
		}
	}
	return nil
}

func (this *SimpleGraph) delNeighbors(x Vertex, y Vertex) {
	x.RemoveNeighbor(y)
	y.RemoveNeighbor(x)
}

func NewSimpleGraph() *SimpleGraph {
	return &SimpleGraph{}
}

// Implementation of Simple Vertex
type SimpleVertex struct {
	key       interface{}
	value     interface{}
	neighbors []Vertex
}

func (this *SimpleVertex) Key() interface{} {
	return this.key
}

func (this *SimpleVertex) Value() interface{} {
	return this.value
}

func (this *SimpleVertex) Neighbors() []Vertex {
	rtns := make([]Vertex, 0, len(this.neighbors))
	for _, neighbor := range this.neighbors {
		rtns = append(rtns, neighbor)
	}
	return rtns
}

func (this *SimpleVertex) AddNeighbor(neighbor Vertex) error {
	if neighbor == nil {
		return ERR_NEIGHBOR_IS_NULL
	}
	this.neighbors = append(this.neighbors, neighbor)
	return nil
}

func (this *SimpleVertex) RemoveNeighbor(neighbor Vertex) error {
	for i, n := range this.neighbors {
		if n.Key() == neighbor.Key() {
			this.neighbors = append(this.neighbors[:i], this.neighbors[i+1:]...)
			return nil
		}
	}
	return ERR_NEIGHBOR_NOT_EXISTS
}

// Implementation of Simple Edge
type SimpleEdge struct {
	from  Vertex
	to    Vertex
	value interface{}
}

func (this *SimpleEdge) Value() interface{} {
	return this.value
}

func (this *SimpleEdge) From() Vertex {
	return this.from
}

func (this *SimpleEdge) To() Vertex {
	return this.to
}

// Implementation of the GraphIterator for SimpleGraph
type SimpleGraphIterator struct {
	graph          *SimpleGraph
	currVertex     Vertex
	visitedVertics set.Set
	method         string
	bfsQueue       list.Queue
	dfsStack       list.Stack
}

func (this *SimpleGraphIterator) HasNext() bool {
	if this.currVertex != nil {
		return true
	}
	return false
}

func (this *SimpleGraphIterator) Next() Vertex {
	next := this.currVertex
	if next == nil {
		return next
	}
	// add the next vertex to visited set
	this.visitedVertics.Add(next)

	if this.method == ITERATE_METHOD_BFS {
		if this.bfsQueue.Size() > 0 {
			this.currVertex = this.findNextNonTraveledVertex()
			if this.currVertex != nil {
				// put the neighbors which the vertex has an edge leading to in the queue
				this.handleNeighborVertics(this.currVertex)
			}
		} else {
			this.currVertex = nil
		}
	}

	if this.method == ITERATE_METHOD_DFS {
		if this.dfsStack.Size() > 0 {
			this.currVertex = this.findNextNonTraveledVertex()
			if this.currVertex != nil {
				// put the neighbors which the vertex has an edge leading to in the stack
				this.handleNeighborVertics(this.currVertex)
			}
		} else {
			this.currVertex = nil
		}
	}

	return next
}

func (this *SimpleGraphIterator) findNextNonTraveledVertex() *SimpleVertex {
	var next *SimpleVertex
	for {
		if this.method == ITERATE_METHOD_BFS {
			if this.bfsQueue.Front() != nil {
				next = this.bfsQueue.DeQueue().(*SimpleVertex)
				if !this.visitedVertics.Contains(next) {
					break
				}
			} else {
				next = nil
				break
			}
		} else if this.method == ITERATE_METHOD_DFS {
			if this.dfsStack.Top() != nil {
				next = this.dfsStack.Pop().(*SimpleVertex)
				if !this.visitedVertics.Contains(next) {
					break
				}
			} else {
				next = nil
				break
			}
		}
	}
	return next
}

func (this *SimpleGraphIterator) handleNeighborVertics(vertex Vertex) {
	if this.method == ITERATE_METHOD_BFS {
		for _, n := range vertex.Neighbors() {
			if this.graph.GetEdge(vertex, n) != nil {
				this.bfsQueue.EnQueue(n)
			}
		}
	} else if this.method == ITERATE_METHOD_DFS {
		for _, n := range vertex.Neighbors() {
			if this.graph.GetEdge(vertex, n) != nil {
				this.dfsStack.Push(n)
			}
		}
	}
}

func newSimpleGraphIterator(graph *SimpleGraph, method string, startFrom *SimpleVertex) *SimpleGraphIterator {
	iter := &SimpleGraphIterator{graph: graph, currVertex: startFrom, visitedVertics: set.NewHashSet(), method: method}
	// For BFS, initialize the queue
	if method == ITERATE_METHOD_BFS {
		iter.bfsQueue = list.NewLinkedList()
	} else if method == ITERATE_METHOD_DFS {
		iter.dfsStack = list.NewLinkedList()
	}
	if startFrom != nil {
		iter.handleNeighborVertics(iter.currVertex)
	}
	return iter
}
