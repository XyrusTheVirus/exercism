package react

import (
	"fmt"
)

type GCell struct {
	value int
}

type ICell struct {
	GCell
}

type CCell struct {
	GCell
	callback func(int)
}

func (c ICell) SetValue(value int) {
	c.value = value
}

func (c CCell) AddCallback(callback func(value int)) Canceler {
	c.callback = callback
	return c
}

func (c CCell) Cancel() {
	c.callback = nil
}

func (c GCell) CreateInput(value int) InputCell {
	return ICell{GCell: GCell{value: value}}
}

func (c GCell) CreateCompute1(c1 Cell, compute func(value int) int) ComputeCell {
	return CCell{GCell: GCell{compute(c1.Value())}, callback: func(value int) { fmt.Println(value) }}
}

func (c GCell) CreateCompute2(c1 Cell, c2 Cell, compute func(value1, value2 int) int) ComputeCell {
	return CCell{GCell: GCell{value: compute(c1.Value(), c2.Value())}, callback: func(value int) { fmt.Println(value) }}
}

func (c GCell) Value() int {
	return c.value
}

func New() Reactor {
	return GCell{}
}
