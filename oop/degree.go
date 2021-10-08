package oop

import "fmt"

var Error float64 = 0

// degree is a bunch of modules
type Degree struct {
	// Name    string
	Modules []Module
}

// assumption: all modules weigh the same amount to the degree
type Module struct {
	Name string
	// Weight     int // weight of the module in the total degree
	Mark       float64
	Components []Component
}

// e.g. assignment
type Component struct {
	// Name   string
	Mark   float64
	Weight float64 // % weight of the assignment in the module
}

// init module
func NewModule(moduleName string) Module {
	module := Module{}
	module.Name = moduleName
	module.Mark = 0
	return module
}

// add module component
func AddModuleComponent(mark, weight float64) Component {
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
		return Error
	}
	sum := 0.0
	for _, component := range module.Components {
		sum += component.Mark * component.Weight / 100
	}
	// fmt.Println(module.Components)

	return sum
}
