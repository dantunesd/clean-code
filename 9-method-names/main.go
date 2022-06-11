package main

type Employee struct {
	name        string
	nationality string
}

// Bad: Methods should contain verbs and follow some conventions. If you are retrieving any data, puts the prefix "get"
func (e *Employee) Name() string {
	return e.name
}

// Good: Methods are actions, actions are verbs. Methods should contain a verb on its name. Prefixing "get" is an convention when retrieving any data
func (e *Employee) GetName() string {
	return e.name
}

// Bad: If you are validating something and your method returns a bool, puts the prefix "is"
func (e *Employee) Brazilian() bool {
	return e.nationality == "BRASIL"
}

// Good: verifies If Employee Is Brazilian
func (e *Employee) IsBrazilian() bool {
	return e.nationality == "BRASIL"
}
