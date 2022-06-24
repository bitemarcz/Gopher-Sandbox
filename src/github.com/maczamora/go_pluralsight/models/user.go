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

func GetUsers() []*User {
	return Users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	Users = append(Users, &u)
	return u, nil
}
