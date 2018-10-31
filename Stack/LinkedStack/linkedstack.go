package linkedstack

// Node is a singly linked list data structure used to create the queue.
type Node struct {
	data interface{}
	next *Node
}

// NewNode creates a node instance with a given data and next variable.
func NewNode(data interface{}, next *Node) *Node {
	node := Node{}
	node.data = data
	node.next = next
	return &node
}

// Data returns the data of a node.
func (node *Node) Data() interface{} {
	return node.data
}

// SetData sets the data for a node.
func (node *Node) SetData(data interface{}) {
	node.data = data
}

// Next returns the next node of a node.
func (node *Node) Next() *Node {
	return node.next
}

// SetNext sets the next node of a node.
func (node *Node) SetNext(next *Node) {
	node.next = next
}

// Stack is a first in, last out data structure.
type Stack struct {
	top    *Node
	length int
}

// NewStack creates an empty stack instance.
func NewStack() *Stack {
	stack := Stack{}
	stack.top = nil
	stack.length = 0
	return &stack
}

// Empty empties the Stack.
func (stack *Stack) Empty() {
	stack.top = nil
	stack.length = 0
}

// IsEmpty returns true if the Stack is empty and false if not.
func (stack *Stack) IsEmpty() bool {
	if stack.length == 0 {
		return true
	}
	return false
}

// Len returns the length of the Stack.
func (stack *Stack) Len() int {
	return stack.length
}

// Peek returns the top Node.
func (stack *Stack) Peek() interface{} {
	return stack.top
}

// Push adds a Node to the top of the Stack.
func (stack *Stack) Push(data interface{}) {
	stack.top = NewNode(data, stack.top)
	stack.length++
}

// Pop removes the top Node in the stack and returns it.
func (stack *Stack) Pop() interface{} {
	temp := stack.top
	stack.top = stack.top.next
	stack.length--
	return temp
}
