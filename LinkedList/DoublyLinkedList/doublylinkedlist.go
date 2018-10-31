package singlylinkedlist

// Node is an item in the doubly linked list data structure.
type Node struct {
	data interface{}
	prev *Node
	next *Node
}

// NewNode creates a node instance with a given data and next variable.
func NewNode(data interface{}, prev *Node, next *Node) *Node {
	node := Node{}
	node.data = data
	node.prev = prev
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

// Previous returns the previous node of a node.
func (node *Node) Previous() *Node {
	return node.prev
}

// SetPrevious sets the previous node of a node.
func (node *Node) SetPrevious(prev *Node) {
	node.next = prev
}

// Next returns the next node of a node.
func (node *Node) Next() *Node {
	return node.next
}

// SetNext sets the next node of a node.
func (node *Node) SetNext(next *Node) {
	node.next = next
}

// List is the doubly linked list data structure.
type List struct {
	head   *Node
	tail   *Node
	length int
}

// new list, add to list, clear list, is empty

// NewList creates an empty list instance.
func NewList() *List {
	list := List{}
	list.head = nil
	list.tail = nil
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

// Tail returns the head node of the list.
func (list *List) Tail() *Node {
	return list.tail
}

// SetTail sets the head node of the list.
func (list *List) SetTail(tail *Node) {
	list.tail = tail
}

// Len returns the length of the list.
func (list *List) Len() int {
	return list.length
}

// Clear clears all data in the list.
func (list *List) Clear() {
	list.head = nil
	list.tail = nil
	list.length = 0
}

// Add inserts the node to the end of the list.
func (list *List) Add(data interface{}) {
	if list.length == 0 {
		list.head = NewNode(data, nil, nil)
	} else {
		list.tail.SetNext(NewNode(data, list.tail, nil))
		list.tail = list.tail.Next()
	}
	list.length++
}

// Remove removes element with given data.
func (list *List) Remove(data interface{}) {
	if list.head != nil {
		temp := list.head
		for temp.next != nil && temp.data != data {
			temp = temp.next
		}
		if temp.data == data {
			list.length--
			temp.prev.next = temp.next
			temp.next.prev = temp.prev
		}
	}
}
