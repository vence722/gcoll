// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)
package gcoll

import (
	"github.com/stretchr/testify/assert"
	"github.com/vence722/gcoll/list"
	"strings"
	"testing"
)

func TestArrayList(t *testing.T) {
	arrayList := list.NewArrayList[string]()
	arrayList.Add("test1")
	arrayList.Add("test2")
	arrayList.Add("test3")

	assert.Equal(t, arrayList.Size(), 3)
	assert.Equal(t, arrayList.IsEmpty(), false)
	assert.Equal(t, arrayList.Contains("test1"), true)
	assert.Equal(t, arrayList.Contains("test4"), false)
	assert.Equal(t, arrayList.MustGet(0), "test1")

	it := arrayList.Iterate()
	for it.HasNext() {
		assert.Equal(t, strings.HasPrefix(it.Next(), "test"), true)
	}

	arrayList.RemoveAt(0)
	assert.Equal(t, arrayList.Size(), 2)
}

func TestLinkedList(t *testing.T) {
	lnkList := list.NewLinkedList[int]()
	lnkList.Add(3)
	lnkList.Add(2)
	lnkList.Set(1, 1)

	assert.Equal(t, lnkList.IsEmpty(), false)
	assert.Equal(t, lnkList.Size(), 2)
}

//
//func TestLRUCache(t *testing.T) {
//	// === Test Fifo LRU Cache ===
//	fifoLRU := lru.NewFifoLRUCache(5)
//	fifoLRU.Put(1, 'a')
//	fifoLRU.Put('a', 8)
//	fifoLRU.Put("3", 0x5F)
//	fifoLRU.Put("aa", "cc")
//	fifoLRU.Put("bb", 0x5F)
//
//	t.Log(fifoLRU.Size())
//	t.Log(fifoLRU.Get(1))
//
//	// Test insert more tan Cap
//	fifoLRU.Put("cc", 5214)
//
//	// Out of FIFO
//	t.Log(fifoLRU.Get(1))
//
//	fifoLRU.Clear()
//	t.Log(fifoLRU.Size())
//
//	// === Test HitsMap LRU Cache ===
//	hitsMapLRU := lru.NewHitsMapLRUCache(5)
//	hitsMapLRU.Put(1, 'a')
//	hitsMapLRU.Put('a', 8)
//	hitsMapLRU.Put("3", 0x5F)
//	hitsMapLRU.Put("aa", "cc")
//	hitsMapLRU.Put("bb", 0x5F)
//
//	t.Log(hitsMapLRU.Size())
//	t.Log(hitsMapLRU.Get(1))
//
//	// Get will add hits count
//	t.Log(hitsMapLRU.Get('a'))
//	t.Log(hitsMapLRU.Get("3"))
//	t.Log(hitsMapLRU.Get("bb"))
//
//	// Test insert more than Cap
//	hitsMapLRU.Put("cc", 5214)
//	// Item will least hit count removed
//	t.Log(fifoLRU.Get("aa"))
//
//	fifoLRU.Clear()
//	t.Log(fifoLRU.Size())
//}
//
//func TestTrees(t *testing.T) {
//	// BST
//	bst := tree.NewBinarySortTree()
//	bst.Put("100", 100)
//	bst.Put("50", 50)
//	bst.Put("220", 220)
//	bst.Put("99", 99)
//
//	// Get
//	t.Log(bst.Get("50"))
//
//	// Keys
//	t.Log(bst.Keys())
//
//	// Values
//	t.Log(bst.Values())
//
//	// Root
//	t.Log(bst.Root())
//
//	// Remove
//	t.Log(bst.Remove("50"))
//
//	// Size
//	t.Log(bst.Size())
//
//	// Get(nil)
//	t.Log(bst.Get("50"))
//}
//
//func TestSimpleGraph(t *testing.T) {
//	g := graph.NewSimpleGraph()
//	v1, _ := g.AddVertex(1, 100)
//	v2, _ := g.AddVertex(2, 200)
//	v3, _ := g.AddVertex(3, 300)
//	g.AddEdge(v1, v2, 30)
//	g.AddEdge(v2, v1, 15)
//	g.AddEdge(v2, v3, 20)
//	g.AddEdge(v3, v2, 10)
//	g.AddEdge(v3, v1, 10)
//	g.AddEdge(v1, v3, 5)
//
//	t.Log("Vertex 1 neighbors size: ", len(g.GetVertex(1).Neighbors()))
//	t.Log("Vertex 1 neighbors: ", g.GetVertex(1).Neighbors()[0].Key(), g.GetVertex(1).Neighbors()[1].Key())
//	t.Log("Vertex 2 neighbors size: ", len(g.GetVertex(2).Neighbors()))
//	t.Log("Vertex 2 neighbors: ", g.GetVertex(2).Neighbors()[0].Key(), g.GetVertex(2).Neighbors()[1].Key())
//	t.Log("Vertex 3 neighbors size: ", len(g.GetVertex(3).Neighbors()))
//	t.Log("Vertex 3 neighbors: ", g.GetVertex(3).Neighbors()[0].Key(), g.GetVertex(3).Neighbors()[1].Key())
//	t.Log("Edge 1 to 2 value: ", g.GetEdge(v1, v2).Value())
//
//	g.RemoveEdge(v1, v2)
//	g.RemoveEdge(v2, v1)
//	t.Log("===After remove edge 1 to 2 & 2 to 1===")
//	t.Log("Vertex 1 neighbors size: ", len(g.GetVertex(1).Neighbors()))
//	t.Log("Vertex 1 neighbors: ", g.GetVertex(1).Neighbors()[0].Key())
//	t.Log("Vertex 2 neighbors size: ", len(g.GetVertex(2).Neighbors()))
//	t.Log("Vertex 2 neighbors: ", g.GetVertex(2).Neighbors()[0].Key())
//	t.Log("Vertex 3 neighbors size: ", len(g.GetVertex(3).Neighbors()))
//	t.Log("Vertex 3 neighbors: ", g.GetVertex(3).Neighbors()[0].Key(), g.GetVertex(3).Neighbors()[1].Key())
//	g.AddEdge(v1, v2, 30)
//	g.AddEdge(v2, v1, 15)
//	t.Log("===Edge 1 to 2 recovered===")
//
//	g.RemoveVertex(3)
//	t.Log("===After remove vertex 3===")
//	t.Log("Vertex 1 neighbors size: ", len(g.GetVertex(1).Neighbors()))
//	t.Log("Vertex 1 neighbors: ", g.GetVertex(1).Neighbors()[0].Key())
//	t.Log("Vertex 2 neighbors size: ", len(g.GetVertex(2).Neighbors()))
//	t.Log("Vertex 2 neighbors: ", g.GetVertex(2).Neighbors()[0].Key())
//}
//
//func TestIterateSimpleGraph(t *testing.T) {
//	g := graph.NewSimpleGraph()
//	v1, _ := g.AddVertex(1, 100)
//	v2, _ := g.AddVertex(2, 200)
//	v3, _ := g.AddVertex(3, 300)
//	v4, _ := g.AddVertex(4, 400)
//	g.AddEdge(v2, v1, 30)
//	g.AddEdge(v2, v3, 5)
//	g.AddEdge(v1, v4, 20)
//
//	// Iterate the graph by BFS order
//	t.Log("===IterateByBFS===")
//	iter := g.IterateByBFS(2)
//	var v graph.Vertex
//	route := ""
//	for iter.HasNext() {
//		v = iter.Next()
//		if route != "" {
//			route = route + "-->"
//		}
//		route = route + fmt.Sprintf("%d", v.Key())
//	}
//	t.Log("route:", route)
//
//	// Iterate the graph by DFS order
//	t.Log("===IterateByDFS===")
//	iter1 := g.IterateByDFS(2)
//	route = ""
//	for iter1.HasNext() {
//		v = iter1.Next()
//		if route != "" {
//			route = route + "-->"
//		}
//		route = route + fmt.Sprintf("%d", v.Key())
//	}
//	t.Log("route:", route)
//}
//
//func TestSimpleWeightedGraph(t *testing.T) {
//	wg := graph.NewSimpleWeightedGraph()
//	va, _ := wg.AddVertex("A", "A")
//	vb, _ := wg.AddVertex("B", "B")
//	vc, _ := wg.AddVertex("C", "C")
//	vd, _ := wg.AddVertex("D", "D")
//	ve, _ := wg.AddVertex("E", "E")
//
//	wg.AddEdgeWithWeight(va, vb, nil, 11)
//	wg.AddEdgeWithWeight(vb, va, nil, 11)
//	wg.AddEdgeWithWeight(va, vd, nil, 9)
//	wg.AddEdgeWithWeight(vd, va, nil, 9)
//	wg.AddEdgeWithWeight(va, ve, nil, 7)
//	wg.AddEdgeWithWeight(ve, va, nil, 7)
//	wg.AddEdgeWithWeight(vb, vc, nil, 10)
//	wg.AddEdgeWithWeight(vc, vb, nil, 10)
//	wg.AddEdgeWithWeight(vc, vd, nil, 12)
//	wg.AddEdgeWithWeight(vd, vc, nil, 12)
//	wg.AddEdgeWithWeight(vd, ve, nil, 8)
//	wg.AddEdgeWithWeight(ve, vd, nil, 8)
//
//	t.Log("===Initial Simple Weighted Graph===")
//	t.Log(wg)
//	t.Log("===Minial Spanning Tree===")
//	t.Log(algo.CreateMinimalSpanningTree(wg, ve))
//
//	t.Log("===Calculate Shortest Path===")
//	path, err := algo.Dijkstra(wg, va, vc)
//	t.Log(path, "total weight:", path.TotalWeight(), err)
//}
//
//func TestAdjacencyMatrixGraph(t *testing.T) {
//	wg := NewAdjacencyMatrixGraph()
//	va, _ := wg.AddVertex("A", "A")
//	vb, _ := wg.AddVertex("B", "B")
//	vc, _ := wg.AddVertex("C", "C")
//	vd, _ := wg.AddVertex("D", "D")
//	ve, _ := wg.AddVertex("E", "E")
//
//	wg.AddEdgeWithWeight(va, vb, nil, 11)
//	wg.AddEdgeWithWeight(vb, va, nil, 11)
//	wg.AddEdgeWithWeight(va, vc, nil, 10)
//	wg.AddEdgeWithWeight(vc, va, nil, 10)
//	wg.AddEdgeWithWeight(va, ve, nil, 20)
//	wg.AddEdgeWithWeight(ve, va, nil, 20)
//	wg.AddEdgeWithWeight(vb, vc, nil, 10)
//	wg.AddEdgeWithWeight(vc, vb, nil, 10)
//	wg.AddEdgeWithWeight(vc, vd, nil, 11)
//	wg.AddEdgeWithWeight(vd, vc, nil, 11)
//	wg.AddEdgeWithWeight(vd, ve, nil, 8)
//	wg.AddEdgeWithWeight(ve, vd, nil, 8)
//
//	t.Log("===Initial Adjacency Matrix Graph===")
//	t.Log(wg)
//
//	t.Log("===Calculate Shortest Path===")
//	path, err := algo.Dijkstra(wg, va, vc)
//	t.Log(path, "total weight:", path.TotalWeight(), err)
//}
//
//func TestMatrix(t *testing.T) {
//	mtrx := matrix.NewLinkedMatrix(5, 5, math.MaxFloat64)
//	t.Log(mtrx.Size())
//	mtrx.Set(0, 0, 1)
//	mtrx.Set(3, 2, 6)
//	t.Log(mtrx.Get(0, 0))
//	t.Log(mtrx.Get(3, 2))
//	mtrx.Resize(1, 1, math.MaxFloat64)
//	t.Log(mtrx.Size())
//	t.Log(mtrx.Get(0, 0))
//	t.Log(mtrx.Get(3, 2))
//}
