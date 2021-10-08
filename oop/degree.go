package oop

type Degree struct {
	Name    string
	Modules []Module
}

type Module struct {
	Name       string
	Weight     int // weight of the module in the total degree
	Mark       int
	Components []Component
}

// e.g. assighment
type Component struct {
	Name   string
	Mark   int
	Weight int // % weight of the assignment in the module
}
