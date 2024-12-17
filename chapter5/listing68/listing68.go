package main

import (
	"chapter5/listing68/counters"
	"fmt"
)

func main() {
	counter := counters.New(10)
	fmt.Printf("Counter: %d\n", counter)
}
