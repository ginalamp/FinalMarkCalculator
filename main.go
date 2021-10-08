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

var Error int = 0

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

func main() {
	fmt.Println("Welcome to Gina's Mark Calculator")
	inputType := stringToInt(readInput("Enter 0 to import a csv, Enter 1 to manually add entries:"))

	if inputType == 0 {
		records := readCsvFile("temp.csv")
		// numModules := len(records) - 1
		// fmt.Println(len(records))

		// add modules (1 module per row except first row)
		var modules []oop.Module
		for i, row := range records {
			// skip title/1st row
			if i == 0 {
				continue
			}
			moduleData := strings.Split(row[0], ";")
			// add init module names to module slice
			module := oop.NewModule(moduleData[0])

			for i := 2; i < len(moduleData[1:]); i += 2 {
				mark := percentageToInt(moduleData[i-1])
				weight := percentageToInt(moduleData[i])
				module.Components = append(module.Components, oop.AddModuleComponent(mark, weight))
			}
			// x := percentageToInt(moduleData[1])
			// y := percentageToInt(moduleData[2])
			// module.Components = append(module.Components, oop.AddModuleComponent(x, y))
			modules = append(modules, module)
			// modules[i].Components = addComponents(moduleData)
			// fmt.Println(oop.NewModule(moduleData[0]))
		}
		fmt.Println(modules)
		// for _, module := range modules {
		// 	fmt.Println(module.Components)

		// }

		// for i := 0; i < numModules; i++ {

		// 	// oop.Component{Name: "Accounting", Mark: 77, Weight: 50}
		// }
	} else if inputType == 1 {
		// read input from terminal
		numModules := 0
		for numModules < 1 {
			numModules = stringToInt(readInput("How many modules do you have?"))
		}
		fmt.Println(numModules)
	}

	// basic assignment initialisation
	a2 := oop.Component{Mark: 77, Weight: 50}
	fmt.Println(a2)
}

// convert 50% to 50
func percentageToInt(s string) int {
	r := strings.Replace(s, "%", "", -1)
	num, err := strconv.Atoi(r)
	if err != nil {
		return 0
	}
	return num
}

// string to int
func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Invalid input. Please type in a number (integer)")
		// handle error
		fmt.Println(err)
		return (Error)
	}
	return (i)
}
