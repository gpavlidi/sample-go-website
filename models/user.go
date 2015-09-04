package models

import ()

type User struct {
	id        int64
	firstName string
	lastName  string
}

func NewUser(firstName, lastName string) *User {
	return &User{firstName: firstName, lastName: lastName}
}
