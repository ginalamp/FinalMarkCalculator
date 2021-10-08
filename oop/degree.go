package oop

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

// init module
func NewModule(moduleName string) Module {
	module := Module{}
	module.Name = moduleName
	module.Mark = 0
	return module
}

// func addModuleComponent(componentName string, mark, weight int) Component {
// 	component := Component{}
// 	component.Name := componentName
// }

// e.g. assighment
type Component struct {
	Name   string
	Mark   int
	Weight int // % weight of the assignment in the module
}
