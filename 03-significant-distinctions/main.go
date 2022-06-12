package main

import "fmt"

type Product struct{}

func main() {
	// Good: it represents a Product
	var product Product

	// Bad: What info? Vague word
	var productInfo Product

	// Bad: What is the difference between Data and Info? and with product? Vague word
	var productData Product

	// Bad: Yes, it's a variable
	var varProduct Product

	// Bad: Yes, it's a variable and a struct
	var productStruct Product

	fmt.Println(product, productInfo, productData, varProduct, productStruct)
}
