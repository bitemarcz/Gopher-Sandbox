package models

import (
	"errors"
	"fmt"
)

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
	if u.ID != 0 {
		return User{}, errors.New("new user must not include id ot it must be set")
	}
	u.ID = nextID
	nextID++
	Users = append(Users, &u)
	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range Users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found,", id)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range Users {
		if candidate.ID == u.ID {
			Users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

func RemoveUserById(id int) error {
	for i, u := range Users {
		if u.ID == id {
			Users = append(Users[:i], Users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
