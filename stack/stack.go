package stack

import (
	"fmt"
)

// Defines a stack data structure implemented by linked list
type Stack struct {
	top           *Node // Denotes the top of the stack
	numOfElements int   // denotes the number of elements in the stack
}

// Inserts item to the top of the stack
// Receives the head of the linked list
// Receives interface{} item to insert
func (s *Stack) Push(head **Node, v interface{}) {
	node := &Node{
		val:  v,
		next: nil,
	}

	if *head == nil {
		*head = node
		s.top = *head
	} else {
		node.next = *head
		*head = node
		s.top = *head
	}

	s.numOfElements++
	return
}

// Removes item from the top of the stack
// Receives the head of the linked list
func (s *Stack) Pop(head **Node) interface{} {

	if *head == nil {
		return nil
	}

	temp := *head
	*head = temp.next
	v := temp.val
	s.top = *head
	temp = nil
	s.numOfElements--
	return v
}

// Returns the top of the stack
func (s *Stack) Top() *Node {
	return s.top
}

// Returns the number of the elements
func (s *Stack) NunOfElements() int {
	return s.numOfElements
}

// Returns if stack is empty or not
func (s *Stack) isEmpty() bool {
	return s.top == nil
}

// Prints the stack
// Receives the head of the linked list
func (s *Stack) Print(head *Node) {
	temp := head
	for temp != nil {
		fmt.Println(temp.val)
		temp = temp.next
	}
}
