package main

import "fmt"

func main() {

	/**
	 * Sometimes you want to use a list to pass as a parameter, and it's ok.
	 * Functions like Printf can receive a variable number of parameters, and still be treated as a monadic, dyadic or triad function.
	 */
	fmt.Printf("%s %s", "name", "surname")
}
