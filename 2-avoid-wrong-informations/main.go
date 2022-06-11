package main

import "fmt"

type Car struct {
	Name string
}

func main() {
	// Bad: It's not an "array"
	var carArray []Car

	// Good: But "list" can canfuse yet
	var listOfcars []Car

	// better: Type doesn't matter at all
	var cars []Car

	fmt.Println(carArray, listOfcars, cars)
}
