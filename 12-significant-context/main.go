package main

import "fmt"

func main() {
	// After you see these 3 variables, you may notice that them is part of a bigger context, like "Adress".
	// Now imagine if they were all spread out. It would be more difficult to get this context
	var number int
	var state string
	var street string

	// Give them context adding prefixes if names are not that self explanatory
	var addressNumber int
	var addressState string

	fmt.Println(street, number, state, addressNumber, addressState)
}

// Or even better... Group them together into a Class or Struct.
type Address struct {
	Street string
	Number int
	State  string
}
