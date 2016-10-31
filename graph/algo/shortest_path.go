package algo

import (
	//	"math"

	"github.com/vence722/gcoll/graph"
)

func Dijkstra(wg graph.WeightedGraph, src graph.Vertex, dest graph.Vertex) (graph.Path, error) {
	srcIndex := indexVertex(wg, src)
	destIndex := indexVertex(wg, dest)
	if srcIndex == -1 || destIndex == -1 {
		return nil, graph.ERR_VERTEX_NOT_EXISTS
	}
	shortestPath := graph.NewSimplePath()
	shortestPath.AddNode(src.Key(), src.Value(), 0)

	//	numVertics := len(wg.Vertices())

	// calculate the adjacency matrix
	//	adjMatrix := wg.GetAdjacencyMatrix()
	// initialize distance vector(distance to MST)
	//	distVec := make([]float64, numVertics)
	//
	//	usedVec := make([]bool, numVertics)
	//	for i := 0; i < numVertics; i++ {
	//		distVec[i] = math.MaxFloat64
	//	}
	//	distVec[srcIndex] = 0

	// Loop until distance vector not changed
	//	for {
	//		// look for the current nearest vertex to src that is not used
	//		var nearest int = -1
	//		var distNearest float64 = math.MaxFloat64
	//		for i := 0; i < numVertics; i++ {
	//			if !usedVec[i] && distVec[i] < distNearest {
	//				nearest = i
	//				distNearest = distVec[i]
	//			}
	//		}

	//	}

	return shortestPath, nil
}
