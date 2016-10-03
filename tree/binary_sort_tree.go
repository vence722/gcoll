// gcoll
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)

package tree

import (
	"github.com/vence722/gcoll/list"
	"github.com/vence722/gcoll/maps"
)

// The BinarySortTree struct
type BinarySortTree struct {
	BaseTree
}

// Return a new BinarySortTree
func NewBinarySortTree() *BinarySortTree {
	return &BinarySortTree{}
}

// StringMap interface
// Return the size of the map
func (this *BinarySortTree) Size() int {
	return this.size
}

// Return the map containing elements or not
func (this *BinarySortTree) IsEmpty() bool {
	return this.size == 0
}

// Return the map conaining specified key or not
func (this *BinarySortTree) ContainsKey(key string) bool {
	return this.Get(key) != nil
}

// Return the map conaining specified value or not
func (this *BinarySortTree) ContainsValue(value interface{}) bool {
	stack := list.NewLinkedList()
	if this.root != nil {
		stack.Push(this.root)
		for stack.Size() > 0 {
			pn := stack.Pop().(*TreeNode)
			if goNode(stack, pn, value) {
				return true
			}
		}
	}
	return false
}

func goNode(stack list.Stack, node *TreeNode, value interface{}) bool {
	if node != nil {
		if node.value == value {
			return true
		}
		if node.right != nil {
			stack.Push(node.right)
		}
		if node.left != nil {
			stack.Push(node.left)
		}
	}
	return false
}

// Add new key value pair to this map
func (this *BinarySortTree) Put(key string, value interface{}) bool {
	if this.root != nil {
		p := this.root
		for {
			if key < p.key {
				if p.left != nil {
					p = p.left
				} else {
					node := &TreeNode{key: key, value: value, parent: p}
					p.left = node
					this.size++
					break
				}
			} else if key > p.key {
				if p.right != nil {
					p = p.right
				} else {
					node := &TreeNode{key: key, value: value, parent: p}
					p.right = node
					this.size++
					break
				}
			} else {
				p.value = value
				break
			}
		}
	} else {
		node := &TreeNode{key: key, value: value}
		this.root = node
		this.size++
	}
	return true
}

// Return the element with the specified key in this map
func (this *BinarySortTree) Get(key string) interface{} {
	if this.root == nil {
		return nil
	}
	pn := this.root
	for pn != nil {
		if key < pn.key {
			pn = pn.left
		} else if key > pn.key {
			pn = pn.right
		} else {
			return pn.value
		}
	}
	return nil
}

// Remove value with specified key from this map
func (this *BinarySortTree) Remove(key string) bool {
	if this.root == nil {
		return false
	}
	pn := this.root
	for pn != nil {
		if key < pn.key {
			pn = pn.left
		} else if key > pn.key {
			pn = pn.right
		} else {
			// find the entry
			if !(pn.left != nil && pn.right != nil) {
				// pn has no more than one child
				var ps *TreeNode
				if pn.left != nil {
					ps = pn.left
				} else if pn.right != nil {
					ps = pn.right
				} else {
					ps = nil
				}
				if pn.parent.left == pn {
					pn.parent.left = ps
				} else if pn.parent.right == pn {
					pn.parent.right = ps
				}
				this.size--
				return true
			} else {
				// pn has both left and right children
				// find pn's successor, remove it and copy it info pn
				// tips:psu have NO left child!
				psu := pn.Successor()
				pn.key = psu.key
				pn.value = psu.value
				if psu.parent != pn {
					psu.parent.left = psu.right
				} else {
					psu.parent.right = psu.right
				}
				this.size--
				return true
			}
		}
	}
	return false
}

// Add another map to this map
func (this *BinarySortTree) PutAll(amap maps.StringMap) bool {
	if amap == nil {
		return false
	}
	entries := amap.Entries()
	for _, entry := range entries {
		this.Put(entry.Key, entry.Value)
	}
	return true
}

