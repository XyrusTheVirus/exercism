package stack

import (
	"react/graph"
)

type Node struct {
	vertex *graph.Vertex
	next   *Node
}
