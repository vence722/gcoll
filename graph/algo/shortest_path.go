package algo

import (
	"math"
	"reflect"

	"gcoll/graph"
	"gcoll/list"
)

type traceEntry struct {
	fromIndex int
	toIndex   int
}

func Dijkstra(wg graph.WeightedGraph, src graph.Vertex, dest graph.Vertex) (graph.Path, error) {
	srcIndex := indexVertex(wg, src)
	destIndex := indexVertex(wg, dest)
	if srcIndex == -1 || destIndex == -1 {
		return nil, graph.ERR_VERTEX_NOT_EXISTS
	}
	shortestPath := graph.NewSimplePath()
	shortestPath.AddNode(src.Key(), src.Value(), 0)

	numVertics := len(wg.Vertices())

	// trace vector for tracing the shortest path
	traceVec := make([]int, numVertics)
	for i := 0; i < len(traceVec); i++ {
		traceVec[i] = -1
	}

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
				wEdge := adjMatrix.Get(nearest, i).(graph.WeightedEdge)
				if wEdge != nil && !reflect.ValueOf(wEdge).IsNil() && distVec[nearest]+wEdge.Weight() < distVec[i] {
					distVec[i] = distVec[nearest] + adjMatrix.Get(nearest, i).(*graph.SimpleWeightedEdge).Weight()

					// update trace vector
					if i != nearest {
						traceVec[i] = nearest
					}

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
		}

		// if no change in distVec, break the loop
		if !distVecUpdated {
			break
		}
	}

	// convert traceVec to shortestPath
	stack := list.NewLinkedList()
	toIndex := destIndex
	fromIndex := traceVec[toIndex]
	for fromIndex != -1 {
		stack.Push(&traceEntry{fromIndex: fromIndex, toIndex: toIndex})
		toIndex = fromIndex
		fromIndex = traceVec[toIndex]
	}

	for !stack.IsEmpty() {
		entry := stack.Pop().(*traceEntry)
		fromVertex := wg.Vertices()[entry.fromIndex]
		toVertex := wg.Vertices()[entry.toIndex]
		edge := wg.GetWeightedEdge(fromVertex, toVertex)

		var weight float64 = 0
		if edge != nil {
			weight = edge.Weight()
		}

		shortestPath.AddNode(toVertex.Key(), toVertex.Value(), weight)
	}

	return shortestPath, nil
}
