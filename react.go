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
	numOfIndepandants int
	inputs            []Cell
	callback          func(int)
	compute1          func(int) int
	compute2          func(int, int) int
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
			switch len(v.GetCell().(*CCell).inputs) {
			case 1:
				v.GetCell().(*CCell).value = v.GetCell().(*CCell).compute1(getValueByType(v.GetCell().(*CCell).inputs[0]))
			case 2:
				v.GetCell().(*CCell).value = v.GetCell().(*CCell).compute2(getValueByType(v.GetCell().(*CCell).inputs[0]), getValueByType(v.GetCell().(*CCell).inputs[1]))
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
	g.AddToGraph(g.CreateVertex(cell), cell.GetId())
	return cell
}

func (r React) CreateCompute1(c1 Cell, compute func(value int) int) ComputeCell {
	cell := &CCell{GCell: GCell{id: r.GenerateUniqueId(), value: compute(c1.Value())}, callback: func(value int) { fmt.Println(value) }, compute1: compute, numOfIndepandants: 1}
	cell.inputs = append(cell.inputs, c1)
	vertex := g.CreateVertex(cell)
	g.AddToGraph(vertex, cell.GetId())
	g.Vertices[getIdByType(c1)].CreateEdge(vertex)
	return cell
}

func (r React) CreateCompute2(c1 Cell, c2 Cell, compute func(value1, value2 int) int) ComputeCell {
	cell := &CCell{GCell: GCell{id: r.GenerateUniqueId(), value: compute(c1.Value(), c2.Value())}, callback: func(value int) { fmt.Println(value) }, compute2: compute, numOfIndepandants: 2}
	cell.inputs = append(cell.inputs, c1, c2)
	vertex := g.CreateVertex(cell)
	g.AddToGraph(vertex, cell.GetId())
	g.Vertices[getIdByType(c1)].CreateEdge(vertex)
	g.Vertices[getIdByType(c2)].CreateEdge(vertex)
	return cell
}

func getValueByType(c interface{}) int {
	var value int

	switch c.(type) {
	case *CCell:
		value = c.(*CCell).Value()
	case *ICell:
		value = c.(*ICell).Value()
	}

	return value
}

func getIdByType(c Cell) string {
	id := ""
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
