package main

//Imports are not shared across different files in the same package
import "fmt"

func presentOptions() {
	fmt.Println("What do you want to do")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}
