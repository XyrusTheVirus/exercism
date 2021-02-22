package react

import (
	"fmt"
	"react/graph"
	"reflect"

	guuid "github.com/google/uuid"
)

var g graph.Graph

type React struct {
	g graph.Graph
}

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

func (c GCell) GetId() string {
	return c.id
}

func (r React) GenerateUniqueId() string {
	return guuid.New().String()
}

func (c GCell) SetValue(value int) {
	c.value = value
	sortedGraph := c.g.GetGraphDependenies()
	for _, v := range sortedGraph {
		if reflect.TypeOf(v.cell).String() == "CCell" {
			switch n := len(v.adjacents); n {
			case 1:
				v.SetValue(v.compute1(v.adjacents[0]))
			case 2:
				v.SetValue(v.compute2(v.adjacents[0], v.adjacents[1]))
			}
		}
	}
}

func 
func (c CCell) AddCallback(callback func(value int)) Canceler {
	c.callback = callback
	return c
}

func (c CCell) Cancel() {
	c.callback = nil
}

func (r React) CreateInput(value int) InputCell {
	cell := ICell{GCell: GCell{id: r.GenerateUniqueId(), value: value}}
	r.g.AddToGraph(r.g.CreateVertex(value), cell.GetId())
	return cell
}

func (r React) CreateCompute1(c1 Cell, compute func(value int) int) ComputeCell {
	cell := CCell{GCell: GCell{id: r.GenerateUniqueId(), value: compute(c1.Value())}, callback: func(value int) { fmt.Println(value) }, compute1: compute}
	vertex := r.g.CreateVertex(cell.Value())
	r.g.AddToGraph(vertex, cell.GetId())
	r.g.Vertices[c1.(CCell).GetId()].CreateEdge(vertex)
	return cell
}

func (r React) CreateCompute2(c1 Cell, c2 Cell, compute func(value1, value2 int) int) ComputeCell {
	cell := CCell{GCell: GCell{value: compute(c1.Value(), c2.Value())}, callback: func(value int) { fmt.Println(value) }, compute2: compute}
	vertex := r.g.CreateVertex(cell.Value())
	r.g.AddToGraph(vertex, cell.GetId())
	r.g.Vertices[c1.(CCell).GetId()].CreateEdge(vertex)
	r.g.Vertices[c2.(CCell).GetId()].CreateEdge(vertex)
	return cell
}

func (c GCell) Value() int {
	return c.value
}

func New() Reactor {
	return React{}
}
