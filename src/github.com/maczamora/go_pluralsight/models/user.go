package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	Users  []*User // slice of pointers to User object
	nextID = 1
)
