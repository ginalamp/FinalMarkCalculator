package main

import (
	"bufio" // user terminal input
	"fmt"

	// "oop"
	"os"      // user input
	"strconv" // type conversion
	"strings"
)

// type Degree struct {
// 	DegreeName string
// 	Modules    []Module
// }

// type Module struct {
// 	ModuleName  string
// 	Weight      int // weight of the module in the total degree
// 	Mark        int
// 	Assignments []Assignment
// }

// type Assignment struct {
// 	AssignmentName string
// 	Mark           int
// 	Weight         int // weight of the assignment in the module
// }

func readInput(userPrompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(userPrompt)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput) // remove whitespace

	fmt.Println("input read - ", userInput)
	return userInput
}

func main() {
	// m := degree.Module{}
	// fmt.Println(m)

	numModules := 0
	for numModules < 1 {
		numModules = stringToInt(readInput("How many modules do you have?"))
	}
	fmt.Println(numModules)

	// degree.Module

	e1 := Assignment{AssignmentName: "name", Mark: 69, Weight: 50}
	fmt.Println(e1)
}

// string to int
func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Invalid input. Please type in the number of modules that you have (as an integer)")
		// handle error
		fmt.Println(err)
		return (0)
	}
	return (i)
}
