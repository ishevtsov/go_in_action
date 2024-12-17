// Sample program to show how unexported fields from an exported
// struct type can't be accessed directly.
package main

import (
	"chapter5/listing71/entities"
	"fmt"
)

func main() {
	u := entities.User{
		Name:  "Bill",
		email: "bill@email.com",
	}

	// listing71/listing71.go:13:3: cannot refer to unexported field email in struct literal of type entities.User

	fmt.Printf("User: %v\n", u)
}
