package main

import (
	"bufio" // user terminal input
	"fmt"

	"ginalamp-mark-tracker/oop"
	"os"      // user input
	"strconv" // type conversion
	"strings"
)

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

	// basic assignment initialisation
	a2 := oop.Assignment{AssignmentName: "Accounting", Mark: 77, Weight: 50}
	fmt.Println(a2)
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
