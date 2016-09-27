package graph

import (
	"errors"
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
	// Add neighbour
	svx.neighbors = append(svx.neighbors, svy)
	svy.neighbors = append(svy.neighbors, svx)
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
