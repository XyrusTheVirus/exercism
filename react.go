package react

import (
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
	inputs    []Cell
	callbacks []func(int)
	compute1  func(int) int
	compute2  func(int, int) int
}

type CCancel struct {
	callback func(int)
}

func (c *GCell) GetId() string {
	return c.id
}

func (r React) GenerateUniqueId() string {
	return guuid.New().String()
}

func (c *GCell) SetValue(value int) {
	if c.value != value {
		c.value = value
		propagate()
	}
}

func propagate() {
	sortedGraph := g.GetGraphDependenies()
	for _, v := range sortedGraph {
		if reflect.TypeOf(v.GetCell()).String() == "*react.CCell" {
			currentVal := v.GetCell().(*CCell).Value()
			switch len(v.GetCell().(*CCell).inputs) {
			case 1:
				v.GetCell().(*CCell).value = v.GetCell().(*CCell).compute1(getValueByType(v.GetCell().(*CCell).inputs[0]))
				executeCallback(v.GetCell(), currentVal)
			case 2:
				v.GetCell().(*CCell).value = v.GetCell().(*CCell).compute2(getValueByType(v.GetCell().(*CCell).inputs[0]), getValueByType(v.GetCell().(*CCell).inputs[1]))
				executeCallback(v.GetCell(), currentVal)
			}
		}
	}
}

func (c *CCell) AddCallback(callback func(value int)) Canceler {
	c.callbacks = append(c.callbacks, callback)
	return &CCancel{callback: callback}
}

func executeCallback(c interface{}, currentVal int) {
	if shouldExecuteCallback(c, currentVal) == true {
		for _, callback := range c.(*CCell).callbacks {
			if callback != nil {
				callback(c.(*CCell).Value())
			}
		}
	}
}

func shouldExecuteCallback(c interface{}, currentVal int) bool {
	return currentVal != c.(*CCell).Value()
}

func (c *CCancel) Cancel() {
	c.callback = nil
}

func (r React) CreateInput(value int) InputCell {
	cell := &ICell{GCell: GCell{id: r.GenerateUniqueId(), value: value}}
	g.AddToGraph(g.CreateVertex(cell), cell.GetId())
	return cell
}

func (r React) CreateCompute1(c1 Cell, compute func(value int) int) ComputeCell {
	cell := &CCell{GCell: GCell{id: r.GenerateUniqueId(), value: compute(c1.Value())}, compute1: compute}
	cell.inputs = append(cell.inputs, c1)
	vertex := g.CreateVertex(cell)
	g.AddToGraph(vertex, cell.GetId())
	g.Vertices[getIdByType(c1)].CreateEdge(vertex)
	return cell
}

func (r React) CreateCompute2(c1 Cell, c2 Cell, compute func(value1, value2 int) int) ComputeCell {
	cell := &CCell{GCell: GCell{id: r.GenerateUniqueId(), value: compute(c1.Value(), c2.Value())}, compute2: compute}
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
