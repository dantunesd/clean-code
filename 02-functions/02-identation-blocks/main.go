package main

import (
	"strconv"
	"strings"
)

/**
 * The first thing we need to do:
 * Decrease the identation level.
 * The less the better.
 * But how do we do this?
 * Extract ifs, fors, etc inside another if, for, etc to a new function
**/

// This function is easy to understand because it has only one identation level (one "if")
func IsValidIp(ip string) bool {
	values := strings.Split(ip, ".")

	if len(values) != 4 {
		return false
	}

	return hasValidValues(values)
}

// This function has the best readability possible, it doesn't have any statement like "if" (yes, it's a simplified if), "for", "while", "switch", etc.
func hasValidRange(number int) bool {
	return number >= 0 && number <= 255
}

/**
 * Previously this function had a difficult readability.
 * It had two identation levels, a confused logic, and so on...
 * To improve it, we apply the identation block rule. Minimize the identation level
 * Also we've implemented the return early approach using the hasValidStart method
 * Now this function has only one identation level
**/
func hasValidValue(value string) bool {
	if !hasValidStart(value) {
		return false
	}

	valueAsNumber, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return hasValidRange(valueAsNumber)
}

// This function still have two identation levels, but it still simple
func hasValidValues(values []string) bool {
	for _, value := range values {
		if !hasValidValue(value) {
			return false
		}
	}

	return true
}

// We've extracted it from inside hasValidValue method
func hasValidStart(value string) bool {
	return value == "0" || string(value[0]) != "0"
}
