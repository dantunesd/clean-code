package main

import (
	"strconv"
	"strings"
)

/**
 * To have good functions, in addition to being small, they must have few identations level. At most 2 levels.
 * The less "if" inside "if", "for" inside another "for", the simpler your code will looks like
 *
 * The first thing we need to do:
 * Decrease the identation level.
 * But how do we do this?
 * Extract "if", "for", etc, from others "if", "for" and put them into a new function
**/

// This function is easy to understand because it has only one identation level (one "if")
func IsValidIp(ip string) bool {
	values := strings.Split(ip, ".")

	if len(values) != 4 {
		return false
	}

	return hasValidValues(values)
}

// This function has the best identation level possible, it doesn't have any statement like "if" (yes, it's a simplified if), "for", "while", "switch", etc.
func hasValidRange(number int) bool {
	return number >= 0 && number <= 255
}

/**
 * Previously this function had a difficult readability beucase its identation level.
 * It had two identation levels, a confused logic, and so on...
 * To improve it, we've extracted some code to minimize the identation level.
 * Also we've implemented the return early approach.
 * Now this function has only one identation level.
**/
func hasValidValue(value string) bool {
	if !hasValidStart(value) {
		return false
	}

	convertedValue, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return hasValidRange(convertedValue)
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
