package graph

type Graph interface {
	Vertices() []Vertex
	Edges() []Edge
	GetVertex(key interface{}) Vertex
	AddVertex(key interface{}, value interface{}) (Vertex, error)
	RemoveVertex(key interface{}) error
	GetEdge(x Vertex, y Vertex) Edge
	AddEdge(x Vertex, y Vertex, value interface{}) (Edge, error)
	RemoveEdge(x Vertex, y Vertex) error
}

type Vertex interface {
	Key() interface{}
	Value() interface{}
	Neighbors() []Vertex
}

type Edge interface {
	Value() interface{}
	From() Vertex
	To() Vertex
}
