package utils

import (
	"log"
	"os"
)

// make directory if not exist
// https://gist.github.com/ivanzoid/5040166bb3f0c82575b52c2ca5f5a60c
func MakeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}

// check if user wan't to exit the program
func UserExit(in string) bool {
	if in == "exit" || in == "quit" {
		return true
	}
	return false
}

// check if error != nil
func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// concatenate strings in csv format
func Concat2StringsCsvFormat(string1, string2 string) string {
	return string1 + ";" + string2
}
