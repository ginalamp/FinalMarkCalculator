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
	fmt.Println("Welcome to Gina's Mark Calculator!")
	fmt.Println("You may enter 'exit' or 'quit' at any input point if you wish to quit the program.")
out:
	for {
		inputType := Empty
		for {
			fmt.Println("-----------------------------MAIN MENU-----------------------------------")
			in := utils.ReadInput("\tEnter 0 to import a csv,\n\tEnter 1 to get the guidelines of how your csv must look like\n\tEnter exit to quit the program")
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
			hasProfile := utils.ReadInput("\nDo you have a profile?\n\tEnter 0 if you have one,\n\tEnter 1 if you don't have one but wish to make one,\n\tEnter any other key if you don't have one and don't wish to make one:")
			// allow user to quit the program
			if utils.UserExit(hasProfile) {
				break out
			}
			profile := oop.NewProfile("")
			switch hasProfile {
			case "0":
				// profile = userHasProfile()
				if userHasProfile() == "m" {
					continue
				}
			case "1":
				userNewProfile()
			default:
				userNoProfile()
				if run(profile) == "exit" {
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
	// assume that there is a <profileusername>marks.csv file outputted for the user
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
		userAction := utils.ReadInput("\tEnter 0 to view your results\n\tEnter 1 to update your results (import a new csv with your updated results)\n\tEnter any other key to go back to the main menu")

		switch userAction {
		case "0":
			// view current results
			file := utils.ReadCsvFile(OutputDirectory + username + "_marks.csv")
			for _, line := range file {
				fmt.Println(line)
			}
		case "1":
			// update results
			profile := oop.NewProfile(username)
			run(profile)
		default:
			return ""
		}
	}
}

// case if user want's to make a profile
func userNewProfile() oop.Profile {
	username := ""
	fmt.Print("\nGreat, let's create a profile for you! ")
out:
	for {
		username = utils.ReadInput("\nWhat is your username?")

		// allow user to quit
		if username == "exit" || username == "quit" {
			return oop.NewProfile(username)
		}
		// username may not be empty
		if username == "" {
			fmt.Println("\nNOTE: Your username may not be empty.")
			continue out
		}
		// username needs to be unique
		usernameFound := false
		for _, profile := range utils.ReadCsvFile("profiles.csv") {

			if strings.EqualFold(profile[0], username) {
				fmt.Printf("\nNOTE: The username %v is already used - please choose a unique username\n", username)
				usernameFound = true
				continue out
			}
		}
		if !usernameFound {
			break
		}
	}

	fmt.Printf("\nHi, %v! Happy to have you here!\n", username)
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
	// log.Println("Appending succeeded")

	return oop.NewProfile(username)
}

// case if user doesn't have a profile and doesn't want to make one
func userNoProfile() {
	fmt.Println("\nThat's okay, you can always create a profile another time.")
}

func run(profile oop.Profile) string {
	// get csv name
	csvFile := utils.ReadInput("\nEnter the relative path of your csv file (default is marks.csv - just click the Enter button to access the default):")
	// allow user to quit the program
	if utils.UserExit(csvFile) {
		return "exit"
	}
	// default set to "marks.csv"
	if len(csvFile) == 0 {
		csvFile = "marks.csv"
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

	// TODO: check if user wants to save results to profile
	profile.Degree = oop.Degree{Name: degree.Name, Modules: &modules}

	fmt.Println("\nOutputting your results to .csv...")
	// utils.OutputCsv(modules, profile)
	utils.OutputCsv(modules, profile, degree)
	return ""
}
