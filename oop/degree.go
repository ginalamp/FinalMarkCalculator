package oop

import "fmt"

var Error float64 = 0

// degree is a bunch of modules
type Degree struct {
	Name    string
	Mark    float64
	Modules *[]Module
}

// assumption: all modules weigh the same amount to the degree
type Module struct {
	Name string
	// Weight     int // weight of the module in the total degree
	Mark       float64
	Components []Component // TODO: change to pointer to component
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

// func DeleteModule(module *Module) {

// }

// func DeleteModuleComponent(component Component, module Module) {

// }

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
	sum := 0.0
	numModules := len(*degree.Modules)
	for _, module := range *degree.Modules {
		sum += module.Mark
	}
	return sum / float64(numModules)
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
	return sum
}
