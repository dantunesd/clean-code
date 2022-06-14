package main

// Bad: Data is too generic, it could be simplified to just Customer. Same goes for Manager, Processor, Common, Utils, etc.
type CustomerData struct {
	Name string
}

// Bad: Info is too generic, it could be simplified to just Car. Same goes for Manager, Processor, Common, Utils, etc.
type CarInfo struct {
	Manufacturer string
}

// Good: a noun without vague or confuse words that "means nothing" at all
type Customer struct {
	Name string
}

// Good: a noun without vague or confuse words that "means nothing" at all
type Car struct {
	Manufacturer string
}
