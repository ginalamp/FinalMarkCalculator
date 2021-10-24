package utils

// input/output helper functions

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"ginalamp-mark-tracker/oop"
	"log"
	"os"
	"strings"
)

// global variables
var Error float64 = -1
var Empty float64 = -1
var Float64Type int = 64
var OutputDirectory string = "marks/"
var ColumnHeaders []string
var DegreeName string

// process terminal input
func InputTerminal() {
	// read input from terminal
	numModules := 0.0
	for numModules < 1 {
		numModules = StringToFloat(ReadInput("How many modules do you have?"))
	}
	fmt.Println(numModules)
	fmt.Println("work in progress...")
}

// // process csv input
// func InputCsv(csvFile string) []oop.Module {
// 	records := ReadCsvFile(csvFile)

// 	// add modules (1 module per row except first row)
// 	var modules []oop.Module
// 	for i, row := range records {
// 		// skip title/1st row
// 		if i == 0 {
// 			continue
// 		}
// 		// convert string to slice
// 		moduleData := strings.Split(row[0], ";")

// 		// add init module names to module slice
// 		module := oop.NewModule(moduleData[0])

// 		// add marks and weights to module components
// 		for i := 2; i <= len(moduleData[1:]); i += 2 {
// 			mark := PercentageToFloat(moduleData[i-1])
// 			weight := PercentageToFloat(moduleData[i])
// 			module.Components = append(module.Components, oop.AddModuleComponent(mark, weight))
// 		}
// 		modules = append(modules, module)
// 	}
// 	return modules
// }

// process csv input
func InputCsv2(csvFile string) []oop.Module {
	records := ReadCsvFile(csvFile)

	// add modules (1 module per row except first row)
	var modules []oop.Module
	for i, row := range records {
		// skip title/1st row
		if i == 0 || i == 1 {
			// set degree name
			if i == 0 {
				DegreeName = strings.Split(row[0], ";")[0]
			}
			// set column headers
			if i == 1 {
				ColumnHeaders = row
			}
			continue
		}
		// replace spaces with commas
		for i, item := range row {
			if item == " " {
				row[i] = ","
			}
		}
		// convert string to slice
		moduleData := strings.Split(row[0], ";")

		// add init module names to module slice
		module := oop.NewModule(moduleData[0])

		// add marks and weights to module components
		colStart := 3
		for i := colStart; i <= len(moduleData[1:]); i += 2 {
			mark := PercentageToFloat(moduleData[i-1])
			weight := PercentageToFloat(moduleData[i])

			// set empty columns to 0
			if weight == Empty {
				mark, weight = 0, 0
			}

			module.Components = append(module.Components, oop.AddModuleComponent(mark, weight))
		}
		modules = append(modules, module)
	}
	return modules
}

// output Module final marks to terminal
func OutputTerminal(modules []oop.Module, degree oop.Degree) {
	// calculate degree mark
	fmt.Printf("Your overall degree mark: %v%%\n", degree.Mark)

	// calculate module mark
	for _, module := range modules {
		fmt.Printf("%v: %v%%\n", module.Name, module.Mark)
	}
}

// output results to csv
// Output csv https://golangcode.com/write-data-to-a-csv-file/
// Create directory https://golangbyexample.com/create-directory-folder-golang/
// func OutputCsv(modules []oop.Module, profile oop.Profile) {
// 	// Create marks directory
// 	err := MakeDirectoryIfNotExists(OutputDirectory)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// output csv to file in directory
// 	fileExtension := "_marks.csv"
// 	if profile.Username == "" {
// 		// output only marks/marks.csv if user has empty profile
// 		fileExtension = "marks.csv"
// 	}
// 	file, err := os.Create(OutputDirectory + profile.Username + fileExtension)

// 	CheckError("Cannot create file", err)
// 	defer file.Close() // always close the file

// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	for _, module := range modules {
// 		value := []string{module.Name, FloatToString(module.Mark)}
// 		err := writer.Write(value)
// 		CheckError("Cannot write to file", err)
// 	}
// }

func OutputFullCsv(modules []oop.Module, profile oop.Profile, degree oop.Degree) {
	// Create marks directory
	err := MakeDirectoryIfNotExists(OutputDirectory)
	if err != nil {
		log.Fatal(err)
	}
	// output csv to file in directory
	fileExtension := "_marks.csv"
	if profile.Username == "" {
		// output only marks/marks.csv if user has empty profile
		fileExtension = "marks.csv"
	}
	file, err := os.Create(OutputDirectory + profile.Username + fileExtension)

	CheckError("Cannot create file", err)
	defer file.Close() // always close the file

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// degree output
	value := []string{Concat2StringsCsvFormat(degree.Name, FloatToString(degree.Mark))}
	err = writer.Write(value)
	CheckError("Cannot write to file", err)

	// set column headers
	err = writer.Write(ColumnHeaders)
	CheckError("Cannot write to file", err)

	// module output
	// for _, module := range modules {
	// 	value = []string{module.Name, FloatToString(module.Mark)}
	// 	err := writer.Write(value)
	// 	CheckError("Cannot write to file", err)
	// }
	// module output 2
	// moduleOutput := []string
	for _, module := range modules {
		value = []string{module.Name, FloatToString(module.Mark)}
		for _, component := range module.Components {
			value = append(value, FloatToString(component.Mark), FloatToString(component.Weight))
		}
		value = []string{strings.Join(value, ";")}
		err := writer.Write(value)
		CheckError("Cannot write to file", err)
	}
}

// read terminal input
func ReadInput(userPrompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(userPrompt + "\n")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput) // remove whitespace

	// fmt.Println("input read: ", userInput)
	return userInput
}

// read csv file
// https://stackoverflow.com/questions/24999079/reading-csv-file-in-go
func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	// csvReader.Comma = ';'          // delimiter = ; instead of ,
	csvReader.FieldsPerRecord = -1 // added https://stackoverflow.com/questions/61336787/how-do-i-fix-the-wrong-number-of-fields-with-the-missing-commas-in-csv-file-in
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
