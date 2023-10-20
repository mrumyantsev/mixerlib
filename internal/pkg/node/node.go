package node

type Node struct {
	value byte
	keep  *NodesKeep
}

type NodesKeep struct {
	nodes  []*Node
	length int
}

func New() *Node {
	return &Node{}
}

func Make(capacity int) *Node {
	node := &Node{}

	if capacity > 0 {
		nodeKeep := &NodesKeep{}
		nodeKeep.nodes = make([]*Node, capacity)
		node.keep = nodeKeep
	}

	return node
}

func (n *Node) GetNodes() []*Node {
	return n.keep.nodes
}

func (n *Node) GetValue() byte {
	return n.value
}

func (n *Node) GetLength() int {
	if n.keep != nil {
		return n.keep.length
	}

	return 0
}

func (n *Node) SetNodes(nodes []*Node) {
	n.keep.nodes = nodes
}

func (n *Node) SetValue(value byte) {
	n.value = value
}

func (n *Node) GetNode(index int) *Node {
	return n.keep.nodes[index]
}

func (n *Node) SetNode(index int, node *Node) {
	n.keep.nodes[index] = node
}

func (n *Node) Push(node *Node) {
	n.keep.nodes[n.keep.length] = node

	n.keep.length++
}

func (n *Node) Pop() *Node {
	if n.keep.length == 0 {
		return nil
	}

	n.keep.length--

	return n.keep.nodes[n.keep.length]
}
