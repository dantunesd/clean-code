package main

// After you see these 3 variables, you may notice that them is part of a bigger context, like "Address".
// Now imagine if they were all spread out. It would be more difficult to get this context
var Number int
var State string
var Street string

// Give them context adding prefixes if names are not that self explanatory
var AddressNumber int
var AddressState string

// Or even better... Group them together into a Class or Struct when it's possible.
type Address struct {
	Street string
	Number int
	State  string
}
