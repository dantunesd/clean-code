package main

import (
	"strconv"
	"strings"
)

/** Even if this function has less than 25 lines, an appropriated function name, some significant variable names, it still difficult to read/understand this function.
 * It's difficult because it's long, it does a lot of things, has a lot statements chained, has different levels of abstraction, has a confuse reading flow, and so on...
 * Bad: for all the things related above
**/

func IsValidIp(ip string) bool {
	values := strings.Split(ip, ".")

	if len(values) != 4 {
		return false
	}

	for _, value := range values {
		if value == "0" || string(value[0]) != "0" {
			valueAsNumber, err := strconv.Atoi(value)
			if err != nil {
				return false
			}

			if valueAsNumber >= 0 && valueAsNumber <= 255 {
				continue
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
