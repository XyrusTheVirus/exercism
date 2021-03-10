package graph

import (
	"react/stack"
)

// Defines the graph entity
type Graph struct {
	Vertices map[string]*Vertex // Map to store all graph vertices
}

// Defines Vertex entity
type Vertex struct {
	cell      interface{} // Stores cell entity
	adjacents []*Vertex   // Defines edge between vertices u and v
	visited   bool        // Indicates whether vertex was visited through the DFS traversal
}

// Returns the vertex's cell entity
func (v *Vertex) GetCell() interface{} {
	return v.cell
}

// Returns an vertex's neighbor
func (v *Vertex) GetAdjacents() []*Vertex {
	return v.adjacents
}

// Creates brand new vertex
// Receives cell as the vertex's cell entity
// Returns *Vertex instance
func (g *Graph) CreateVertex(cell interface{}) *Vertex {
	return &Vertex{cell: cell, visited: false}
}

// Inserts vertex to the graph data structure
// Receives *Vertex v to insert the graph
// Receives string id to indicate entry in graph
func (g *Graph) AddToGraph(v *Vertex, id string) {
	g.Vertices[id] = v
}

// Creates edge between vertices v and u, which indicates v is connected to u (v -> u)
// Receives u *Vertex to connect with vertex v
func (v *Vertex) CreateEdge(u *Vertex) {
	v.adjacents = append(v.adjacents, u)
}

// Returns all graph's vertices dependencies as and sorted array
// Returns sorted array with the proper dependencies order
func (g *Graph) GetGraphDependenies() []*Vertex {
	return topologicalSort(g, &stack.Stack{})
}

// Returns the vertices dependencies by using the DFS traversal
// Receives the graph to apply the sort on
// Receives the stack, sent to the DFS traversal
// Returns the reversed order of the DFS output
func topologicalSort(g *Graph, s *stack.Stack) []*Vertex {
	return reverse(DFS(g, s))
}

// Reversing an array
// Receives []*Vertex v to be reversed
// Returns the reversed array
func reverse(v []*Vertex) []*Vertex {
	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}

	return v
}

//
func DFS(g *Graph, s *stack.Stack) []*Vertex {
	var sortedGraph []*Vertex

	for _, v := range g.Vertices {
		if v.visited == false {
			visit(v, s, &sortedGraph)
		}
	}

	g.restoreVisitedVerteces()
	return sortedGraph
}

func visit(v *Vertex, s *stack.Stack, sortedGraph *[]*Vertex) {
	v.visited = true
	top := s.Top()
	s.Push(&top, v)

	for _, adjacent := range v.adjacents {
		if adjacent.visited == false {
			visit(adjacent, s, sortedGraph)
		}
	}

	top = s.Top()
	*sortedGraph = append(*sortedGraph, (s.Pop(&top)).(*Vertex))
	return
}

func (g *Graph) restoreVisitedVerteces() {
	for _, v := range g.Vertices {
		v.visited = false
	}
}
