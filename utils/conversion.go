package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// **************************************************************************************
// *** helper functions
// **************************************************************************************

// convert "50%" to 50.0
func PercentageToFloat(s string) float64 {
	r := strings.Replace(s, "%", "", -1)
	num, err := strconv.ParseFloat(r, Float64Type)
	if err != nil {
		return Error
	}
	return num
}

// string input to float
func StringToFloat(s string) float64 {
	i, err := strconv.ParseFloat(s, Float64Type)
	if err != nil {
		fmt.Println("Invalid input. Please type in a number (integer)")
		return (Error)
	}
	return (i)
}

// convert float64 to string with 0 decimals
func FloatToString(f float64) string {
	return fmt.Sprintf("%.0f", f)
}