// Remove all elements from this map
func (this *BinarySortTree) Clear() {
	this.root = nil
	this.size = 0
}

// Views
// Return a list containing all the keys in the map
func (this *BinarySortTree) Keys() []string {
	keys := list.NewArrayList()
	stack := list.NewLinkedList()
	if this.root != nil {
		stack.Push(this.root)
		for stack.Size() > 0 {
			pn := stack.Pop().(*TreeNode)
			goNodeKeys(stack, pn, keys)
		}
	}
	skeys := make([]string, INIT_LEN, INIT_CAP)
	for i := 0; i < keys.Size(); i++ {
		skeys = append(skeys, keys.Get(i).(string))
	}
	return skeys
}

func goNodeKeys(stack list.Stack, node *TreeNode, keys list.List) {
	if node != nil {
		keys.Add(node.key)
		if node.right != nil {
			stack.Push(node.right)
		}
		if node.left != nil {
			stack.Push(node.left)
		}
	}
}

// Return a list containing all the values in the map
func (this *BinarySortTree) Values() []interface{} {
	values := list.NewArrayList()
	stack := list.NewLinkedList()
	if this.root != nil {
		stack.Push(this.root)
		for stack.Size() > 0 {
			pn := stack.Pop().(*TreeNode)
			goNodeValues(stack, pn, values)
		}
	}
	svalues := make([]interface{}, INIT_LEN, INIT_CAP)
	for i := 0; i < values.Size(); i++ {
		svalues = append(svalues, values.Get(i))
	}
	return svalues
}

func goNodeValues(stack list.Stack, node *TreeNode, values list.List) {
	if node != nil {
		values.Add(node.value)
		if node.right != nil {
			stack.Push(node.right)
		}
		if node.left != nil {
			stack.Push(node.left)
		}
	}
}

// Return a list containing all the entries in the map
func (this *BinarySortTree) Entries() []maps.StringMapEntry {
	entries := list.NewArrayList()
	stack := list.NewLinkedList()
	if this.root != nil {
		stack.Push(this.root)
		for stack.Size() > 0 {
			pn := stack.Pop().(*TreeNode)
			goNodeEntries(stack, pn, entries)
		}
	}
	smes := make([]maps.StringMapEntry, INIT_LEN, INIT_CAP)
	for i := 0; i < entries.Size(); i++ {
		smes = append(smes, entries.Get(i).(maps.StringMapEntry))
	}
	return smes
}

func goNodeEntries(stack list.Stack, node *TreeNode, entries list.List) {
	if node != nil {
		entries.Add(maps.StringMapEntry{node.key, node.value})
		if node.right != nil {
			stack.Push(node.right)
		}
		if node.left != nil {
			stack.Push(node.left)
		}
	}
}

//// Return the string that describes the contains of this map
//func (this *BinarySortTree) String() string {

//}

// Tree interface
// Return the root of the tree
func (this *BinarySortTree) Root() *TreeNode {
	return this.root
}

// Return the pointer to the TreeNode with the specified key
func (this *BinarySortTree) Node(key string) *TreeNode {
	if this.root == nil {
		return nil
	}
	pn := this.root
	for pn != nil {
		if key < pn.key {
			pn = pn.left
		} else if key > pn.key {
			pn = pn.right
		} else {
			return pn
		}
	}
	return nil
}

// Return the pointer to the key-minimum TreeNode of the subtree that the root of which is this
func (this *TreeNode) MinimumAsSubTreeRoot() *TreeNode {
	p := this
	for p.left != nil {
		p = p.left
	}
	return p
}

// Return the pointer to the successor TreeNode of this node
func (this *TreeNode) Successor() *TreeNode {
	if this.right != nil {
		return this.right.MinimumAsSubTreeRoot()
	}
	p := this.parent
	s := this
	for p != nil && p.right == s {
		s = p
		p = p.parent
	}
	return p
}
