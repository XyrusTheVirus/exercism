package graph

import (
	"react/stack"
)

type Graph struct {
	Vertices []*Vertex
}

type Vertex struct {
	cell      int //*react.Cell
	adjacents []*Vertex
	visited   bool
}

func (g *Graph) CreateVertex(cell int) *Vertex {
	return &Vertex{cell: cell, visited: false}
}

func (g *Graph) AddToGraph(v *Vertex) {
	g.Vertices = append(g.Vertices, v)
}

func (v *Vertex) CreateEdge(u *Vertex) {
	v.adjacents = append(v.adjacents, u)
}

func (g *Graph) UpdateGraphDependenies() {
	topologicalSort(g, &stack.Stack{})
}

func topologicalSort(g *Graph, s *stack.Stack) []*Vertex {
	return reverse(DFS(g, s))
}

func reverse(v []*Vertex) []*Vertex {
	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}

	return v
}

func DFS(g *Graph, s *stack.Stack) []*Vertex {
	var sortedGraph []*Vertex
	for _, v := range g.Vertices {
		if v.visited == false {
			visit(v, s, &sortedGraph)
		}
	}

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
	*sortedGraph = append(*sortedGraph, s.Pop(&top))
	return
}
