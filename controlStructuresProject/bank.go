package main

import "fmt"
import "examples.com/bank/fileops"
import "github.com/Pallinder/go-randomdata"

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("--------")
		//panic("Cant continue, sorry.")
	}
	fmt.Println("Welcome to Nishant Banking Services")
	fmt.Println("--------")
	fmt.Println("For any issues, please contact this phone number", randomdata.PhoneNumber())

	for {
		presentOptions()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		//wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
			if err != nil {
				fmt.Println("Error")
				fmt.Println(err)
				fmt.Println("--------")
				return
			}
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("Enter the amount you want to deposit: ")
			var amount float64
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Your cannot deposit an amount less than or equal to zero")
				//No code is executed after this return. It stops the execution of the function
				continue
			}
			accountBalance += amount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Your updated balance is: ", accountBalance)
		case 3:
			fmt.Print("Enter the amount you want to withdraw: ")
			var amount float64
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Your cannot withdraw an amount less than or equal to zero")
				//No code is executed after this return. It stops the execution of the function
				continue
			}

			if amount > accountBalance {
				fmt.Println("You cannot withdraw this amount as you dont have sufficient balance")
				continue
			}
			accountBalance -= amount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Your updated balance is: ", accountBalance)
		default:
			fmt.Println("Good bye!")
			fmt.Println("Thanks for using our bank")
			return
			//break
		}
	}
}
