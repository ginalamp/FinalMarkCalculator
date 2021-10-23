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
func outputTerminal(modules []oop.Module) {
	// calculate module mark
	for _, module := range modules {
		fmt.Printf("%v: %v%%\n", module.Name, module.CalculateMark())
	}
}

// output results to csvq
// https://golangcode.com/write-data-to-a-csv-file/
func outputCsv(modules []oop.Module) {
	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close() // always close the file

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, module := range modules {
		value := []string{module.Name, float2string(module.CalculateMark())}
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

// run program
func main() {
	fmt.Println("Welcome to Gina's Mark Calculator")
	for {
		inputType := Empty
		for {
			inputType = stringToFloat(readInput("Enter 0 to import a csv, Enter 1 to manually add entries:"))
			if !(inputType == 0 || inputType == 1) {
				continue
			}
			break
		}
		switch inputType {
		case 0:
			csvFile := readInput("Enter the name of your mark csv file (default is marksInput.csv)")
			if len(csvFile) == 0 {
				csvFile = "marksInput.csv"
			}
			modules := inputCsv(csvFile)
			outputTerminal(modules)
			fmt.Println("Outputting results to csv")
			outputCsv(modules)
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
		fmt.Println("Thank you for using Gina's mark calculator!")
		break
	}
}

// **************************************************************************************
// *** utility functions
// **************************************************************************************

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
		// handle error
		// fmt.Println(err)
		return (Error)
	}
	return (i)
}

// check if error != nil
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// convert float64 to string with 0 decimals
func float2string(f float64) string {
	return fmt.Sprintf("%.0f", f)
}
