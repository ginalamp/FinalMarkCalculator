package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

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
			in := utils.ReadInput("Mark input:\n\tEnter 0 to import a csv,\n\tEnter 1 to manually add entries:")
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
			hasProfile := utils.ReadInput("Do you have a profile?\n\tEnter 0 if you have one,\n\tEnter 1 if you don't have one but wish to make one,\n\tEnter any other key if you don't have one and don't wish to make one:")
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
		run := strings.ToLower(utils.ReadInput("Would you like to calculate another profile's mark?\n\tEnter 'Y' if you do\n\tEnter any key to exit the program"))
		if run == "yes" || run == "y" {
			continue
		}
		break
	}
	fmt.Println("Thank you for using Gina's mark calculator!")
}

// case if user has a profile
func userHasProfile() {
	username := utils.ReadInput("What is your username?")
	fmt.Printf("Welcome back, %v!\n", username)
	// assume that there is a <profileusername>marks.csv file outputted for the user

	userFound := false
	for _, profile := range utils.ReadCsvFile("profiles.csv") {
		if profile[0] == username {
			log.Printf("User %v found\n", username)
			userFound = true
		}
	}
	if !userFound {
		fmt.Println("Oops... seems like we cannot find your profile")
		// TODO: do something
		return
	}
	file := utils.ReadCsvFile(OutputDirectory + username + "_marks.csv")
	for _, line := range file {
		fmt.Println(line)
	}
}

// case if user want's to make a profile
func userNewProfile() oop.Profile {
	username := ""
	fmt.Print("Great, let's create a profile for you! ")
out:
	for {
		username := utils.ReadInput("What is your username?")

		// allow user to quit
		if username == "exit" || username == "quit" {
			return oop.NewProfile(username)
		}
		// username needs to be unique
		usernameFound := false
		for _, profile := range utils.ReadCsvFile("profiles.csv") {
			if profile[0] == username {
				fmt.Printf("The username %v is already used - please choose a unique username\n", username)
				usernameFound = true
				continue out
			}
		}
		if !usernameFound {
			break
		}
	}

	fmt.Printf("Hi, %v! Happy to have you here!\n", username)
	profile := oop.NewProfile(username)

	run(profile)

	// append user data to profiles.csv
	f, err := os.OpenFile("profiles.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var data [][]string
	password := "123" // TODO make this user input
	data = append(data, []string{profile.Username, password})

	w := csv.NewWriter(f)
	w.WriteAll(data)
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	log.Println("Appending succeeded")

	return oop.NewProfile(username)
}

// case if user doesn't have a profile and doesn't want to make one
func userNoProfile() {
	fmt.Println("That's okay, you can always create a profile another time.")
}

func run(profile oop.Profile) string {
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
