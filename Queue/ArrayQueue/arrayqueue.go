package arrayqueue

// Queue is a first in first out data structure.
type Queue struct {
	array  []interface{}
	length int
}

// NewQueue creates an empty queue instance.
func NewQueue() *Queue {
	queue := Queue{}
	queue.array = make([]interface{}, 0)
	queue.length = 0
	return &queue
}

// Empty empties the Queue.
func (queue *Queue) Empty() {
	queue.array = make([]interface{}, 0)
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
func (queue *Queue) Peek() interface{} {
	return queue.array[0]
}

// Enqueue adds a Node to the end of the Queue.
func (queue *Queue) Enqueue(Data interface{}) {
	queue.array = append(queue.array, Data)
	queue.length++
}

// Dequeue removes the first Node in the Queue and returns it.
func (queue *Queue) Dequeue() interface{} {
	temp := queue.array[0]
	queue.array = queue.array[1:]
	queue.length--
	return temp
}
