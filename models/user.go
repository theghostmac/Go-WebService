package models

// create User struct with f,l names and id
type User struct {
	FirstName string
	LastName  string
	ID        int
}

// create variables with users of type slice of user struct and nextID of type
var (
	users  []*User
	nextID = 1
)

// create a function to return users slice
func ReturnUsers() []*User {
	return users
}

// create a function to add user and return user or error
func AddUser(usr User) {

}

//
