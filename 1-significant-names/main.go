package main

import "fmt"

func main() {
	fmt.Println(sq([]int{1, 2}))
	fmt.Println(sumSquare([]int{1, 2}))

	fmt.Println(get(2, 2))
	fmt.Println(multiply(2, 2))
}

// Bad: What does sq mean? and "s"?
func sq(s []int) int {
	// What is T?
	var t int

	// I don't know what is "s", and now I don't know what is "n"
	for _, n := range s {
		t += n * n
	}

	// Well, let's just return it :(
	return t
}

// Good: Significant function and parameter name
func sumSquare(numbers []int) int {
	var totalSum int

	// it's easier to comprehend now that we are doing a square sum operation with the given numbers
	for _, number := range numbers {
		totalSum += number * number
	}

	// returning the total
	return totalSum
}

// Bad: Get what? What "a" and "b" represents?
func get(a, b int) int {
	r := a * b
	return r
}

// Good: A function that multiplies the firstNumber with the secondNumber and returns their result
func multiply(firstNumber, secondNumber int) int {
	result := firstNumber * secondNumber
	return result
}
