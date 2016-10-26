package graph

type SimplePath struct {
	nodes []Node
	start Node
	dest  Node
}

func (this *SimplePath) Start() Node {
	return this.start
}

func (this *SimplePath) Dest() Node {
	return this.dest
}

func (this *SimplePath) Nodes() []Node {
	return this.nodes
}

func (this *SimplePath) AddNode(key interface{}, value interface{}) error {
	if this.GetNode(key) != nil {
		return ERR_NODE_EXISTS
	}
	node := &SimpleNode{key: key, value: value}
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

type SimpleNode struct {
	key   interface{}
	value interface{}
	next  *SimpleNode
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
