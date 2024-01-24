package main

import (
	"errors"
	"fmt"
)

/*
  A class or module should have one reason to change.  A type of function should
  have a single responsibility, making code easier to understand and maintain.
*/

// Consider this example, the BadUser objects are responsible for storing user
// Meta data as well as saving.  This is more two responsibilities.
type BadUser struct {
	FirstName string
	LastName  string
}

func (u *BadUser) FullName() string {
	return u.FirstName + u.LastName
}

func (u *BadUser) Save() error {
	// Save user to a database
	return errors.New("Might Error...")
}

// Let's adhere to the single responsibility principle
// and improve this code.
type User struct {
	Firstname string
	LastName  string
}

func (u *User) FullName() string {
	return u.Firstname + u.LastName
}

type UserRepository struct {
}

func (u *UserRepository) Save(user *User) error {
	// Save user to the database
	return errors.New("Might error..")
}

func main() {
	fmt.Println("Single Responsibility Principle")
	repository := &UserRepository{}
	user := &User{Firstname: "Joe", LastName: "Bloggs"}
	if err := repository.Save(user); err != nil {
		fmt.Printf(fmt.Sprintf("%s\n", err.Error()))
	}
}
