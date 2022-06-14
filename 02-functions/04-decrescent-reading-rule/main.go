package main

import (
	"strconv"
	"strings"
)

func IsValidIp(ip string) bool {
	values := splitIp(ip)
	return validateValues(values)
}

func splitIp(ip string) []string {
	return strings.Split(ip, ".")
}

func validateValues(values []string) bool {
	if !isAnOctet(values) {
		return false
	}

	return hasValidValues(values)
}

func isAnOctet(values []string) bool {
	return len(values) == 4
}

func hasValidValues(values []string) bool {
	for _, value := range values {
		if !hasValidValue(value) {
			return false
		}
	}
	return true
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

func hasValidStart(value string) bool {
	return value == "0" || string(value[0]) != "0"
}

func convertToNumber(value string) (int, error) {
	return strconv.Atoi(value)
}

func hasValidRange(number int) bool {
	return number >= 0 && number <= 255
}
