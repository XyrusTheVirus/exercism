package graph

import (
	"fmt"
	"react/stack"
)

type Graph struct {
	Vertices map[string]*Vertex
}

type Vertex struct {
	cell      interface{}
	adjacents []*Vertex
	visited   bool
}

func (v *Vertex) GetCell() interface{} {
	return v.cell
}

func (v *Vertex) GetAdjacents() []*Vertex {
	return v.adjacents
}

func (g *Graph) CreateVertex(cell interface{}) *Vertex {
	return &Vertex{cell: cell, visited: false}
}

func (g *Graph) AddToGraph(v *Vertex, id string) {
	g.Vertices[id] = v
}

func (v *Vertex) CreateEdge(u *Vertex) {
	v.adjacents = append(v.adjacents, u)
}

func (g *Graph) GetGraphDependenies() []*Vertex {
	return topologicalSort(g, &stack.Stack{})
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
	fmt.Println(g.Vertices)
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
	fmt.Println("Current Strack: ")
	s.Push(&top, v)
	s.Print(top)
	for _, adjacent := range v.adjacents {
		if adjacent.visited == false {
			visit(adjacent, s, sortedGraph)
		}
	}

	top = s.Top()
	*sortedGraph = append(*sortedGraph, (s.Pop(&top)).(*Vertex))
	return
}
