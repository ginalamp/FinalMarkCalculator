package main

import (
	"fmt"

	"ginalamp-mark-tracker/oop"   // Profile, Degree, Module, Component
	"ginalamp-mark-tracker/utils" // utility functions
	"strings"
)

// global variables
var Error float64 = -1
var Empty float64 = -1
var Float64Type int = 64
var OutputDirectory string = "marks/"

// run program
func main() {
	fmt.Println("Welcome to Gina's Mark Calculator")
	fmt.Println("You may enter 'exit' or 'quit' at any input point if you wish to quit the program")
out:
	for {
		inputType := Empty
		for {
			in := utils.ReadInput("Enter 0 to import a csv, Enter 1 to manually add entries:")
			// allow user to quit the program
			if utils.UserExit(in) {
				break out
			}
			// check if valid input given
			inputType = utils.StringToFloat(in)
			if !(inputType == 0 || inputType == 1) {
				continue
			}
			break
		}
		switch inputType {
		case 0:
			// check if user has a profile
			hasProfile := utils.ReadInput("Do you have a profile? Enter 0 if you have one, Enter 1 if you don't have one but wish to make one, Enter any other key if you don't have one and don't wish to make one:")
			// allow user to quit the program
			if utils.UserExit(hasProfile) {
				break out
			}
			profile := oop.NewProfile("")
			switch hasProfile {
			case "0":
				// profile = userHasProfile()
				userHasProfile()
			case "1":
				profile = userNewProfile()
				if run(profile) == "exit" {
					break out
				}
			default:
				userNoProfile()
				if run(profile) == "exit" {
					break out
				}
			}
		case 1:
			utils.InputTerminal()
		default:
			fmt.Println("why u like dis")
		}

		// check if user wants to continue with the program
		run := strings.ToLower(utils.ReadInput("Would you like to calculate another profile's mark? (Enter 'Y' if you do, otherwise enter any key to exit the program):"))
		if run == "yes" || run == "y" {
			continue
		}
		break
	}
	fmt.Println("Thank you for using Gina's mark calculator!")
}

// case if user has a profile
func userHasProfile() {
	name := utils.ReadInput("What is your name?")
	fmt.Printf("Welcome back, %v!\n", name)
	// assume that there is a <profileName>marks.csv file outputted for the user

	userFound := false
	for _, profile := range utils.ReadCsvFile("profiles.csv") {
		if profile[0] == name {
			fmt.Printf("User %v found\n", name)
			userFound = true
		}
		fmt.Println(profile)
	}
	if !userFound {
		fmt.Println("Oops... seems like we cannot find your profile")
		// TODO: do something
		return
	}
	file := utils.ReadCsvFile(OutputDirectory + name + "_marks.csv")
	for _, line := range file {
		fmt.Println(line)
	}

}

// case if user want's to make a profile
func userNewProfile() oop.Profile {
	name := utils.ReadInput("Great, let's create a profile for you! What is your name?")
	fmt.Printf("Hi, %v! Happy to have you here!\n", name)
	return oop.NewProfile(name)
}

// case if user doesn't have a profile and doesn't want to make one
func userNoProfile() {
	fmt.Println("That's okay, you can always create a profile another time.")
}

func run(profile oop.Profile) string {
	fmt.Println("Profile ---> ", profile)
	// get csv name
	csvFile := utils.ReadInput("Enter the name of your mark csv file (default is marks.csv):")
	// allow user to quit the program
	if utils.UserExit(csvFile) {
		return "exit"
	}
	// default set to "marks.csv"
	if len(csvFile) == 0 {
		csvFile = "marks.csv"
	}
	modules := utils.InputCsv(csvFile)

	// set calculated module mark
	for i, module := range modules {
		modules[i].Mark = module.CalculateMark()
	}
	// set calculated degree mark
	degree := oop.Degree{Modules: &modules}
	degree.Mark = degree.CalculateMark()

	utils.OutputTerminal(modules, degree)

	// check if user wants to save results to profile
	profile.Degree = oop.Degree{Modules: &modules}

	fmt.Println("Outputting results to csv...")
	utils.OutputCsv(modules, profile)
	return ""
}
