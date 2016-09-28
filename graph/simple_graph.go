package graph

import (
	"errors"

	"github.com/vence722/gcoll/list"
	"github.com/vence722/gcoll/set"
)

const (
	ITERATE_METHOD_DFS string = "DFS"
	ITERATE_METHOD_BFS string = "BFS"
)

var (
	ERR_VERTEX_KEY_EXISTS     error = errors.New("Vertex key exists")
	ERR_VERTEX_KEY_NOT_EXISTS error = errors.New("Vertex key not exists")
	ERR_EDGE_EXISTS           error = errors.New("Edge exists")
	ERR_EDGE_NOT_EXISTS       error = errors.New("Edge not exists")
)

type SimpleGraph struct {
	vertics []*SimpleVertex
	edges   []*SimpleEdge
}

func (this *SimpleGraph) Vertices() []Vertex {
	rtns := make([]Vertex, 0, len(this.vertics))
	for _, v := range this.vertics {
		rtns = append(rtns, v)
	}
	return rtns
}

func (this *SimpleGraph) Edges() []Edge {
	rtns := make([]Edge, 0, len(this.edges))
	for _, e := range this.edges {
		rtns = append(rtns, e)
	}
	return rtns
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
	v := &SimpleVertex{key: key, value: value, neighbors: []*SimpleVertex{}}
	this.vertics = append(this.vertics, v)
	return v, nil
}

// Remove vertex with specified key, and delete the edges related to it
func (this *SimpleGraph) RemoveVertex(key interface{}) error {
	var vertex *SimpleVertex = nil
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
	for _, n := range vertex.neighbors {
		// delete edges for both directions
		this.RemoveEdge(vertex, n)
		this.RemoveEdge(n, vertex)
	}
	return nil
}

func (this *SimpleGraph) GetEdge(x Vertex, y Vertex) Edge {
	for _, e := range this.edges {
		if e.from.Key() == x.Key() && e.to.Key() == y.Key() {
			return e
		}
	}
	return nil
}

func (this *SimpleGraph) AddEdge(x Vertex, y Vertex, value interface{}) (Edge, error) {
	if this.GetEdge(x, y) != nil {
		return nil, ERR_EDGE_EXISTS
	}
	svx := x.(*SimpleVertex)
	svy := y.(*SimpleVertex)
	e := &SimpleEdge{from: svx, to: svy, value: value}
	this.edges = append(this.edges, e)
	// Add neighbour if not exists
	var exists = false
	for _, n := range svx.neighbors {
		if n.Key() == svy.Key() {
			exists = true
		}
	}
	if !exists {
		svx.neighbors = append(svx.neighbors, svy)
	}

	exists = false
	for _, n := range svy.neighbors {
		if n.Key() == svx.Key() {
			exists = true
		}
	}
	if !exists {
		svy.neighbors = append(svy.neighbors, svx)
	}

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
		startVertex = this.vertics[0]
	}
	return newSimpleGraphIterator(ITERATE_METHOD_BFS, startVertex)
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
		startVertex = this.vertics[0]
	}
	return newSimpleGraphIterator(ITERATE_METHOD_DFS, startVertex)
}

func (this *SimpleGraph) delEdge(x Vertex, y Vertex) error {
	if this.GetEdge(x, y) == nil {
		return ERR_EDGE_NOT_EXISTS
	}
	for i, e := range this.edges {
		if e.from.Key() == x.Key() && e.to.Key() == y.Key() {
			this.edges = append(this.edges[:i], this.edges[i+1:]...)
			break
		}
	}
	return nil
}

func (this *SimpleGraph) delNeighbors(x Vertex, y Vertex) {
	svx := x.(*SimpleVertex)
	svy := y.(*SimpleVertex)
	for i, n := range svx.neighbors {
		if n.Key() == svy.Key() {
			svx.neighbors = append(svx.neighbors[:i], svx.neighbors[i+1:]...)
			break
		}
	}
	for i, n := range svy.neighbors {
		if n.Key() == svx.Key() {
			svy.neighbors = append(svy.neighbors[:i], svy.neighbors[i+1:]...)
			break
		}
	}
}

func NewSimpleGraph() *SimpleGraph {
	return &SimpleGraph{}
}

type SimpleVertex struct {
	key       interface{}
	value     interface{}
	neighbors []*SimpleVertex
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

type SimpleEdge struct {
	from  *SimpleVertex
	to    *SimpleVertex
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
	currVertex     *SimpleVertex
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
			this.currVertex = this.findNextNonTraveledVertex(ITERATE_METHOD_BFS)
			if this.currVertex != nil {
				// put the neighbors of the next vertex to the queue
				for _, n := range this.currVertex.neighbors {
					this.bfsQueue.EnQueue(n)
				}
			}
		} else {
			this.currVertex = nil
		}
	}

	if this.method == ITERATE_METHOD_DFS {
		if this.dfsStack.Size() > 0 {
			this.currVertex = this.findNextNonTraveledVertex(ITERATE_METHOD_DFS)
			if this.currVertex != nil {
				// put the neighbors of the next vertex to the stack
				for _, n := range this.currVertex.neighbors {
					this.dfsStack.Push(n)
				}
			}
		}
	}

	return next
}

func (this *SimpleGraphIterator) findNextNonTraveledVertex(method string) *SimpleVertex {
	var next *SimpleVertex
	for {
		if method == ITERATE_METHOD_BFS {
			if this.bfsQueue.Front() != nil {
				next = this.bfsQueue.DeQueue().(*SimpleVertex)
				if !this.visitedVertics.Contains(next) {
					break
				}
			} else {
				next = nil
				break
			}
		} else if method == ITERATE_METHOD_DFS {
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

func newSimpleGraphIterator(method string, startFrom *SimpleVertex) *SimpleGraphIterator {
	iter := &SimpleGraphIterator{currVertex: startFrom, visitedVertics: set.NewHashSet(), method: method}
	// For BFS, initialize the queue
	if method == ITERATE_METHOD_BFS {
		iter.bfsQueue = list.NewLinkedList()
		if startFrom != nil {
			for _, n := range startFrom.neighbors {
				iter.bfsQueue.EnQueue(n)
			}
		}
	} else if method == ITERATE_METHOD_DFS {
		iter.dfsStack = list.NewLinkedList()
		if startFrom != nil {
			for _, n := range startFrom.neighbors {
				iter.dfsStack.Push(n)
			}
		}
	}
	return iter
}
