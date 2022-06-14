package main

import "fmt"

/**
 * Functions (mainly monadics) should form a well pair of "verb" and "noun". Like the example below:
 * It prints the name
 */
func Print(name string) {
	fmt.Printf("Name: %s", name)
}
