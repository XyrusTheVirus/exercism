package main

import (
	"react/graph"
)

func main() {
	g := &graph.Graph{}

	for i := 1; i <= 5; i++ {
		g.AddToGraph(g.CreateVertex(i))
	}

	g.Vertices[0].CreateEdge(g.Vertices[1])
	g.Vertices[0].CreateEdge(g.Vertices[2])
	g.Vertices[1].CreateEdge(g.Vertices[3])
	g.Vertices[2].CreateEdge(g.Vertices[3])
	g.Vertices[1].CreateEdge(g.Vertices[4])
	g.Vertices[3].CreateEdge(g.Vertices[4])
	g.UpdateGraphDependenies()
}
