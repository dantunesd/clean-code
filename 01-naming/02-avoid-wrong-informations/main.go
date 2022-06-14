package main

type Car struct {
	Name string
}

// Bad: It's not an "array"
var CarArray []Car

// Good: But "list" may cause confusion yet
var ListOfcars []Car

// Better: The type doesn't matter at all
var Cars []Car
