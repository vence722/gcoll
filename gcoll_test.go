// gcoll
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)
package gcoll

import (
	"testing"

	"github.com/vence722/gcoll/lru"
	"github.com/vence722/gcoll/tree"
)

func TestLists(t *testing.T) {

}

func TestLRUCache(t *testing.T) {
	// === Test Fifo LRU Cache ===
	fifoLRU := lru.NewFifoLRUCache(5)
	fifoLRU.Put(1, 'a')
	fifoLRU.Put('a', 8)
	fifoLRU.Put("3", 0x5F)
	fifoLRU.Put("aa", "cc")
	fifoLRU.Put("bb", 0x5F)

	t.Log(fifoLRU.Size())
	t.Log(fifoLRU.Get(1))

	// Test insert more tan Cap
	fifoLRU.Put("cc", 5214)

	// Out of FIFO
	t.Log(fifoLRU.Get(1))

	fifoLRU.Clear()
	t.Log(fifoLRU.Size())

	// === Test HitsMap LRU Cache ===
	hitsMapLRU := lru.NewHitsMapLRUCache(5)
	hitsMapLRU.Put(1, 'a')
	hitsMapLRU.Put('a', 8)
	hitsMapLRU.Put("3", 0x5F)
	hitsMapLRU.Put("aa", "cc")
	hitsMapLRU.Put("bb", 0x5F)

	t.Log(hitsMapLRU.Size())
	t.Log(hitsMapLRU.Get(1))

	// Get will add hits count
	t.Log(hitsMapLRU.Get('a'))
	t.Log(hitsMapLRU.Get("3"))
	t.Log(hitsMapLRU.Get("bb"))

	// Test insert more tan Cap
	hitsMapLRU.Put("cc", 5214)
	// Item will least hit count removed
	t.Log(fifoLRU.Get("aa"))

	fifoLRU.Clear()
	t.Log(fifoLRU.Size())
}

func TestTrees(t *testing.T) {
	// BST
	bst := tree.NewBinarySortTree()
	bst.Put("100", 100)
	bst.Put("50", 50)
	bst.Put("220", 220)
	bst.Put("99", 99)

	// Get
	t.Log(bst.Get("50"))

	// Keys
	t.Log(bst.Keys())

	// Values
	t.Log(bst.Values())

	// Root
	t.Log(bst.Root())

	// Remove
	t.Log(bst.Remove("50"))

	// Size
	t.Log(bst.Size())

	// Get(nil)
	t.Log(bst.Get("50"))
}
