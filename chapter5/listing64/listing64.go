package main

import "fmt"

func main() {
	counter := counters.alertCounter(10)
	fmt.Printf("Counter: %d\n", counter)
}
