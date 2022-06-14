package main

import (
	"strconv"
	"strings"
)

/** The first thing we need to do:
 * Keep functions small.
 * The smaller the better.
 * But how do we do this?
 * Breaking it into many small functions (Like hasValidRange)
**/

// Now the main function is a way simplier and easier than before
func IsValidIp(ip string) bool {
	values := strings.Split(ip, ".")

	if len(values) != 4 {
		return false
	}

	return hasValidValues(values)
}

// Simple methods like this one is basicaly what it says it does
func hasValidRange(number int) bool {
	return number >= 0 && number <= 255
}

// This function still does a lot of things. We can apply the same rule here, break it into small functions. But we will cover it later.
func hasValidValue(value string) bool {
	if value == "0" || string(value[0]) != "0" {
		valueAsNumber, err := strconv.Atoi(value)
		if err != nil {
			return false
		}

		if !hasValidRange(valueAsNumber) {
			return false
		}

		return true
	}

	return false
}

// We extracted the for statement into a small function, and it's verification into another one
func hasValidValues(values []string) bool {
	for _, value := range values {
		if !hasValidValue(value) {
			return false
		}
	}
	return true
}
