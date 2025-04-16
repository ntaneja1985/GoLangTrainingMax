package main

import (
	"examples.com/structs/user"
	"fmt"
)

type customString string

func (str customString) log() {
	fmt.Println(str)
}

func main() {
	var name customString = "John"
	name.log()

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	var appUser *user.User
	admin, _ := user.NewAdmin("email@example.com", "test123")
	admin.OutputUserDetails()
	//appUser = &user.User{
	//	firstName: firstName,
	//}
	appUser, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}
	// ... do something awesome with that gathered data!
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

//func outputUserDetails(u *user) {
//	//Note we are not dereferencing the pointer here.
//	//fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate, (*u).createdAt)
//	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
//}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
