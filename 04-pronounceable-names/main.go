package main

import (
	"time"
)

// Bad: It's not a word, difficult to pronounce and it confuses
var Clt = time.Now()

// Good: it's a word, easy to pronounce
var CurrentLocalTime = time.Now()
