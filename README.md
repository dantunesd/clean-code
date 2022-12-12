
# Clean Code

A summary of clean code best practices based on its own book.

## Content

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
    - [Indentation blocks](#indentation-blocks)
    - [One abstraction level per function](#one-abstraction-level-per-function)
    - [Decrescent reading](#decrescent-reading)
    - [Function parameters](#function-parameters)
    - [Verbs and key-words](#verbs-and-key-words)
    - [Avoid collateral effect](#avoid-collateral-effect)
    - [Output parameters](#output-parameters)
    - [Command-Query separation](#command-query-separation)
    - [Prefer exceptions instead of returning error codes](#prefer-exceptions-instead-of-returning-error-codes)
    - [Avoid repetition](#avoid-repetition)
- [Comments](#comments)
    - [Avoid comments](#avoid-comments)
    - [Avoid redundant comments](#avoid-redundant-comments)
    - [Avoid imperative comments](#avoid-imperative-comments)


## Naming

### Meaningful names

Use names that reveals its purpose. 

Choosing a good name may take some time. But you will economize later.

All those who reads you code will be grateful. 

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

Avoid naming things which may trick other programmers.

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

Avoid coding only for the compiler or interpreter. Code for other programmers.

Let your code be explicit, name everything with meaningful distinctions.

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

We people are good with words, with our language. Not using it is a waste.

Name things with pronounceable words.

#### Bad

```golang
// It's not a word, it's hard to pronounce
var crtTm = time.Now()
```

#### Good

```golang
// Words, they are easy to pronounce
var currentTime = time.Now()
```

---

### Searchable names

Names with one letter or number is hard to find throughout a code.

If a variable is used in many places, prefer to name it well and seachable.

#### Bad

```golang
// Imagine searching for just "a" on your IDE
var a = 22
```

#### Good

```golang
// Easier to find all ocurrencies
var averageMonthlyWorkingDays = 22
```

---

### Avoid mental mapping

Prefer always to use conventions and ubiquitous language(DDD).

It will facilitate everybody who is reading your code, avoiding them to have to translate your code to what it was supposed to be. 

#### Bad

```golang
for ç := 0; ç < count; ç++ {
    fmt.Println(ç)
}
```


```golang
var motorizedMachine Vehicle
var vehicleThatTakesPeople Vehicle
```

#### Good

```golang
for i := 0; i < count; i++ {
    fmt.Println(i)
}
```

```golang
var car Vehicle
var taxi Vehicle
```

---

### Class names

Classes should be nouns.

Avoid generic names such as Data, Info, Manager, Process, and so on.

#### Bad

```golang
type CarInfo struct {
    Name string
}

type CustomerData struct {
    Name string
}
```

#### Good

```golang
type Car struct {
    Name string
}

type Customer struct {
    Name string
}
```

---

### Method names

Methods are actions, actions are verbs. Methods should contain a verb on its name.

Preffer following some conventions like:  

Get: to get a field

Set: to set a value on a field

Is: to answer if something is something else

However, there's a golang convention which says you don't need to prefix "get" when getting something. So this topic is controversial :X

#### Bad

```golang
type Employee struct {
    name        string
    nationality string
}

// What method "Name" is doing? Where is the verb?
func (e *Employee) Name() string {
    return e.name
}

// What a "Brazilian" method should do?
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

// Get the Name
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

Use the same word for the same abstract context. Naming them different may confuses other programmers. 

Like, what's the difference between "fetching" or "getting" something? Or "creating" or "saving"?

#### Bad

```golang
func (c *Car) GetName() string {
    return c.Name
}

func (c *Customer) RetrieveName() string {
    return c.Name
}

func (a *Animal) FetchName() string {
    return a.Name
}
```

#### Good

```golang
func (c *Car) GetName() string {
    return c.Name
}

func (c *Customer) GetName() string {
    return c.Name
}

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
type BusStrategy struct{
    // ...
}

type CarFactory struct{
        // ...
}

type ActiveAccountState struct{
        // ...
}
```

---

### Meaningful context

Only after reading these 3 variables below you notice they belong to a bigger context: Address.

Now imagine if they were all spread out in some code. Would you still make it?

There are few words which are meaningful for its own. This is why we put such variables into classes, structs, etc

#### Bad

```golang
var number int
var state string
var street string
```

#### Good

```golang
// Group them together into a Class or Struct when it's possible.
type Address struct {
    Street string
    Number int
    State  string
}
```

```golang
// If there is no other way, add prefixes.
var addressNumber int
var addressState string
var adressStreet string
```

---

### Avoid unnecessary context

Well, this is the opposite from the topic above.

If your variables or classes are well contextualized, Don't give them unnecessary context.

#### Bad

```golang
package whatever

type AddressFromWhateverPackage struct {
    Street        string
    AddressNumber int
    AdressState   string
}
```

#### Good

```golang
package whatever

type Address struct {
    Street string
    Number int
    State  string
}
```

---

## Functions

### Small

Great functions are hard to understand because they do a lot different things, have lots of statements chained, have different levels of abstraction, have a confused reading flow, and so on...

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

### Indentation blocks

To have good functions, in addition of being small, they must have few indentation levels. At most 2 levels.

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

The problem with this is that your function will become harder to understand, probably because it does a lot of things with these parameters. It's hard to test all the combinations, you may change the order of them while calling it, and so on.

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

### Output parameters
Parameters are generally known as "input" to functions. But there are also some parameters known as "output parameters". Ex:


```golang
AppendFooter(r)
```

Is AppendFooter appending "r" as footer to another object or it's appending a footer on "r"?

Generally, avoid using parameters as "output parameters" (AppendFooter appends footer on "r"). If any function needs to change the state of any object, delegate this responsibility to its own object/struct.

#### Bad

```golang
type Report struct {
    Footer string
}

func AppendFooter(report *Report) {
    report.Footer = "footer example"

    // This function is not appending the report on another object
    // It's appending a footer on report. It's changing the Report state
}
```

#### Good

```golang
type Report struct {
    Footer string
}

func (r *Report) AppendFooter() {
    r.Footer = "footer example"

    // Now this function is changing the state of its own "object".
}
```

---

### Command-Query separation

Functions should do or answer something, not both. Doing both things could generate confusion. Avoid it.

Try to implement command-query separation principle

#### Bad

```golang
package main

type Car struct {
    Name string
}

func (c *Car) GetName() string {
    if c.Name != "" {
        c.Name = "Default name"

        // doing more things here related to this object
        // ...
        // ...
    }

    return c.Name
}
```

#### Good

```golang
package main

type Car struct {
    Name string
}

func (c *Car) GetName() string {
    return c.Name
}

func (c *Car) SetName(name string) {
    c.Name = name
}
```

---

### Prefer exceptions instead of returning error codes

Well, golang doesn't have exceptions. But you still don't need to return error codes.

Returning error codes leds to creating many indentation blocks and giving the caller responsibility to validate these codes.

Prefer to return golang "errors" and handle them.

Also, error handling (like try/catch blocks) is "one thing". Avoid mixing them with business logic. So, create a separate function for it.

#### Bad

```golang
package main

import "log"

// A simple example of an error code
const E_OK = "OK"

// A simple struct
type Customer struct {
    // fields
}

// This function needs to know about error codes to continue its logic.
func Create(customer *Customer) {
    if Validate(customer) == E_OK {
        if Save(customer) == E_OK {
            return
        } else {
            log.Println("failed to save customer")
        }
    } else {
        log.Println("failed to validate customer")
    }
}

// Returning error codes.
func Validate(customer *Customer) string {
    // validating
    return E_OK
}

// Returning error codes.
func Save(customer *Customer) string {
    // saving on database
    return E_OK
}
```

#### Good

```golang
package main

import "log"

// A simple struct
type Customer struct {
    // fields
}

// A "main" function applying something like a "try-catch" block.
func CreateAndLogError(customer *Customer) {
    if err := Create(customer); err != nil {
        log.Println(err.Error())
    }
}

// Now this function doesn't need to bother with codes neither logging, it just passes on the errors.
func Create(customer *Customer) error {
    if err := Validate(customer); err != nil {
        return err
    }

    if err := Save(customer); err != nil {
        return err
    }

    return nil
}

// Returning errors instead of error codes.
func Validate(customer *Customer) error {
    // validating
    return errors.New("failed to validate customer")
}

// Returning errors instead of error codes.
func Save(customer *Customer) error {
    // saving on database
    return errors.New("failed to save customer")
}
```

---

### Avoid repetition

Follow [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself) principle.

---

## Comments

_Do not insert comments in a bad code, rewrite it._

### Avoid comments

Your code change, evolves, it's moved from a place to another, it's renamed, deleted, improved, and so on.

But comments rarely evolves at the same rate. 

Comments lie, get outdated, pollute and obfuscate your code. They tell things that aren't true anymore.

The only thing that can tell you exactly what your code does... it is the code itself!

Don't waste energy explaning a bad code with a comment. Clean it.


#### Bad

```golang
// Verifies if employee has the rights to all benefits
if employee.age > 65 && employee.contract == HOURLY_CONTRACT {
    ...
}
```

#### Good

```golang
if employee.isEligibleForFullBenefits() {
    ...
}
```

---

### Avoid redundant comments

Avoid commenting a code that already explains itself. Reading the comment will be just a waste of time.

#### Bad

```golang
// If one is higher than two, prints "hey", if not, prints "ho"
if 1 > 2 {
    fmt.Println("hey")
} else {
    fmt.Println("ho")
}
```

#### Good

```golang
if 1 > 2 {
    fmt.Println("hey")
} else {
    fmt.Println("ho")
}
```

#### Better

if you have a complex code (this is not the case), extract it to smaller functions and give then a self explained name.

```golang
func printsHeyOrHo() {
    if oneIsHigherThanTwo() {
        printsHey()
    } else {
        printsHo()
    }
}
```

#### Some other bad ones

```golang
// Day of month
var dayOfMonth int
```
Really?

```golang
/**
 * @return Returns the day of the month
 */ 
func getDayOfTheMounth() int {
    return dayOfMonth
}
```

Hmm, got it.

---

### Avoid imperative comments

It's basicaly the same as redundant comments. Adding code blocks to explain something that is self explained is unecessary.

### Bad

```golang
/**
 * @param title The title of the CD
 * @param author The author of the CD
 */
func addCD(title, author string) {
    ...
}
```

### Good

```golang
func addCD(title, author string) {
    ...
}
```

---