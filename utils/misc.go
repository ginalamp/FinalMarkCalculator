package utils

// miscellaneous helper functions

import (
	"log"
	"os"
	"strings"
)

// make directory if not exist
// https://gist.github.com/ivanzoid/5040166bb3f0c82575b52c2ca5f5a60c
func MakeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}

// check if user wants to exit the program
func UserExit(in string) bool {
	if strings.ToLower(in) == "exit" || strings.ToLower(in) == "quit" {
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
