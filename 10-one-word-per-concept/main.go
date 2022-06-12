package main

type Customer struct {
	Name string
}

type Car struct {
	Name string
}

type Animal struct {
	Name string
}

// All these methods are doing the same thing, but they are named different. Avoid it. Try to use the same word for the same concept
func (c *Customer) GetName() string {
	return c.Name
}

// The method name could be GetName (Or all the others could be Retrieve...)
func (c *Car) RetrieveName() string {
	return c.Name
}

// The method name could be GetName (Or all the others could be Fetch...)
func (a *Animal) FetchName() string {
	return a.Name
}
