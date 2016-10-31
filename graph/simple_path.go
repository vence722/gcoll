package graph

type SimplePath struct {
	nodes []Node
	src   Node
	dest  Node
}

func NewSimplePath() *SimplePath {
	return &SimplePath{}
}

func (this *SimplePath) Src() Node {
	return this.src
}

func (this *SimplePath) Dest() Node {
	return this.dest
}

func (this *SimplePath) Nodes() []Node {
	return this.nodes
}

func (this *SimplePath) AddNode(key interface{}, value interface{}, weightToPrev float64) error {
	if this.GetNode(key) != nil {
		return ERR_NODE_EXISTS
	}
	node := &SimpleNode{key: key, value: value, weightToPrev: weightToPrev}
	dest := this.dest.(*SimpleNode)
	dest.next = node
	this.nodes = append(this.nodes, node)
	return nil
}

func (this *SimplePath) GetNode(key interface{}) Node {
	for _, node := range this.Nodes() {
		if node.Key() == key {
			return node
		}
	}
	return nil
}

func (this *SimplePath) TotalWeight() float64 {
	var totalW float64 = 0
	curr := this.src
	for curr != nil {
		totalW += curr.WeightToPrev()
		curr = curr.Next()
	}
	return totalW
}

type SimpleNode struct {
	key          interface{}
	value        interface{}
	next         *SimpleNode
	weightToPrev float64
}

func (this *SimpleNode) Key() interface{} {
	return this.key
}

func (this *SimpleNode) Value() interface{} {
	return this.value
}

func (this *SimpleNode) Next() Node {
	return this.next
}

func (this *SimpleNode) WeightToPrev() float64 {
	return this.weightToPrev
}
