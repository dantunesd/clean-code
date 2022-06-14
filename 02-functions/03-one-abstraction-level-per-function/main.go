package main

import (
	"strconv"
	"strings"
)

/**
 * The first thing we need to do:
 * Leave the function with only one abstraction level.
 * The less the better.
 * But how do we do this?
 * Extract low details into smaller functions.
**/

/**
 * This function used to have different abstraction levels:
 * It used to split the ip by its dots into a list of values, it used to validate the values and call hasValidValues to validate each value
 * Now it is delegating all those operations to other functions. Those operations are abstracted into new functions
**/
func IsValidIp(ip string) bool {
	values := splitIp(ip)

	if !isAnOctet(values) {
		return false
	}

	return hasValidValues(values)
}

// The low details are encapsulated here, others functions don't need to know how to split an IP
func splitIp(ip string) []string {
	return strings.Split(ip, ".")
}

// Same here. Others functions don't need to know the details of an octet
func isAnOctet(values []string) bool {
	return len(values) == 4
}

func hasValidRange(number int) bool {
	return number >= 0 && number <= 255
}

func hasValidValue(value string) bool {
	if !hasValidStart(value) {
		return false
	}

	valueAsNumber, err := convertToNumber(value)
	if err != nil {
		return false
	}

	return hasValidRange(valueAsNumber)
}

func convertToNumber(value string) (int, error) {
	return strconv.Atoi(value)
}

func hasValidValues(values []string) bool {
	for _, value := range values {
		if !hasValidValue(value) {
			return false
		}
	}

	return true
}

func hasValidStart(value string) bool {
	return value == "0" || string(value[0]) != "0"
}
