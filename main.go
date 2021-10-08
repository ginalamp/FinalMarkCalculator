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
var Error float64 = 0
var Empty float64 = 0
var Float64Type int = 64

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
func inputCsv() {
	records := readCsvFile("marksInput.csv")
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
	fmt.Println(modules)
	// calculate module mark
	for _, module := range modules {
		fmt.Printf("%v: %v%%\n", module.Name, module.CalculateMark())
	}
	// fmt.Println(modules[0].CalculateMark())
}

// run program
func main() {
	fmt.Println("Welcome to Gina's Mark Calculator")
	inputType := stringToFloat(readInput("Enter 0 to import a csv, Enter 1 to manually add entries:"))

	switch inputType {
	case 0:
		inputCsv()
	case 1:
		inputTerminal()
	default:
		fmt.Println("why u like dis")
	}
}

// **************************************************************************************
// *** helper functions
// **************************************************************************************

// read terminal input
func readInput(userPrompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(userPrompt)
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

// convert "50%" to 50.0
func percentageToFloat(s string) float64 {
	r := strings.Replace(s, "%", "", -1)
	num, err := strconv.ParseFloat(r, Float64Type)
	if err != nil {
		return Empty
	}
	return num
}

// string input to float
func stringToFloat(s string) float64 {
	i, err := strconv.ParseFloat(s, Float64Type)
	if err != nil {
		fmt.Println("Invalid input. Please type in a number (integer)")
		// handle error
		fmt.Println(err)
		return (Error)
	}
	return (i)
}
