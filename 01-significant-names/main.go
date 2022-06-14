package main

// Bad: What does sq mean? And about "s"?
func Sq(s []int) int {
	// What is T?
	var t int

	// I don't know what is "s", neither what is "n"
	for _, n := range s {
		t += n * n
	}

	// Well, let's just return it :(
	return t
}

// Good: Significant function and parameter name
func SumSquare(numbers []int) int {
	var totalSum int

	// it's easier now to comprehend that it's doing a square sum operation with the given numbers
	for _, number := range numbers {
		totalSum += number * number
	}

	// returning the total
	return totalSum
}

// Bad: Get what? What "a" and "b" represent?
func Get(a, b int) int {
	r := a * b
	return r
}

// Good: A function that multiplies the firstNumber and secondNumber and returns the result
func Multiply(firstNumber, secondNumber int) int {
	result := firstNumber * secondNumber
	return result
}
