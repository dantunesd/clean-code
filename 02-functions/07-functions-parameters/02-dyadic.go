package main

/**
 * Dyadic functions are functions that receive two parameters. It's relatively easy to understand its parameters
 * But even them sometimes may be difficult to understand. Probably at least one time you have changed the order of the AssertEquals parameters when testing.
 * If you can simplify it to a monadic function, do it.
 */

// A super simplified example of an assert
func AssertEquals(expected, actual string) bool {
	return expected == actual
}

// An example how to avoid confusion with parameters. There are way more code implemented now, but we won't confuse its parameters anymore
type Assert struct {
	expected string
	actual   string
}

// Receives the expected value, monadic function (Better)
func (a *Assert) SetExpected(expected string) {
	a.expected = expected
}

// Receives the actual value, monadic function (Better)
func (a *Assert) SetEquals(actual string) {
	a.actual = actual
}

// Executes the validation, 0 parameters (Even better)
func (a *Assert) IsEquals() bool {
	return a.expected == a.actual
}
