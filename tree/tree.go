// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package tree

const (
	INIT_LEN = 0
	INIT_CAP = 16
)

// The BaseTree struct
type BaseTree struct {
	root *TreeNode
	size int
	//layers int
}

// The TreeNode struct
type TreeNode struct {
	parent *TreeNode
	left   *TreeNode
	right  *TreeNode
	key    string
	value  any
}
