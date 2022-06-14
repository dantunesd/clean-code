package main

import (
	"strconv"
	"strings"
)

/**
 * Good functions must do only one thing. But, how to know if it's doing one thing? It must have only one abstraction level.
 * Low and high details together it's a sign the it's doing more than one thing.
 *
 * The first thing we need to do:
 * Leave the function with only one abstraction level.
 * But how do we do this?
 * Extract low details into smaller functions.
**/

/**
 * This function previously had different abstraction levels:
 * It had a statement to split the ip by its dots into a list of values, it used to validate values length and call hasValidValues to validate each value
 * Now it is delegating all those operations to smaller functions. They are encapsulating those details.
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

// We've applied this rule here before.
func hasValidRange(number int) bool {
	return number >= 0 && number <= 255
}

// We've extracted strconv.Atoi(value) to "convertToNumber" function
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

// strconv.Atoi(value) could have the same abstraction level as the others functions in "hasValidValue".
// But, I've decided to move it to apply "pronounceable-names" rule
func convertToNumber(value string) (int, error) {
	return strconv.Atoi(value)
}

// We've applied this rule here before.
func hasValidValues(values []string) bool {
	for _, value := range values {
		if !hasValidValue(value) {
			return false
		}
	}

	return true
}

// We've applied this rule here before.
func hasValidStart(value string) bool {
	return value == "0" || string(value[0]) != "0"
}
