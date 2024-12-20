// Sample program to show how unexported fields from an exported
// struct type can't be accessed directly.
package main

import (
	"chapter5/listing74/entities"
	"fmt"
)

func main() {
	// Create a value of type Admin from the entities package.
	a := entities.Admin{
		Rights: 10,
	}

	// Set the exported fields from the unexported
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}
