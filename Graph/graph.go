package graph

type Node struct {
	data  interface{}
	edges []*Node
}

func NewNode(data interface{}) *Node {
	node := Node{}
	node.data = data
	node.edges = make([]*Node, 0)
	return &node
}

func (node *Node) AddEdge(edge *Node) {
	node.edges = append(node.edges, edge)
}

func (node *Node) RemoveEdge(edge *Node) {
	for index, cur := range node.edges {
		if cur == edge {
			node.edges = append(node.edges[:index], node.edges[index:]...)
			return
		}
	}
}

type Graph struct {
	nodes []*Node
}

func NewGraph() *Graph {
	graph := Graph{}
	graph.nodes = make([]*Node, 0)
	return &graph
}

func (graph *Graph) AddNode(data interface{}) {
	graph.nodes = append(graph.nodes, NewNode(data))
}

func (graph *Graph) AddEdge(node1 *Node, node2 *Node) {
	node1.AddEdge(node2)
}

func (graph *Graph) RemoveNode(data interface{}) {
	// graph.nodes = append(graph.nodes, NewNode(data))
}

func (graph *Graph) RemoveEdge(node1 *Node, node2 *Node) {
	node1.RemoveEdge(node2)
}
