package tree

import "strconv"

// Node the representation of each node on a tree, including the root.
type Node struct {
	key   int
	value interface{}
	left  *Node
	right *Node
}

// NewNode creates a new tree node.
func NewNode(key int, value interface{}) *Node {
	node := Node{}
	node.key = key
	node.value = value
	node.left = nil
	node.right = nil
	return &node
}

// Add a node into the binary tree.
func (node *Node) Add(new *Node) {
	if new.key < node.key {
		if node.left == nil {
			node.left = new
		} else {
			node.left.Add(new)
		}
	} else {
		if node.right == nil {
			node.right = new
		} else {
			node.right.Add(new)
		}
	}
}

// Remove a node from the binary tree.
func (node *Node) Remove(key int) *Node {
	if node == nil {
		return nil
	} else if key < node.key {
		node.left = node.left.Remove(key)
		return node
	} else if key > node.key {
		node.right = node.right.Remove(key)
		return node
	} else if node.left == nil && node.right == nil {
		node = nil
		return nil
	} else if node.left == nil {
		node = node.right
		return node
	} else if node.right == nil {
		node = node.left
		return node
	}

	newParent := node.right
	for newParent != nil && newParent.left != nil {
		newParent = newParent.left
	}
	node.key, node.value = newParent.key, newParent.value
	node.right = node.right.Remove(node.key)
	return node
}

// Get the value from a given key in the binary tree.
func (node *Node) Get(key int) interface{} {
	if key == node.key {
		return node.value
	} else if key < node.key {
		if node.left == nil {
			return nil
		} else {
			return node.left.Get(key)
		}
	} else {
		if node.right == nil {
			return nil
		} else {
			return node.right.Get(key)
		}
	}
}

// Temp View Tree
func (node *Node) String() []string {
	a := make([]string, 0)
	a = append(a, strconv.Itoa(node.key))
	if node.left != nil {
		a = append(a, node.left.String()...)
	}
	if node.right != nil {
		a = append(a, node.right.String()...)
	}
	return a
}
