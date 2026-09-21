package model

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

type Admin struct {
	email    string
	password string
	User
}

func (a *Admin) Email(email string) string {
	return a.email
}

func (a *Admin) SetEmail(email string) {
	(*a).email = email
}

func (a *Admin) Password(password string) string {
	return a.password
}

func (a *Admin) SetPassword(password string) {
	(*a).password = password
}

// func (a *Admin) User() User {
// 	return a.User
// }

// func (a *Admin) SetUser(user User) {
// 	(*a).User = user
// }


func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		fmt.Println("Error: All fields are required.")
		return nil, errors.New("all fields are required")
	}
	
	return &User{
		firstName: firstName,
		lastName: lastName	,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) (*Admin, error) {
	if email == "" || password == "" {
		fmt.Println("Error: Email and password are required.")
		return nil, errors.New("email and password are required")
	}

	return &Admin{
		email:    email,
		password: password,
		User:     User{
			firstName: "Admin",
			lastName:  "Admin",
			birthdate: "01/01/1970",
			createdAt: time.Now(),
		},
	}, nil

}
