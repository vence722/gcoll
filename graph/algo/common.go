package algo

import (
	"github.com/vence722/gcoll/graph"
)

func indexVertex(g graph.WeightedGraph, v graph.Vertex) int {
	for i, vertex := range g.Vertices() {
		if vertex.Key() == v.Key() {
			return i
		}
	}
	return -1
}
