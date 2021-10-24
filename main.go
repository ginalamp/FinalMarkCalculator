package main

import (
	"bufio" // user terminal input
	"encoding/csv"
	"fmt"
	"log"

	"ginalamp-mark-tracker/oop"
	"os"      // user input
	"strconv" // type conversion
	"strings"
)

// global variables
var Error float64 = -1
var Empty float64 = -1
var Float64Type int = 64

// run program
func main() {
	fmt.Println("Welcome to Gina's Mark Calculator")
	fmt.Println("You may enter 'exit' or 'quit' at any input point if you wish to quit the program")
out:
	for {
		inputType := Empty
		for {
			in := readInput("Enter 0 to import a csv, Enter 1 to manually add entries:")
			// allow user to quit the program
			if in == "exit" || in == "quit" {
				break out
			}

			inputType = stringToFloat(in)
			if !(inputType == 0 || inputType == 1) {
				continue
			}
			break
		}
		switch inputType {
		case 0:
			csvFile := readInput("Enter the name of your mark csv file (default is marks.csv), Enter exit to quit the program:")
			// allow user to quit the program
			if csvFile == "exit" || csvFile == "quit" {
				break out
			}

			if len(csvFile) == 0 {
				csvFile = "marks.csv"
			}
			modules := inputCsv(csvFile)

			// set calculated module mark
			for i, module := range modules {
				modules[i].Mark = module.CalculateMark()
			}
			// set calculated degree mark
			degree := oop.Degree{Modules: &modules}
			degree.Mark = degree.CalculateMark()

			outputTerminal(modules, degree)

			// check if user wants to save results to profile
			profile := oop.NewProfile("Pietie")
			profile.Degree = oop.Degree{Modules: &modules}

			fmt.Println("Outputting results to csv...")
			outputCsv(modules, profile)
		case 1:
			inputTerminal()
		default:
			fmt.Println("why u like dis")
		}

		// check if user wants to continue with the program
		run := strings.ToLower(readInput("Would you like to calculate another profile's mark? (Enter 'Y' if you do, otherwise enter any key to exit the program)"))
		if run == "yes" || run == "y" {
			continue
		}
		break
	}
	fmt.Println("Thank you for using Gina's mark calculator!")
}

// **************************************************************************************
// *** input/output
// **************************************************************************************

// process terminal input
func inputTerminal() {
	// read input from terminal
	numModules := 0.0
	for numModules < 1 {
		numModules = stringToFloat(readInput("How many modules do you have?"))
	}
	fmt.Println(numModules)
}

// process csv input
func inputCsv(csvFile string) []oop.Module {
	records := readCsvFile(csvFile)
	// numModules := len(records) - 1

	// add modules (1 module per row except first row)
	var modules []oop.Module
	for i, row := range records {
		// skip title/1st row
		if i == 0 {
			continue
		}
		// convert string to slice
		moduleData := strings.Split(row[0], ";")

		// add init module names to module slice
		module := oop.NewModule(moduleData[0])

		// add marks and weights to module components
		for i := 2; i <= len(moduleData[1:]); i += 2 {
			mark := percentageToFloat(moduleData[i-1])
			weight := percentageToFloat(moduleData[i])
			module.Components = append(module.Components, oop.AddModuleComponent(mark, weight))
		}
		modules = append(modules, module)
	}
	return modules
}

// output Module final marks to terminal
func outputTerminal(modules []oop.Module, degree oop.Degree) {
	// calculate degree mark
	fmt.Printf("Your overall degree mark: %v%%\n", degree.Mark)

	// calculate module mark
	for _, module := range modules {
		fmt.Printf("%v: %v%%\n", module.Name, module.Mark)
	}
}

// output results to csvq
// https://golangcode.com/write-data-to-a-csv-file/
func outputCsv(modules []oop.Module, profile oop.Profile) {
	file, err := os.Create(profile.Name + "_marks.csv")
	checkError("Cannot create file", err)
	defer file.Close() // always close the file

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, module := range modules {
		value := []string{module.Name, floatToString(module.Mark)}
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

// read terminal input
func readInput(userPrompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(userPrompt + ": \n")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput) // remove whitespace

	// fmt.Println("input read: ", userInput)
	return userInput
}

// read csv file
// https://stackoverflow.com/questions/24999079/reading-csv-file-in-go
func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

// **************************************************************************************
// *** helper functions
// **************************************************************************************

// check if error != nil
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// convert "50%" to 50.0
func percentageToFloat(s string) float64 {
	r := strings.Replace(s, "%", "", -1)
	num, err := strconv.ParseFloat(r, Float64Type)
	if err != nil {
		return Error
	}
	return num
}

// string input to float
func stringToFloat(s string) float64 {
	i, err := strconv.ParseFloat(s, Float64Type)
	if err != nil {
		fmt.Println("Invalid input. Please type in a number (integer)")
		return (Error)
	}
	return (i)
}

// convert float64 to string with 0 decimals
func floatToString(f float64) string {
	return fmt.Sprintf("%.0f", f)
}
