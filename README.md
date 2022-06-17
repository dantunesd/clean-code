
# Clean Code

A brief description of what this project does and who it's for

## Summary

- [Naming](#naming)
    - [Meaningful names](#meaningful-names)
    - [Avoid wrong information](#avoid-wrong-information)
    - [Meaningful distinction](#meaningful-distinction)
    - [Pronounceable names](#pronounceable-names)
    - [Searchable names](#searchable-names)
    - [Avoid mental mapping](#avoid-mental-mapping)
    - [Class names](#class-names)
    - [Method names](#method-names)
    - [One word per context](#one-word-per-context)
    - [Use domain solution and problem names](#use-domain-solution-and-problem-names)
    - [Meaningful context](#meaningful-context)
    - [Avoid unnecessary contex](#avoid-unnecessary-context)
- [Functions](#functions)
    - [Small](#small)
    - [Identation blocks](#identation-blocks)
    - [One abstraction level per function](#one-abstraction-level-per-function)
    - [Decrescent reading](#decrescent-reading)
    - [Function parameters](#function-parameters)
    - [Verbs and key-words](#verbs-and-key-words)
    - [Avoid collateral effect](#avoid-collateral-effect)


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
// A function which receives some numbers, does squared operations and sum the results
func SumNumbersSquared(numbers []int) int {
    var result int
    for _, number := range numbers {
        result += number * number
    }
    return result
}
```

---

### Avoid wrong information

#### Bad

```golang
// It does not store the time in yyyymmdd format
var yyyymmddTime = time.Now()
```

#### Good

```golang
// It stores the current time, whatever the format
var currentTime = time.Now()
```

---

### Meaningful distinction

#### Bad

```golang
// What's the difference between them?
var time1 = time.Now()
var time2 = time.Now().AddDate(0, 0, 1)
var timeObject = time.Now(AddDate(1, 0, 0))
```

#### Good

```golang
// It's clear now
var now = time.Now()
var tomorrow = time.Now().AddDate(0, 0, 1)
var nextYear = time.Now(AddDate(1, 0, 0))
```

---

### Pronounceable names

#### Bad

```golang
// It's not a word, it's difficult to pronounce
var crtTm = time.Now()
```

#### Good

```golang
// Words, they are easy to pronounce
var currentTime = time.Now()
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

Prefer always using conventions and ubiquitous language(DDD). It will facilitate everybody who is reading your code, avoiding them having to translate your code to what it was supposed to be while reading. 

#### Bad

```golang
for ç := 0; ç < count; ç++ {
    fmt.Println(ç)
}

var motorizedMachine Vehicle
var vehicleThatTakesPeople Vehicle
```

#### Good

```golang
for i := 0; i < count; i++ {
    fmt.Println(i)
}

var car Vehicle
var taxi Vehicle
```

---

### Class names

#### Bad

```golang
// Avoid generic names such as Data, Info, Manager, Process, and so on.
type CarInfo struct {
    Name string
}

type CustomerData struct {
    Name string
}
```

#### Good

```golang
// Prefer using just nouns for naming 
type Car struct {
    Name string
}

type Customer struct {
    Name string
}
```

---

### Method names

#### Bad

```golang
type Employee struct {
    name        string
    nationality string
}

// Imagine reading this method being called... Employee, "Name".
func (e *Employee) Name() string {
    return e.name
}

// What "Brazilian" method should do?
func (e *Employee) Brazilian() bool {
    return e.nationality == "BRASIL"
}
```

#### Good

```golang
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
// This function returns the object name
func (c *Car) GetName() string {
    return c.Name
}

// This function also returns the object name. But, why naming it different?
func (c *Customer) RetrieveName() string {
    return c.Name
}

// Same here
func (a *Animal) FetchName() string {
    return a.Name
}
```

#### Good

```golang
// This function returns the object name
func (c *Car) GetName() string {
    return c.Name
}

// Using the same name for the same concept
func (c *Customer) GetName() string {
    return c.Name
}

// Same here
func (a *Animal) GetName() string {
    return a.Name
}
```

---

### Use domain solution and problem names

There's no problem in using "domain solution" names. Other programmers probably know what it means and it may facilitate to them (Ex: Design Patterns)

You can just use "domain problem" names if there aren't any "domain solution" name. At least some domain specialist may know more about it and you can ask him.

#### Good

```golang
// 
type BusStrategy struct{}
type CarFactory struct{}
type ActiveAccountState struct{}
```

---

### Meaningful context

#### Bad

```golang

// It could represent whatever number.
var number int

// It could represent a state name? like "active" or "inactive".
var state string

// Now i'm getting the context...
var street string

// Only after reading these 3 variables you may have noticed that them is part of a bigger context: Address.
// Now imagine if they were all spread out in some code.
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

## Functions

### Small

Great functions are difficult to understand because they do a lot different things, have lots of statements chained, have different levels of abstraction, have a confused reading flow, and so on...

Small functions should have preferably at most 5 lines. The less the better.

#### Bad

```golang
func SomeWeirdOperation() {
    if 1 < 10 {
        fmt.Println("1 is lower than 10")
    }

    for i := 0; i < 5; i++ {
        fmt.Printf("hello")
    }

    currentTime := time.Now()
    fmt.Printf("Now: %s\n", currentTime.String())

    // Doing another thing, now we have 9 lines of code

    // Doing one more, now 10 lines of code

    // Still going, now 11 lines of code
}
```

#### Good

```golang
// Now this function delagates everything and it has only 3 lines of code
func SomeWeirdOperation() {
    printIfNumberOneIsLowerThanTen()
    printHelloFiveTimes()
    printCurrentTime()
}

// 3 lines of code (counting the brackets)
func printIfNumberOneIsLowerThanTen() {
    if 1 < 10 {
        fmt.Println("1 is lowen than 10")
    }
}

// 3 lines of code
func printHelloFiveTimes() {
    for i := 0; i < 5; i++ {
        fmt.Printf("hello")
    }
}

// 2 lines of code
func printCurrentTime() {
    currentTime := time.Now()
    fmt.Printf("Now: %s\n", currentTime.String())
}
```

---

### Identation blocks

To have good functions, in addition of being small, they must have few identation levels. At most 2 levels.

The less "if" inside another "if", "for" inside another "for", the simpler your code will looks like

#### Bad

```golang
func AnotherWeirdOperation() {
    // Only one for? no problem...
    for i := 0; i <= 10; i++ {

        // Ok, one if inside a for
        if i == 0 {
            fmt.Println("starting")
        }

        // Another
        if i == 5 {
            fmt.Println("we've arrived at half")
        }

        // Another
        if i == 10 {
            fmt.Println("let's just loop two more times")

            // Another for? It's getting confused now
            for j := 0; j < 2; j++ {
                fmt.Println("last two loops")
            }

            fmt.Println("bye")
        }
    }
}

```

#### Good

```golang
// Just one for with one method call
func AnotherWeirdOperation() {
    for i := 0; i <= 10; i++ {
        PrintMessages(i)
    }
}

// Just calling methods
func PrintMessages(i int) {
    PrintStartMessage(i)
    PrintMiddleMessage(i)
    PrintFinalMessage(i)
}

// Just one if
func PrintStartMessage(i int) {
    if i == 0 {
        fmt.Println("starting")
    }
}

// Just one if
func PrintMiddleMessage(i int) {
    if i == 5 {
        fmt.Println("we've arrived at half")
    }
}

// Just one if and a method call
func PrintFinalMessage(i int) {
    if i == 10 {
        PrintTwoMoreLoopsMessage()
        fmt.Println("bye")
    }
}

// Just one for
func PrintTwoMoreLoopsMessage() {
    fmt.Println("let's just loop two more times")
    for i := 0; i < 2; i++ {
        fmt.Println("last two loops")
    }
}
```

---

### One abstraction level per function

 Good functions do only one thing. But, how to know if it's really doing only one thing? It must have only one abstraction level.
 Low and high details together it's a good sign that it's doing more than one thing.

#### Bad

```golang
// Many different abstraction levels here. It knows details of how to validate a customer and how to log, but it delegates saving operation.
func CreateCustomer(name, email string) error {
    if email != "just.an.example@gmail.com" {
        return errors.New("invalid email")
    }

    if err := saveCustomerOnDatabase(name, email); err != nil {
        return err
    }

    log.Printf("Date: %s Message: %s", time.Now(), "Customer created")

    return nil
}

// Imagine this function saves on database
func saveCustomerOnDatabase(name, email string) error {
    return nil
}
```

#### Good

```golang
// Now this function delegates all the lowest level details to other functions. They are now abstracted.
func CreateCustomer(name, email string) error {
    if err := validateCustomer(name, email); err != nil {
        return err
    }

    if err := saveCustomerOnDatabase(name, email); err != nil {
        return err
    }

    logCustomerCreated()

    return nil
}

// Still delegating some lower level details
func validateCustomer(name, email string) error {
    if !isValidEmail(email) {
        return errors.New("invalid email")
    }

    return nil
}

// This function contains the lowest detail level
func isValidEmail(email string) bool {
    return email == "just.an.example@gmail.com"
}

// Same here
func saveCustomerOnDatabase(name, email string) error {
    return nil
}

// Same here
func logCustomerCreated() {
    log.Printf("Date: %s Message: %s", time.Now(), "Customer created")
}
```

---

### Decrescent reading

Good functions are generally small and broken into many smaller functions. But also it must have a good reading flow. Like if we were reading a book.
Decrescent reading rule is a practice to put functions in order right below where they are called

#### Bad

```golang
package main

import "fmt"

func ICallTheLastOne() {
    fmt.Println("I am the last")
}

func IamTheThird() {
    fmt.Println("I am the third")
}

func main() {
    IamTheFirst()
    IamTheSecond()
    IamTheThird()
    IamSpecial()
}

func IamSpecial() {
    ICallTheLastOne()
}

func IamTheSecond() {
    fmt.Println("I am the second")
}

func IamTheFirst() {
    fmt.Println("I am the first")
}
```

#### Good

```golang
package main

import "fmt"

func main() {
    IamTheFirst()
    IamTheSecond()
    IamTheThird()
    IamSpecial()
}

func IamTheFirst() {
    fmt.Println("I am the first")
}

func IamTheSecond() {
    fmt.Println("I am the second")
}

func IamTheThird() {
    fmt.Println("I am the third")
}

func IamSpecial() {
    ICallTheLastOne()
}

func ICallTheLastOne() {
    fmt.Println("I am the last")
}
```

---

### Function parameters

Functions can receive as many parameters as you like. 

The problem with this is that your function will become harder to understand, probably because it does a lot of things with these parameters. It's difficult to test all the combinations, you may change the order of them while calling it, and so on.

So, the less parameters your functions receive the better. 0 is the best. 1 (monadic) is good, 2 (dyadic) is ok, 3 or more... well, better no.

Tip: If you find functions which have 3 or more parameters, it's a good sign that they belong to a context. So you can group them in a representative class/struct.

#### Bad

```golang
func CreateCustomer(name, email, phone, birthdate string) {
    // creating customer
}
```

#### Good

```golang
type Customer struct {
    name      string
    email     string
    phone     string
    birthdate string
}

func CreateCustomer(customer *Customer) {
    // creating customer
}
```

---

### Verbs and key-words

Functions (mainly monadics) should form a well pair of "verb" and "noun".

#### Good

```golang
type Customer struct {
    // fields
}

func Create(customer *Customer) {
    // Creating customer
}
```

---

### Avoid collateral effect

If your function promises to do something (this promise is on its name), so it must do only it. 

Programmers believe on functions names. Don't let their names fool them (us).

#### Bad

```golang
// This function only calculates... and sends an emaaaail?
func Calculate(firstNumber, secondNumber int) int {
    result := firstNumber + secondNumber
    SendResultOnEmail(result)
    return result
}

// Sends an email
func SendResultOnEmail(result int) {
    // Getting the result and sending an email
}
```

#### Good

```golang
// Now this function only calculates
func Calculate(firstNumber, secondNumber int) int {
    return firstNumber + secondNumber
}

// Still sending an email
func SendResultOnEmail(result int) {
    // Getting the result and sending an email
}

// At least we have now a different function which does those two operations with a explicit name.
func CalculateAndSendResultOnEmail(firstNumber, secondNumber int) int {
    result := Calculate(firstNumber, secondNumber)
    SendResultOnEmail(result)
    return result
}
```

---
