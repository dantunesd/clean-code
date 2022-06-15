
# Clean Code

A brief description of what this project does and who it's for

## Summary

- [Naming](#naming)
    - [Meaningful names](#meaningful-names)
    - [Avoid wrong information](#avoid-wrong-information)
    - [Meaningful distinction](#meaningful-distinction)
    - [Searchable names](#searchable-names)
    - [Avoid mental mapping](#avoid-mental-mapping)
    - [Class names](#class-names)
    - [Method names](#method-names)
    - [One word per context](#one-word-per-context)
    - [Use domain solution and problem names](#use-domain-solution-and-problem-names)
    - [Meaningful context](#meaningful-context)
    - [Avoid unnecessary contex](#avoid-unnecessary-context)

## Naming

### Meaningful names

#### Bad

```golang
// What dooes "Sq" do? What "s" means? What about "t" and "n"?
func Sq(s []int) int {
	var t int
	for _, n := range s {
		t += n * n
	}
	return t
}
```

#### Good

```golang
// A function which receives some numbers and does a square sum operation
func SumSquare(numbers []int) int {
	var totalSum int
	for _, number := range numbers {
		totalSum += number * number
	}
	return totalSum
}
```

---

### Avoid wrong information

#### Bad

```golang
// A simple struct
type Car struct {}

// It's not really an array
var carArray []Car

// It does not store the time in yyyymmdd format
var yyyymmdd time.Now()
```

#### Good

```golang
// A simple struct
type Car struct {}

// Just Cars
var cars []Car

// It does store the current time
var currentTime time.Now()
```

---

### Meaningful distinction

#### Bad

```golang
// A simple struct
type Car struct {}

// Good name
var car Car

// But, what's the difference between car2 and car variables?
var car2 Car

// What's the real difference between the functions below?
func GetCar() Car {
    // Returning a Car
}

func GetCarsInfo() Car {
    // Returning Cars info (?)
}

func GetCarData() string {
    // Returning a Car data (?)
}
```

#### Good

```golang
// A simple struct
type Car struct {}

// Meaningful distinction. It's my car
var myCar Car

// Meaningful distinction. It's my wife's car
var myWifesCar Car

// Difference between them is clear now
func GetCar() Car {
    // returning a Car
}

func GetCars() []Car {
    // returning a list of cars
}

func GetCarName() string {
    // returning car name
}
```

---

### Searchable names

#### Bad

```golang
// Imagine searching for just "a" on your IDE
var a = 22
```

#### Good

```golang
// It will be easier to find all ocurrencies
var averageMonthlyWorkingDays = 22
```

---

### Interface implementations

```
to be implemented
```

---

### Avoid mental mapping

#### Bad

```golang
// It's not common to use letters other than "i", "j", "k" in a for statement
for ç := 0; ç < 2; ç++ {
    fmt.Println(ç)
}

// An example of a domain object
type Vehicle struct {}

// Maybe it's not common word to represents its meaning when talking to your teammates, stakeholders, etc.
var motorizedMachine Vehicle
var vehicleThatTakesPeople Vehicle

```

#### Good

```golang
// Prefer conventions
for i := 0; i < 2; i++ {
    fmt.Println(i)
}

// An example of a domain object
type Vehicle struct {}

// Use ubiquitous language.
var car Vehicle
var bus Vehicle
var taxi Vehicle
```

---

### Class names

#### Bad

```golang
// Avoid generic names such as Data, Info, Manager, Process, and so on.
type CustomerData struct {
	Name string
}

type CarInfo struct {
	Manufacturer string
}
```

#### Good

```golang
// Prefer using just nouns for naming 
type Customer struct {
	Name string
}

type Car struct {
	Manufacturer string
}
```

---

### Method names

#### Bad

```golang
// A simple struct
type Employee struct {
	name        string
	nationality string
}

// Methods should contain verbs on its name. Puts "get" when getting some data
func (e *Employee) Name() string {
	return e.name
}

// If you are asking if something is something else, puts "is" as prefix
func (e *Employee) Brazilian() bool {
	return e.nationality == "BRASIL"
}
```

#### Good

```golang
// A simple struct
type Employee struct {
	name        string
	nationality string
}

// Methods are actions, actions are verbs. Methods should contain a verb on its name.
func (e *Employee) GetName() string {
	return e.name
}

// Is the employee brazilian?
func (e *Employee) IsBrazilian() bool {
	return e.nationality == "BRASIL"
}
```

---

### One word per context

#### Bad

```golang
// A simple struct
type Customer struct {
	Name string
}

// Another simple struct
type Car struct {
	Name string
}

// This function returns the object name
func (c *Customer) GetName() string {
	return c.Name
}

// This function also returns the object name. But, why naming it different?
func (c *Car) RetrieveName() string {
	return c.Name
}
```

#### Good

```golang
// A simple struct
type Customer struct {
	Name string
}

// Another simple struct
type Car struct {
	Name string
}

// This function returns the object name
func (c *Customer) GetName() string {
	return c.Name
}

// Using the same name for the same concept
func (c *Car) GetName() string {
	return c.Name
}
```

---

### Use domain solution and problem names


#### Good

```golang
// There's no problem in using "domain solution" names. Others programmers probably know what it means and it may facilitate to them (Ex: Design Patterns)
type BusStrategy struct{}
type CarFactory struct{}
type ActiveAccountState struct{}

// You can just use "domain problem" names if there aren't any "domain solution" name. At least some domain specialist may know more about it and you can ask him.
type Account struct{}
```

---

### Meaningful context

#### Bad

```golang

// it could represent whatever number.
var number int

// it could represent a state name? like "active" or "inactive".
var state string

// now i'm getting the context... an address
var street string

// Only after reading these 3 variables you may notice that them is part of a bigger context: Address. Now imagine if they were all spread out in some code.
```

#### Good

```golang
// Adding prefixes could help to give them a context
var addressNumber int
var addressState string
var adressStreet string

// Or even better... Group them together into a Class or Struct when it's possible.
type Address struct {
	Street string
	Number int
	State  string
}
```

---

### Avoid unnecessary context

#### Bad

```golang
package whatever

// Adding unecessary context
type whateverPkgAddress struct {
	AddressStreet string
	whateverNumber int
	AdressState  string
}
```

#### Good

```golang
package whatever

// Simple, contextualized.
type Address struct {
	Street string
	Number int
	State  string
}
```

---

