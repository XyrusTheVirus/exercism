package stack

// Defines a node within the linked list
type Node struct {
	val  interface{} // The node's value
	next *Node       // Pointer to the next node
}
