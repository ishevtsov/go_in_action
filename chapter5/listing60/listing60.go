// Sample program to show how to embed a type into another type and
// the relationship between the inner and outer type.
package main

import "fmt"

// notifier is an interface that defined notification
// type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method that can be called via
// a value of type user.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// admin represents an admin user with privileges.
type admin struct {
	user  // Embedded Type
	level string
}

// notify implements a method that can be called via
// a value of type admin.
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

func main() {
	// Create an admin user.
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// Send the admin user a notification.
	// The embedded inner type's implementation of the
	// interface is "promoted" to the outer type.
	sendNotification(&ad)

	// We can access the inner type's method directly.
	ad.user.notify()

	// The inner type's method is NOT promoted.
	ad.notify()
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
