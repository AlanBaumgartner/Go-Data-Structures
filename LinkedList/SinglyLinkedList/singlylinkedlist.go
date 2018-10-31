package singlylinkedlist

// Node is the singly linked list data structure.
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

// new list, add to list, clear list, is empty

// List is the doubly linked list data structure.
type List struct {
	head   *Node
	length int
}

// new list, add to list, clear list, is empty

// NewList creates an empty list instance.
func NewList() *List {
	list := List{}
	list.head = nil
	list.length = 0
	return &list
}

// Head returns the head node of the list.
func (list *List) Head() *Node {
	return list.head
}

// SetHead sets the head node of the list.
func (list *List) SetHead(head *Node) {
	list.head = head
}

// Len returns the length of the list.
func (list *List) Len() int {
	return list.length
}

// Clear clears all data in the list.
func (list *List) Clear() {
	list.head = nil
	list.length = 0
}

// Add inserts the node to the end of the list.
func (list *List) Add(data interface{}) {
	if list.length == 0 {
		list.head = NewNode(data, nil)
	} else {
		temp := list.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = NewNode(data, nil)
	}
	list.length++
}

// Remove removes element with given data.
func (list *List) Remove(data interface{}) {
	if list.head != nil {
		temp := list.head
		for temp.next != nil && temp.next.data != data {
			temp = temp.next
		}
		if temp.next.data == data {
			list.length--
			temp.next = temp.next.next
		}
	}
}
