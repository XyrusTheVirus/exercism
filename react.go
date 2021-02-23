package react

import (
	"fmt"
	"react/graph"
	"reflect"

	guuid "github.com/google/uuid"
)

var g *graph.Graph = &graph.Graph{Vertices: make(map[string]*graph.Vertex)}

type React struct{}

type GCell struct {
	value int
	id    string
}

type ICell struct {
	GCell
}

type CCell struct {
	GCell
	callback func(int)
	compute1 func(int) int
	compute2 func(int, int) int
}

func (c *GCell) GetId() string {
	return c.id
}

func (r React) GenerateUniqueId() string {
	return guuid.New().String()
}

func (c *GCell) SetValue(value int) {
	c.value = value
	sortedGraph := g.GetGraphDependenies()
	for _, v := range sortedGraph {
		if reflect.TypeOf(v.GetCell()).String() == "*react.CCell" {
			switch n := len(v.GetAdjacents()); n {
			case 1:
				v.GetCell().(*GCell).value = v.GetCell().(CCell).compute1(v.GetAdjacents()[0].GetCell().(*GCell).Value())
			case 2:
				v.GetCell().(*GCell).value = v.GetCell().(CCell).compute2(v.GetAdjacents()[0].GetCell().(*GCell).Value(), v.GetAdjacents()[1].GetCell().(*GCell).Value())
			}
		}
	}
}

func (c CCell) AddCallback(callback func(value int)) Canceler {
	c.callback = callback
	return c
}

func (c CCell) Cancel() {
	c.callback = nil
}

func (r React) CreateInput(value int) InputCell {
	cell := &ICell{GCell: GCell{id: r.GenerateUniqueId(), value: value}}
	g.AddToGraph(g.CreateVertex(value), cell.GetId())
	return cell
}

func (r React) CreateCompute1(c1 Cell, compute func(value int) int) ComputeCell {
	cell := &CCell{GCell: GCell{id: r.GenerateUniqueId(), value: compute(c1.Value())}, callback: func(value int) { fmt.Println(value) }, compute1: compute}
	vertex := g.CreateVertex(cell.Value())
	g.AddToGraph(vertex, cell.GetId())
	g.Vertices[getIdByType(c1)].CreateEdge(vertex)
	return cell
}

func (r React) CreateCompute2(c1 Cell, c2 Cell, compute func(value1, value2 int) int) ComputeCell {
	cell := &CCell{GCell: GCell{value: compute(c1.Value(), c2.Value())}, callback: func(value int) { fmt.Println(value) }, compute2: compute}
	vertex := g.CreateVertex(cell.Value())
	g.AddToGraph(vertex, cell.GetId())
	g.Vertices[getIdByType(c1)].CreateEdge(vertex)
	g.Vertices[getIdByType(c2)].CreateEdge(vertex)
	return cell
}

func getIdByType(c Cell) string {
	id := ""
	fmt.Println(reflect.TypeOf(c))
	switch c.(type) {
	case *CCell:
		id = c.(*CCell).GetId()
	case *ICell:
		id = c.(*ICell).GetId()
	}

	return id
}

func (c *GCell) Value() int {
	return c.value
}

func New() Reactor {
	return React{}
}
