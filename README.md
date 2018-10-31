# Overview
Data structures and abstract data types coded in Google's Goland.

## Abstract Data Types
### Singly Linked List
#### What is a Singly Linked List?
A singly linked list is a linear data type similar to an array with a different philosophy on length mutability. An array normally has a static size whereas a linked list can be any size, all it has to do is store a reference to the next node in the current last node. A singly linked list only stores a reference to the first node and the next node.

### Doubly Linked List
#### What is a Doubly Linked List?
A doubly linked list is almost identical to a singly linked list with the addition of a tail node and a previous node reference. This takes up more memory per node but allows you to easily backtrack when iterating through the nodes.

## Data Structures
### Queue
#### What is a Queue?
A queue is a first in first out data structure. An easy way to think of a queue like standing in a line (or queue for the British) at a grocery store, the first person in the line is the first person to get helped. The methods for adding and removing from a queue are simply Enqueue() and Dequeue(). You can also view what item is going to be "Dequeued" next by using Peek(). In the source I have provided both an array based and linked list based implementation of a queue.

### Stack
#### What is a Stack?
A stack is a first in last out data structure. An easy way to think of a stack is like doing the dishes. When you have a stack of plates, the last plate to be put on the stack is the first plate to get cleaned. The methods for adding and removing from a queue are Push() and Pop(). You can also view what item is going to be "Popped" next by using Peek(). In the source I have provided both an array based and linked list based implementation of a stack.