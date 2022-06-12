package main

// If class/struct name already gives you a well defined context, avoid adding unnecessary context.
type Address struct {

	// Bad: We are already in Address Struct, there is no need to repeat "Address" again
	AddressStreet string

	// Good: We already know this number and state is part of an Address
	Number int
	State  string
}

// The same rule could be applied to modules, class names, folder names, etc. We are already in a folder named 'avoid-unnecessary-context' folder.
type AvoidUnnecessaryContextAddress struct{}
