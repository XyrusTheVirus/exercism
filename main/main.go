package main

import (
	"react/graph"
	"strconv"
)

func main() {
	g := graph.Graph{Vertices: make(map[string]*graph.Vertex)}

	for i := 1; i <= 5; i++ {
		g.AddToGraph(g.CreateVertex(i), strconv.Itoa(i-1))
	}

	g.Vertices["0"].CreateEdge(g.Vertices["1"])
	g.Vertices["0"].CreateEdge(g.Vertices["2"])
	g.Vertices["1"].CreateEdge(g.Vertices["3"])
	g.Vertices["2"].CreateEdge(g.Vertices["3"])
	g.Vertices["1"].CreateEdge(g.Vertices["4"])
	g.Vertices["3"].CreateEdge(g.Vertices["4"])
	g.GetGraphDependenies()
}
