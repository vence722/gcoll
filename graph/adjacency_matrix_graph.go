package graph

import (
	"fmt"

	"gcoll/matrix"
)

type AdjacencyMatrixGraph struct {
	adjacencyMatrix matrix.Matrix
	vertics         []Vertex
}

func NewAdjacencyMatrixGraph() *AdjacencyMatrixGraph {
	return &AdjacencyMatrixGraph{adjacencyMatrix: matrix.NewLinkedMatrix(0, 0, nil)}
}

func (this *AdjacencyMatrixGraph) Vertices() []Vertex {
	return this.vertics
}

func (this *AdjacencyMatrixGraph) Edges() []Edge {
	var edges []Edge
	width, height := this.adjacencyMatrix.Size()
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if i == j {
				continue
			}
			val := this.adjacencyMatrix.Get(i, j)
			if val == nil {
				continue
			}
			edges = append(edges, val.(Edge))
		}
	}
	return edges
}

func (this *AdjacencyMatrixGraph) WeightedEdges() []WeightedEdge {
	var edges []WeightedEdge
	width, height := this.adjacencyMatrix.Size()
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if i == j {
				continue
			}
			val := this.adjacencyMatrix.Get(i, j)
			if val == nil {
				continue
			}
			edges = append(edges, val.(WeightedEdge))
		}
	}
	return edges
}

func (this *AdjacencyMatrixGraph) GetVertex(key any) Vertex {
	for _, v := range this.vertics {
		if v.Key() == key {
			return v
		}
	}
	return nil
}

func (this *AdjacencyMatrixGraph) AddVertex(key any, value any) (Vertex, error) {
	if this.GetVertex(key) != nil {
		return nil, ERR_VERTEX_KEY_EXISTS
	}
	v := &AdjacencyMatrixVertex{key: key, value: value}
	this.vertics = append(this.vertics, v)

	// Need to expand the adjacency matrix
	sizeRows, sizeCols := this.adjacencyMatrix.Size()
	this.adjacencyMatrix.Resize(sizeRows+1, sizeCols+1, nil)

	return v, nil
}

func (this *AdjacencyMatrixGraph) RemoveVertex(key any) error {
	var index int = -1
	for i, v := range this.vertics {
		if v.Key() == key {
			index = i
			this.vertics = append(this.vertics[:i], this.vertics[i+1:]...)
			break
		}
	}
	if index == -1 {
		return ERR_VERTEX_KEY_NOT_EXISTS
	}
	// update adjacency matrix
	x, _ := this.adjacencyMatrix.Size()
	for i := 0; i < x; i++ {
		this.adjacencyMatrix.RemoveAt(index)
	}
	return nil
}

func (this *AdjacencyMatrixGraph) GetEdge(x Vertex, y Vertex) Edge {
	xIndex, yIndex := this.indexEdge(x, y)
	if xIndex == -1 || yIndex == -1 {
		return nil
	}
	return this.adjacencyMatrix.Get(xIndex, yIndex).(Edge)
}

func (this *AdjacencyMatrixGraph) GetWeightedEdge(x Vertex, y Vertex) WeightedEdge {
	xIndex, yIndex := this.indexEdge(x, y)
	if xIndex == -1 || yIndex == -1 {
		return nil
	}
	if this.adjacencyMatrix.Get(xIndex, yIndex) == nil {
		return nil
	}
	return this.adjacencyMatrix.Get(xIndex, yIndex).(WeightedEdge)
}

func (this *AdjacencyMatrixGraph) AddEdge(x Vertex, y Vertex, value any) (Edge, error) {
	return this.AddEdgeWithWeight(x, y, value, 0)
}

func (this *AdjacencyMatrixGraph) AddEdgeWithWeight(x Vertex, y Vertex, value any, weight float64) (WeightedEdge, error) {
	xIndex, yIndex := this.indexEdge(x, y)
	if xIndex == -1 || yIndex == -1 {
		return nil, ERR_INDEX_OUT_OF_BOUND
	}
	edge := &SimpleWeightedEdge{SimpleEdge: &SimpleEdge{from: x, to: y, value: value}, weight: weight}
	this.adjacencyMatrix.Set(xIndex, yIndex, edge)
	return edge, nil
}

func (this *AdjacencyMatrixGraph) RemoveEdge(x Vertex, y Vertex) error {
	return nil
}

func (this *AdjacencyMatrixGraph) IterateByBFS(startKey any) GraphIterator {
	return nil
}

func (this *AdjacencyMatrixGraph) IterateByDFS(startKey any) GraphIterator {
	return nil
}

func (this *AdjacencyMatrixGraph) GetAdjacencyMatrix() matrix.Matrix {
	return this.adjacencyMatrix
}

func (this *AdjacencyMatrixGraph) String() string {
	str := ""
	for _, edge := range this.Edges() {
		str += fmt.Sprintf("%s--->%s, weight:%f\n", edge.From().Key(), edge.To().Key(), edge.(*SimpleWeightedEdge).weight)
	}
	return str
}

func (this *AdjacencyMatrixGraph) indexVertex(key any) int {
	for i, v := range this.vertics {
		if v.Key() == key {
			return i
		}
	}
	return -1
}

func (this *AdjacencyMatrixGraph) indexEdge(x Vertex, y Vertex) (int, int) {
	xIndex := this.indexVertex(x.Key())
	yIndex := this.indexVertex(y.Key())
	return xIndex, yIndex
}

type AdjacencyMatrixVertex struct {
	key       any
	value     any
	neighbors []Vertex
}

func (this *AdjacencyMatrixVertex) Key() any {
	return this.key
}

func (this *AdjacencyMatrixVertex) Value() any {
	return this.value
}

func (this *AdjacencyMatrixVertex) Neighbors() []Vertex {
	return this.neighbors
}

func (this *AdjacencyMatrixVertex) AddNeighbor(neighbor Vertex) error {
	if neighbor == nil {
		return ERR_NEIGHBOR_IS_NULL
	}
	this.neighbors = append(this.neighbors, neighbor)
	return nil
}

func (this *AdjacencyMatrixVertex) RemoveNeighbor(neighbor Vertex) error {
	for i, n := range this.neighbors {
		if n.Key() == neighbor.Key() {
			this.neighbors = append(this.neighbors[:i], this.neighbors[i+1:]...)
			return nil
		}
	}
	return ERR_NEIGHBOR_NOT_EXISTS
}
