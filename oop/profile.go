package oop

// import "fmt"

// Profile
type Profile struct {
	Name   string
	Degree Degree
	// InputFile  string
	// OutputFile string
}

// init profile
func NewProfile(name string) Profile {
	profile := Profile{}
	profile.Name = name
	profile.Degree = Degree{}
	return profile
}
