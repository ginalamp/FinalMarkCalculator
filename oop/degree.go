package oop

import "fmt"

type Degree struct {
	Name    string
	Modules []Module
}

// assumption: all modules weigh the same amount to the degree
type Module struct {
	Name string
	// Weight     int // weight of the module in the total degree
	Mark       int
	Components []Component
}

// e.g. assignment
type Component struct {
	// Name   string
	Mark   int
	Weight int // % weight of the assignment in the module
}

// init module
func NewModule(moduleName string) Module {
	module := Module{}
	module.Name = moduleName
	module.Mark = 0
	return module
}

// add module component
func AddModuleComponent(mark, weight int) Component {
	component := Component{}
	component.Mark = mark
	component.Weight = weight
	return component
}

// calculate final mark depending on type
type Marker interface {
	CalculateMark() float64
}

// calculate degree overall mark
func (degree *Degree) CalculateMark() float64 {
	if degree == nil {
		fmt.Println("<nil>")
		return 0
	}
	return 80
}

// calculate module overall mark
func (module *Module) CalculateMark() float64 {
	if module == nil {
		fmt.Println("<nil>")
		return 0
	}
	return 80
}
