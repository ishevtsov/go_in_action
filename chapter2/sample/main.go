package main

import (
	"log"
	"os"

	_ "sample/matchers"
	"sample/search"
)

func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
