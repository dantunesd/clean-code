
# Clean Code

A brief description of what this project does and who it's for

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


### Searchable names

#### Bad

```golang
// Imagine searching for just "a" on your IDE
var a = 22
```

#### Good

```golang
// It will be easier to find all ocurrencies
var averageMonthlyWorkDays = 22
```


### Interface implementations

```
to be implemented
```


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
// Preffer conventions
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