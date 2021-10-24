package oop

// import "fmt"

// Profile
type Profile struct {
	Username string
	Degree   Degree
	// InputFile  string
	// OutputFile string
}

// init profile
func NewProfile(username string) Profile {
	profile := Profile{}
	profile.Username = username
	profile.Degree = Degree{}
	return profile
}
