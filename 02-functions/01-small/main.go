package main

import (
	"strconv"
	"strings"
)

/**
 * Good functions are small and they do only one thing. Good and Small functions have just few lines, like at most 4 or 5 lines. The less the better.
 *
 * The first thing we need to do:
 * Keep functions small.
 * But how do we do this?
 * Breaking it into many small functions (Like hasValidRange)
**/

// Now the main function is a way simplier and easier than before. It has 5 lines.
func IsValidIp(ip string) bool {
	values := strings.Split(ip, ".")

	if len(values) != 4 {
		return false
	}

	return hasValidValues(values)
}

// This is the best scenario. It has only 1 line, does literally only one thing.
func hasValidRange(number int) bool {
	return number >= 0 && number <= 255
}

/**
 * This function still does a lot of things. It has 11 lines (counting the closing brackets)
 * We can apply the same rule here, breaking it into smaller functions. But we will cover more things on it latter.
**/
func hasValidValue(value string) bool {
	if value == "0" || string(value[0]) != "0" {
		convertedValue, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if !hasValidRange(convertedValue) {
			return false
		}

		return true
	}

	return false
}

// We've extracted the "for" statement into a small function and its verifications statements into another one. It has 6 lines (counting the closing brackets)
func hasValidValues(values []string) bool {
	for _, value := range values {
		if !hasValidValue(value) {
			return false
		}
	}
	return true
}
