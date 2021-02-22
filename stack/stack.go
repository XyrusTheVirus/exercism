package stack

import (
	"fmt"
)

type Stack struct {
	top           *Node
	numOfElements int
}

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

func (s *Stack) Pop(head **Node) interface{} {

	if *head == nil {
		return nil
	}

	temp := *head
	*head = temp.next
	v := temp.val
	s.top = temp
	temp.next = nil
	temp = nil
	s.numOfElements--
	return v
}

func (s *Stack) Top() *Node {
	return s.top
}

func (s *Stack) NunOfElements() int {
	return s.numOfElements
}

func (s *Stack) Print(head *Node) {
	temp := head
	for temp != nil {
		fmt.Println(temp.val)
		temp = temp.next
	}
}
