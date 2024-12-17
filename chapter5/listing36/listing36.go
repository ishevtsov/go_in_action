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

// notify implements a method with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>",
		u.name,
		u.email)
}

func main() {
	// Create a value of type User and send a notification.
	u := user{"Bill", "bill@email.com"}

	// ./listing36.go:28:19: cannot use u (variable of struct type user)
	// as notifier value in argument to sendNotification:
	// user does not implement notifier (method notify has pointer receiver)
	// fix sendNotification(&u)
	sendNotification(u)
}

func sendNotification(n notifier) {
	n.notify()
}
