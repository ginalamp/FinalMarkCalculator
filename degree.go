package main

type Degree struct {
	DegreeName string
	Modules    []Module
}

type Module struct {
	ModuleName  string
	Weight      int // weight of the module in the total degree
	Mark        int
	Assignments []Assignment
}

type Assignment struct {
	AssignmentName string
	Mark           int
	Weight         int // % weight of the assignment in the module
}
