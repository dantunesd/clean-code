package main

/**
 * Function can receive as many parameters as you like.
 * The problem with this is that your function will become harder to understand, probably because it does a lot of things with these parameters
 * It's difficult to test all the combinations, and so on.
 *
 * So, the less parametes your functions receive the better. 0 is the best. 1 (monadic) is good, 2 (dyadic) is ok, 3... well, better no.
 *
 * There are two good reasons to create a function that receives only one parameter (monadic function):
 * Or you are asking something about it or you are transforming it into something else .
 */

// Good: A commom monadic function, a function that only validates its parameter and return a boolean
func IsSomething(param string) bool {
	// verify if param is something
	return true
}

// Good: A commom monadic function, a function that uses its parameter to return something else
func OpenFile(filename string) string {
	// opening the file
	return "file-opened-example.txt"
}

// Good: A less commom monadic function, a function that receives its parameter to do something else but doesn't return anything. Use it cautiously. It should be obvious what it will do.
func SendEmail(email string) {
	// sending the email
}

type MyObject struct {
	SomeField string
}

// Try to avoid functions that use its parameter as an "output parameters"
func ModifyMyObjectField(myObject *MyObject) {
	myObject.SomeField = "different value"
}

// Preffer returning the parameter even if its the same
func ModifyMyObjectField2(myObject *MyObject) *MyObject {
	myObject.SomeField = "different value"
	return myObject
}
