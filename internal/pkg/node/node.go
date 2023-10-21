package node

type Node struct {
	value byte
	nodes *keep
}

type keep struct {
	arr   []*Node
	count int
}

func New() *Node {
	return &Node{}
}

func Make(capacity int) *Node {
	node := &Node{}

	if capacity > 0 {
		nodeKeep := &keep{}
		nodeKeep.arr = make([]*Node, capacity)
		node.nodes = nodeKeep
	}

	return node
}

func (n *Node) GetNodes() []*Node {
	return n.nodes.arr
}

func (n *Node) GetValue() byte {
	return n.value
}

func (n *Node) GetCount() int {
	if n.nodes != nil {
		return n.nodes.count
	}

	return 0
}

func (n *Node) SetNodes(nodes []*Node) {
	n.nodes.arr = nodes
}

func (n *Node) SetValue(value byte) {
	n.value = value
}

func (n *Node) GetNode(index int) *Node {
	return n.nodes.arr[index]
}

func (n *Node) SetNode(index int, node *Node) {
	n.nodes.arr[index] = node
}

func (n *Node) ResetCount() {
	n.nodes.count = cap(n.nodes.arr)
}

func (n *Node) Push(node *Node) {
	n.nodes.arr[n.nodes.count] = node

	n.nodes.count++
}

func (n *Node) Pop() *Node {
	n.nodes.count--

	return n.nodes.arr[n.nodes.count]
}
