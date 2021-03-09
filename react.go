package react

import (
	"react/graph"
	"reflect"

	guuid "github.com/google/uuid"
)

// Holds the graph structure
var g *graph.Graph = &graph.Graph{Vertices: make(map[string]*graph.Vertex)}

type React struct{}

// Gcell struct acts as father of ICell (input) & CCell (compute) and holds their common attributes
type GCell struct {
	value int
	id    string
}

// Defines InputCell interface behavior
type ICell struct {
	GCell // Inherit Gcell properties
}

// Defines ComputeCell interface  behavior
type CCell struct {
	GCell                        // Inherits Gcelll properties
	inputs    []Cell             // Stores all compute cell's independent cells
	callbacks []*Callback        // Stores all compute cell calbacks
	compute1  func(int) int      // Stores computation function with one independent
	compute2  func(int, int) int // Stores computation function with two independents
}

// Defines Cancel interface behavior
type CCancel struct {
	callback *Callback
}

// Defines a callback structure for CCancel entity
type Callback struct {
	f func(int) // Stores the CCancel entity callback
}

// Returns cell's id (string)
func (c *GCell) GetId() string {
	return c.id
}

// Generates unique id for the cell's id to avoid collision in the graph's map holds all cells
// Returns string
func (r React) GenerateUniqueId() string {
	return guuid.New().String()
}

// Setts the value of the cell
func (c *GCell) SetValue(value int) {
	if c.value != value {
		c.value = value
		propagate()
	}
}

// This function responsible to update all the desired cells, according to their dependencies cells value modification.
// The values modification process goes as follow:
// 1. Send the cells graph to a topological sorting
// 2. Go through the sorted array
// 3. Update cell value according to its correct order
func propagate() {
	// The graph received from the topological sort
	sortedGraph := g.GetGraphDependenies()
	for _, v := range sortedGraph {
		// Input cells shouldn't be updated at all, only via 'SetValue' function
		if reflect.TypeOf(v.GetCell()).String() == "*react.CCell" {
			currentVal := v.GetCell().(*CCell).Value()
			// Determine to which compute function we should call, according to its list of parameter
			switch len(v.GetCell().(*CCell).inputs) {
			case 1:
				// Update cell value
				v.GetCell().(*CCell).value = v.GetCell().(*CCell).compute1(getValueByType(v.GetCell().(*CCell).inputs[0]))
				// Call the callback to notify cell value modification
				executeCallback(v.GetCell(), currentVal)
			case 2:
				// Update cell value
				v.GetCell().(*CCell).value = v.GetCell().(*CCell).compute2(getValueByType(v.GetCell().(*CCell).inputs[0]), getValueByType(v.GetCell().(*CCell).inputs[1]))
				// Call the callback to notify cell value modification
				executeCallback(v.GetCell(), currentVal)
			}
		}
	}
}

// Adds callback to designated compute cell
func (c *CCell) AddCallback(callback func(value int)) Canceler {
	function := &Callback{f: callback}
	c.callbacks = append(c.callbacks, function)
	return &CCancel{callback: function}
}

// Executes callback compute cell callback
func executeCallback(c interface{}, currentVal int) {
	if shouldExecuteCallback(c, currentVal) == true {
		callbacks := c.(*CCell).callbacks
		for _, callback := range callbacks {
			if callback.f != nil {
				callback.f(c.(*CCell).Value())
			}
		}
	}
}

func shouldExecuteCallback(c interface{}, currentVal int) bool {
	return currentVal != c.(*CCell).Value()
}

func (c *CCancel) Cancel() {
	c.callback.f = nil
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
