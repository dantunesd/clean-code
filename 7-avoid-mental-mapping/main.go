package main

import "fmt"

func main() {
	useConventions()
	useUbiquitousLanguage()
}

func useConventions() {
	// Bad: It's a "convention" to use one letter in "for" statements to represent "index" (i) or "value" (v). But "ç" is definitely not common.
	for ç := 0; ç < 2; ç++ {
		fmt.Println(ç)
	}

	// Good: Use conventions. Even though "i" isn't a significant name, a pronounceable word, but "i" is more commonly used than "ç". You don't have to translate it mentally to understand its meaning
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}
}

func useUbiquitousLanguage() {
	// An example of a domain object
	type Vehicle struct {
		Type           string
		EngineCapacity int
	}

	// Bad: Maybe it's not common word to represents its meaning when talking with your teammates, stakeholders, etc.
	var motorMachine Vehicle

	// Good: Maybe it's better because it expresses a known name for the company, team, etc. You don't have to translate it mentally to understand its meaning
	var car Vehicle
	var bus Vehicle

	fmt.Println(motorMachine, car, bus)
}
