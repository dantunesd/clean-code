package main

import (
	"strconv"
	"strings"
)

/**
 * Good functions is generally small and broken into many smaller functions. But also it must have a good reading flow. Like we were reading a book.
 * Decrescent reading rule is a practice to put functions right below where it's been called
 *
 * The first thing we need to do:
 * If you are calling a function, place that function declaration right below.
**/

/**
 * This function calls splitIp, isAnOctet and hasValidValues functions.
 * Note that they are now right bellow this function, simplifying the navigation and readbility.
 *
**/
func IsValidIp(ip string) bool {
	values := splitIp(ip)

	if !isAnOctet(values) {
		return false
	}

	return hasValidValues(values)
}

func splitIp(ip string) []string {
	return strings.Split(ip, ".")
}

func isAnOctet(values []string) bool {
	return len(values) == 4
}

// The same rule was applied here. this function calls hasValidValue and it is right below this one.
func hasValidValues(values []string) bool {
	for _, value := range values {
		if !hasValidValue(value) {
			return false
		}
	}
	return true
}

// The same rule was also applied here.
func hasValidValue(value string) bool {
	if !hasValidStart(value) {
		return false
	}

	convertedValue, err := convertToNumber(value)
	if err != nil {
		return false
	}

	return hasValidRange(convertedValue)
}

func hasValidStart(value string) bool {
	return value == "0" || string(value[0]) != "0"
}

func convertToNumber(value string) (int, error) {
	return strconv.Atoi(value)
}

func hasValidRange(value int) bool {
	return value >= 0 && value <= 255
}
