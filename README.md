
# Clean Code

A brief description of what this project does and who it's for

## Naming

### Meaningful names

#### Bad

```golang
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
func SumSquare(numbers []int) int {
	var totalSum int
	for _, number := range numbers {
		totalSum += number * number
	}
	return totalSum
}
```