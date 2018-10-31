package linkedqueue

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

// Queue is a first in first out data structure.
type Queue struct {
	front  *Node
	length int
}

// NewQueue creates an empty queue instance.
func NewQueue() *Queue {
	queue := Queue{}
	queue.front = nil
	queue.length = 0
	return &queue
}

// Empty empties the Queue.
func (queue *Queue) Empty() {
	queue.front = nil
	queue.length = 0
}

// IsEmpty returns true if the Queue is empty and false if not.
func (queue *Queue) IsEmpty() bool {
	if queue.length == 0 {
		return true
	}
	return false
}

// Len returns the length of the Queue.
func (queue *Queue) Len() int {
	return queue.length
}

// Peek returns front node.
func (queue *Queue) Peek() *Node {
	return queue.front
}

// Enqueue adds a Node to the end of the Queue.
func (queue *Queue) Enqueue(data interface{}) {
	if queue.length == 0 {
		queue.front = NewNode(data, nil)
	} else {
		temp := queue.front
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = NewNode(data, nil)
	}
	queue.length++
}

// Dequeue removes the first Node in the Queue and returns it.
func (queue *Queue) Dequeue() *Node {
	temp := queue.front
	queue.front = queue.front.next
	queue.length--
	return temp
}
