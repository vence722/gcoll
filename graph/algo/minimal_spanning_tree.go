package algo

import (
	"math"
	"reflect"

	"gcoll/graph"
)

// Build the minimal spanning tree from the graph with Prim Algorithm
// Started from a specific vertex
// *** The graph should be UNDIRECTED ***
// Returns the retult tree and the total minimal weight
func CreateMinimalSpanningTree(g graph.WeightedGraph, root graph.Vertex) (graph.WeightedGraph, float64, error) {
	var rootIndex int = indexVertex(g, root)
	var totalWeight float64 = 0
	if rootIndex == -1 {
		return nil, totalWeight, graph.ERR_VERTEX_NOT_EXISTS
	}
	mst := graph.NewSimpleWeightedGraph()
	mst.AddVertex(root.Key(), root.Value())
	// vector to remember the index of the vertics that added to minimal spanning tree
	indexVec := []int{rootIndex}
	// calculate the adjacency matrix
	adjMtrx := g.GetAdjacencyMatrix()
	// initialize distance vector(distance to MST)
	distVec := make([]float64, len(g.Vertices()))
	for i := 0; i < len(distVec); i++ {
		distVec[i] = math.MaxFloat64
	}
	distVec[rootIndex] = 0
	// Loop until distance vector not changed
	for {
		minDistToMST := math.MaxFloat64
		minDistIndex := -1
		minDistMSTIndex := -1
		for i, _ := range g.Vertices() {
			// distance > 0 means vertex is not already in MST
			if distVec[i] > 0 {
				// determine the distance between the vertex and the MST
				distToMST := math.MaxFloat64
				distMSTIndex := -1
				for _, index := range indexVec {
					wEdge := adjMtrx.Get(i, index).(graph.WeightedEdge)
					if wEdge != nil && !reflect.ValueOf(wEdge).IsNil() && wEdge.Weight() < distToMST {
						distToMST = wEdge.Weight()
						distMSTIndex = index
					}
				}
				// update distance vector
				distVec[i] = distToMST
				if distToMST < minDistToMST {
					minDistToMST = distToMST
					minDistIndex = i
					minDistMSTIndex = distMSTIndex
				}
			}
		}
		if minDistIndex != -1 {
			// add the nearest point to MST
			nearestVertex := g.Vertices()[minDistIndex]
			mst.AddVertex(nearestVertex.Key(), nearestVertex.Value())
			// add the shortest edge mutually
			shortestEdge := g.GetEdge(g.Vertices()[minDistIndex], g.Vertices()[minDistMSTIndex])
			mst.AddEdgeWithWeight(shortestEdge.From(), shortestEdge.To(), shortestEdge.Value(), minDistToMST)
			mst.AddEdgeWithWeight(shortestEdge.To(), shortestEdge.From(), shortestEdge.Value(), minDistToMST)
			// update total weight
			totalWeight = totalWeight + minDistToMST
			// mark the vertex already in MST
			distVec[minDistIndex] = 0
			// update indexVec
			indexVec = append(indexVec, minDistIndex)
		} else {
			// no new vertex needed to add, can exit the loop
			break
		}
	}
	return mst, totalWeight, nil
}
