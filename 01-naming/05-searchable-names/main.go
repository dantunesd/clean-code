package main

// Bad: What does "a" mean? Imagine you need to change all ocurrences of this "a", but there are others "a" with different meaning
var a = 20

// Bad: The same goes for "g". Imagine this function being used in many places, and you have to find it or change its name
func G(n int) int {
	// Bad: The same goes for "n".
	if n == 0 {
		return 0
	}

	return n / a
}

// Good: Significant, pronounceable and easy to find all occurrences
var averageMonthlyWorkingDays = 20

// Good: Significant, pronounceable and easy to find all occurrences
func GetAverageTasksDeliveredPerDay(tasks int) int {
	// Good: Significant, pronounceable and easy to find all occurrences
	if tasks == 0 {
		return 0
	}

	return tasks / averageMonthlyWorkingDays
}
