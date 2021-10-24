package main

// main running of program

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
var marksFilePath string = "marks.csv"

// run program
func main() {
	fmt.Println("Welcome to Gina's Mark Calculator!")
	fmt.Println("You may enter 'exit' or 'quit' at any input point if you wish to quit the program.")
out:
	for {
		inputType := Empty
		for {
			fmt.Println("-----------------------------MAIN MENU-----------------------------------")
			in := utils.ReadInput("\tEnter 0 to start\n\tEnter 1 to get the guidelines of how your csv must look like\n\tEnter exit to quit the program")
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
			hasProfile := utils.ReadInput("\nDo you have a profile?\n\tEnter 0 if you have a profile,\n\tEnter 1 if you would like to make a new profile,\n\tEnter any other key to continue without a profile")
			// allow user to quit the program
			if utils.UserExit(hasProfile) {
				break out
			}
			profile := oop.NewProfile("")
			switch hasProfile {
			case "0":
				choice := userHasProfile()
				// allow user to quit the program
				if utils.UserExit(choice) {
					break out
				}
				// main menu
				if choice == "m" {
					continue
				}
			case "1":
				choice := userNewProfile()
				// allow user to quit the program
				if utils.UserExit(choice) {
					break out
				}
				// main menu
				if choice == "m" {
					continue
				}
			default:
				userNoProfile()
				if run(profile, marksFilePath) == "exit" {
					break out
				}
			}
		case 1:
			utils.InputCsvGuidelines()
			continue
		}

		// check if user wants to continue with the program
		run := strings.ToLower(utils.ReadInput("\nWould you like to calculate another profile's mark?\n\tEnter 'Y' if you do\n\tEnter any key to exit the program"))
		if run == "yes" || run == "y" {
			continue
		}
		break
	}
	fmt.Println("\nThank you for using Gina's mark calculator!")
}

// case if user has a profile
func userHasProfile() string {
	username := ""
out:
	for {
		username = utils.ReadInput("\nWhat is your username?")

		userFound := false
		for _, profile := range utils.ReadCsvFile("profiles.csv") {
			if strings.EqualFold(profile[0], username) {
				log.Printf("User %v found\n", username)
				userFound = true
				break out
			}
		}
		if !userFound {
			fmt.Printf("\nOops... seems like we don't have '%v' in out database. Make sure you've spelt it correctly\n", username)
			menu := utils.ReadInput("\tEnter 'm' to go back to the main menu\n\tEnter any other key to retry entering your username.")
			if menu == "m" || menu == "menu" {
				return "m"
			}
			continue
		}
	}

	fmt.Printf("\nWelcome back, %v!\n", username)
	for {
		fmt.Println("\n---------------------------PROFILE INPUT MENU----------------------------------")
		userAction := utils.ReadInput("\tEnter 0 to view your results\n\tEnter 1 to update your results (import a new csv with your updated results)\n\tEnter any other key to go back to the main menu")

		// assume that there is a <profileusername>_marks.csv file outputted for the user
		marksFilePath := OutputDirectory + username + "_marks.csv"
		switch userAction {
		case "0":
			// view current results
			modules := utils.InputCsv(marksFilePath)
			// calculate reults
			for i, module := range modules {
				modules[i].Mark = module.CalculateMark()
			}
			degree := oop.Degree{Modules: &modules}
			degree.Mark = degree.CalculateMark()
			degree.Name = utils.DegreeName

			// output results
			utils.OutputTerminal(modules, degree)
		case "1":
			// update results
			profile := oop.NewProfile(username)

			run(profile, marksFilePath)
		default:
			return ""
		}
	}
}

// case if user want's to make a profile
func userNewProfile() string {
	username := ""
	fmt.Print("\nGreat, let's create a profile for you! ")
out:
	for {
		username = strings.ToLower(utils.ReadInput("\nChoose a username:"))

		// allow user to quit
		if utils.UserExit(username) {
			return "exit"
		}
		// username may not be empty
		if username == "" {
			fmt.Println("\nNOTE: Your username may not be empty.")
			menu := utils.ReadInput("\tEnter 'm' to go back to the main menu\n\tEnter any other key to retry entering your username.")
			if menu == "m" || menu == "menu" {
				return "m"
			}
			// allow user to quit
			if utils.UserExit(username) {
				return "exit"
			}
			continue out
		}
		// username needs to be unique
		usernameFound := false
		for _, profile := range utils.ReadCsvFile("profiles.csv") {

			if strings.EqualFold(profile[0], username) {
				fmt.Printf("\nNOTE: The username '%v' is already used - please choose a unique username\n", username)
				usernameFound = true
				menu := utils.ReadInput("\tEnter 'm' to go back to the main menu\n\tEnter any other key to retry entering your username.")
				if menu == "m" || menu == "menu" {
					return "m"
				}
				// allow user to quit
				if utils.UserExit(username) {
					return "exit"
				}
				continue out
			}
		}
		// username not in records
		if !usernameFound {
			break
		}
	}
	// create new profile
	fmt.Printf("\nHi, %v! Happy to have you here!\n", username)
	profile := oop.NewProfile(username)
	run(profile, marksFilePath)

	// append user data to profiles.csv
	f, err := os.OpenFile("profiles.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var data [][]string
	data = append(data, []string{profile.Username})

	w := csv.NewWriter(f)
	w.WriteAll(data)
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	return ""
}

// case if user doesn't have a profile and doesn't want to make one
func userNoProfile() {
	fmt.Println("\nThat's okay, you can always create a profile another time.")
}

// run marks csv input & output functions
func run(profile oop.Profile, marksFilePath string) string {
	// get csv name
	csvFile := utils.ReadInput("\nEnter the relative path of your csv file\n\tIf you have a profile the default is set to marks/<username>_marks\n\tThe default is marks.csv if you just created a profile, or are an anonymous user")
	// allow user to quit the program
	if utils.UserExit(csvFile) {
		return "exit"
	}
	// default file path
	if len(csvFile) == 0 {
		csvFile = marksFilePath
	}
	// modules := utils.InputCsv(csvFile)
	modules := utils.InputCsv(csvFile)

	// set calculated module mark
	for i, module := range modules {
		modules[i].Mark = module.CalculateMark()
	}
	// set calculated degree mark
	degree := oop.Degree{Modules: &modules}
	degree.Mark = degree.CalculateMark()
	degree.Name = utils.DegreeName

	utils.OutputTerminal(modules, degree)
	profile.Degree = oop.Degree{Name: degree.Name, Modules: &modules}

	fmt.Println("\nWriting your results to .csv...")
	utils.OutputCsv(modules, profile, degree)
	return ""
}
