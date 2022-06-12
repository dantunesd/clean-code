package main

type Product struct{}

// Good: It represents a Product
var product Product

// Bad: What is the difference between product and product2?
var product2 Product

// Bad: What info? Vague word
var productInfo Product

// Bad: What is the difference between Data and Info? and with product? Vague word
var productData Product

// Bad: Yes, it's a variable
var varProduct Product

// Bad: Yes, it's a variable and a struct
var productStruct Product
