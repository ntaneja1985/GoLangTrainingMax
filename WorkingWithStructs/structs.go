package main

import "fmt"
import "time"

// Group related fields together
type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	var appUser user
	appUser = user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}
	// ... do something awesome with that gathered data!
	appUser.outputUserDetails()
}

//func outputUserDetails(u *user) {
//	//Note we are not dereferencing the pointer here.
//	//fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate, (*u).createdAt)
//	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
//}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
