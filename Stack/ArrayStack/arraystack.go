package arraystack

// Stack is a first in, last out data structure.
type Stack struct {
	array  []interface{}
	length int
}

// NewStack creates an empty stack instance.
func NewStack() *Stack {
	stack := Stack{}
	stack.array = make([]interface{}, 0)
	stack.length = 0
	return &stack
}

// Empty empties the Stack.
func (stack *Stack) Empty() {
	stack.array = make([]interface{}, 0)
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

// Peek returns the top node.
func (stack *Stack) Peek() interface{} {
	return stack.array[stack.length-1]
}

// Push adds a Node to the top of the Stack.
func (stack *Stack) Push(Data interface{}) {
	stack.array = append(stack.array, Data)
	stack.length++
}

// Pop removes the top Node in the stack and returns it.
func (stack *Stack) Pop() interface{} {
	temp := stack.array[stack.length-1]
	stack.array = stack.array[:stack.length-1]
	stack.length--
	return temp
}
