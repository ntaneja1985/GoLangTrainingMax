package main

import "fmt"

func main() {
	age := 32 //regular variable

	//Has address of age variable
	var agePointer *int
	agePointer = &age

	fmt.Println("Age:", *agePointer)
	getAdultYears(&age)
	fmt.Println("Get Adult Years:", age)
}

func getAdultYears(age *int) {
	//Store the value directly inside the variable referenced by the pointer
	*age = *age - 18
}
