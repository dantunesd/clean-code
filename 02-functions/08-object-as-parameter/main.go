package main

/**
 * If you find functions which have 3 or more parameters, it's a sign you can group its parameters into a representative class, giving it also a significant context
 */

// Receiving now two params. We've grouped x and y into a representative struct/class, giving it a context and simplifying this function
func MakeCircle(center Point, radius float64) string {
	return "creating a circle with the two params now"
}

type Point struct {
	X float64
	Y float64
}
