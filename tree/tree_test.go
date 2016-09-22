// gcoll v1.0
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)

package tree

import (
	"testing"
)

func TestTrees(t *testing.T) {
	// BST
	bst := NewBinarySortTree()
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
