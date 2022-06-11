package main

import (
	"fmt"
	"time"
)

func main() {
	// Bad: it's not a word, difficult to pronounce and it confuses
	clt := time.Now()

	// Good: it's a word, easy to pronounce
	currentLocalTime := time.Now()

	fmt.Println(clt, currentLocalTime)
}
