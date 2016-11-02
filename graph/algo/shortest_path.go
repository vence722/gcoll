package algo

import (
	"math"

	"github.com/vence722/gcoll/graph"
	"github.com/vence722/gcoll/list"
)

func Dijkstra(wg graph.WeightedGraph, src graph.Vertex, dest graph.Vertex) (graph.Path, error) {
	srcIndex := indexVertex(wg, src)
	destIndex := indexVertex(wg, dest)
	if srcIndex == -1 || destIndex == -1 {
		return nil, graph.ERR_VERTEX_NOT_EXISTS
	}
	shortestPath := graph.NewSimplePath()
	shortestPath.AddNode(src.Key(), src.Value(), 0)
	shortestEdges := []graph.WeightedEdge{}

	numVertics := len(wg.Vertices())

	// calculate the adjacency matrix
	adjMatrix := wg.GetAdjacencyMatrix()
	// initialize distance vector(distance to MST)
	distVec := make([]float64, numVertics)
	for i := 0; i < numVertics; i++ {
		distVec[i] = math.MaxFloat64
	}
	distVec[srcIndex] = 0
	// the vector to store the status whether the vertics are used
	usedVec := make([]bool, numVertics)

	// loop until distance vector not changed
	for {
		// look for the current nearest vertex to src that is not used
		var nearest int = -1
		var distNearest float64 = math.MaxFloat64
		for i := 0; i < numVertics; i++ {
			if !usedVec[i] && distVec[i] < distNearest {
				nearest = i
				distNearest = distVec[i]
			}
		}

		// distVec update flag
		distVecUpdated := false

		if nearest != -1 {
			// update distance vector after choosing the nearest vertex
			for i := 0; i < numVertics; i++ {
				if usedVec[i] {
					continue
				}

				if adjMatrix.Get(nearest, i) == nil {
					continue
				}

				// update distance if the path through "nearest" vertex is nearer than the direct path
				if distVec[nearest]+adjMatrix.Get(nearest, i).(*graph.SimpleWeightedEdge).Weight() < distVec[i] {
					distVec[i] = distVec[nearest] + adjMatrix.Get(nearest, i).(*graph.SimpleWeightedEdge).Weight()
					distVecUpdated = true
				}
			}

			// mark vertex used
			usedVec[nearest] = true

			// Add path information
			vertex := wg.Vertices()[nearest]
			lastKey := shortestPath.Dest().Key()
			lastVertex := wg.GetVertex(lastKey)

			if vertex == nil || vertex.Key() == nil || vertex.Key() == lastVertex.Key() {
				continue
			}

			edge := wg.GetWeightedEdge(lastVertex, vertex)
			if edge == nil {
				return nil, graph.ERR_UNKNOWN
			}
			shortestEdges = append(shortestEdges, edge)
		}

		// if no change in distVec, break the loop
		if !distVecUpdated {
			break
		}
	}

	// convert shortestEdges to shortestPath
	stack := list.NewLinkedList()
	var currKey string = dest.Key().(string)
	for currKey != src.Key() {
		for _, edge := range shortestEdges {
			if edge.To().Key() == currKey {
				stack.Push(edge)
				currKey = edge.From().Key().(string)
				break
			}
		}
	}
	for !stack.IsEmpty() {
		edge := stack.Pop().(graph.WeightedEdge)
		shortestPath.AddNode(edge.To().Key(), edge.To().Value(), edge.Weight())
	}

	return shortestPath, nil
}
