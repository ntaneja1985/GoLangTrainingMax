package user

import (
	"errors"
	"fmt"
	"time"
)

// User struct: Group related fields together
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// New : Not a feature built in go, just a pattern/convention
// Think of constructor functions
func New(firstName, lastName, birthdate string) (*User, error) {
	//Can be used to do validations in one central place
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name and Last name and birthdate must be provided")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password must be provided")
	}
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}, nil
}
