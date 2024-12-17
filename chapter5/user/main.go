package user

// user defines a user in the program.
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

// Declare a variable of type user.
var bill user

// Declare a variable of type user and initialize all the fields.
lisa := user{
	name: "Lisa",
	email: "lisa@email.com",
	ext: 123,
	privileged: true,
}

// admin represents an admin user with privileges.
type admin struct{
	person user
	level string
}

fred := admin{
	person: user{
		name: "Fred",
		email: "fred@email.com",
		ext: 321,
		privileged: true,
	},
	level: "super",
}
