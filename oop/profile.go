package oop

// Profile

// Profile
type Profile struct {
	Username string
	Degree   Degree
}

// init profile
func NewProfile(username string) Profile {
	profile := Profile{}
	profile.Username = username
	profile.Degree = Degree{}
	return profile
}
