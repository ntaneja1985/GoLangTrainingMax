# GoLangTrainingMax
Golang Training

- Golang is an open source programming language developed by Google
```go
package main

import "fmt"

func main(){
	fmt.Println("Hello, World!")
}

```
- ![alt text](image-1.png)
- ![img_8.png](img_8.png)

## Go Essentials

### Key Components of a Go Program
- When we write Go Code, we split our code across packages.
- We must have atleast one package per application.
- A single package can also be split across multiple files.
- ![img.png](img.png)
- ![img_1.png](img_1.png)
- We can have multiple packages in one Go Project
- The idea of package is similar to namespaces in C#
- Packages help to organize our code
- We can export and import features across our files.
```go
import "fmt"
```
- This fmt package is imported from Go Standard library
- Go comes with a huge standard packages library of built-in packages i.e lot of them come installed
- ![img_2.png](img_2.png)
- "main" is a special reserved package name, and it tells go that it is main entry point of the application we are building.
- We run "go build" and it tells go to create an executable file which can run on systems which dont have go installed
- ![img_3.png](img_3.png)
- Go has modules and each module contains of multiple packages
- Think of module as a C# project(.csproj) file
- We need to run a special command
```shell
go mod init example.com/first-app
```
- It converts our code into a module and creates a go.mod file which has the following contents
```text
module example.com/first-app

go 1.24.2

```
- Now we can run "go build" command
- This then creates an .exe file if we are running on Windows
- On macOS and Linux we just get a filename with the name of the module and without an extension
- We run it by writing ./first-app
- ![img_4.png](img_4.png)
- This file can be run without having go installed on the system
- We also need to name our main function as main()
- ![img_5.png](img_5.png)
- This main function specifies the code to be executed when application starts running.
- **Any package can have only one main() function.**
- ![img_6.png](img_6.png)
- If we are building a library which just has helper functions, we don't need a main() function.
- Such a package needs to be just imported, and we can use functions from that package. 

### Values and Types
- In every programming language we store values in some container(called variable), and then reference that container to use that value from any point in our code.
- In Go we can create a variable using the "var" keyword.
- ![img_7.png](img_7.png)
- Go is a statically typed language. Every value in Go Program is a specific type.
- investmentAmount is of type int and expectedReturnRate is of type float64(allows decimals)
- These types matter in the calculation and we cannot just mix and match types.
- We can do type casting(similar to what we have in C#)
- ![img_9.png](img_9.png)
- The float64() function converts the investmentAmount to a float64 variable. So technically 1000 becomes 1000.0
- In Go we have no built in power operator. We have "math" package
- ![img_10.png](img_10.png)
- ![img_11.png](img_11.png)
- Go comes with a couple of built-in basic types:
- int => A number WITHOUT decimal places (e.g., -5, 10, 12 etc)
- float64 => A number WITH decimal places (e.g., -5.2, 10.123, 12.9 etc)
- string => A text value (created via double quotes or backticks: "Hello World", `Hi everyone')
- bool => true or false
- But there also are some noteworthy "niche" basic types which you'll typically not need that often but which you should still know about
- uint => An unsigned integer which means a strictly non-negative integer (e.g., 0, 10, 255 etc)
- int32 => A 32-bit signed integer, which is an integer with a specific range from -2,147,483,648 to 2,147,483,647 (e.g., -1234567890, 1234567890)
- rune => An alias for int32, represents a Unicode code point (i.e., a single character), and is used when dealing with Unicode characters (e.g., 'a', 'ñ', '世')
- uint32 => A 32-bit unsigned integer, an integer that can represent values from 0 to 4,294,967,295
- int64 => A 64-bit signed integer, an integer that can represent a larger range of values, specifically from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
- There also are more types like int8 or uint8 which work in the same way (=> integer with smaller number range)

### Null Values
- All Go value types come with a so-called "null value" which is the value stored in a variable if no other value is explicitly set.
- For example, the following int variable would have a default value of 0 (because 0 is the null value of int, int32 etc):
```go
var age int // age is 0
```
- Here's a list of the null values for the different types:
- int => 0
- float64 => 0.0
- string => "" (i.e., an empty string)
- bool => false

## Outputting Values
- ![img_12.png](img_12.png)
- We can just run this command to execute the module
```shell
go run .
```
- ![img_13.png](img_13.png)
- ![img_14.png](img_14.png)

### Automatically infer types in Go
- ![img_15.png](img_15.png)
- We can also declare several variables in the same line
- This is not allowed:
- ![img_16.png](img_16.png)
- If we do specify the type of the variables in single line, then they all will have the same type
- ![img_17.png](img_17.png)
- In the above code both investmentAmount and years will have the type float64
- We can automatically infer types with the following code:
- ![img_18.png](img_18.png)
- We can now move all our variables in one line:
- ![img_19.png](img_19.png)
- This however makes the code difficult to read, so to know what variable is of what type we can always use this code:
```go
var investmentAmount float64 = 1000;
```
#### The difference between const and var in Go is that variable can be reassigned, but const cannot be reassigned.
- ![img_20.png](img_20.png)
#### To accept input from the user we use the Scan() function
- ![img_21.png](img_21.png)
```go
func main() {
const inflationRate float64 = 6.5
//If we don't assign a value to investmentAmount variable, it will take the default null value for float64 which is 0.0
var investmentAmount float64
years, expectedReturnRate := 10.0, 5.5

//investmentAmount already has a value of 1000 initially, but it gets overridden by the value entered by the user
//Here we directly update the value inside the data container(variable) investmentAmount in this case, but passing its pointer
//The pointer is a memory reference
fmt.Scan(&investmentAmount)

futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
//Start a new line after outputting the value
fmt.Println(futureValue)
fmt.Println(futureRealValue)
}
```
- If we try to assign a value to constant by passing its memory reference we get an error:
- ![img_22.png](img_22.png)
- Code to get all the values from the user is like this:
```go
func main() {
	const inflationRate float64 = 6.5
	//If we don't assign a value to investmentAmount variable, it will take the default null value for float64 which is 0.0
	var investmentAmount float64
	years, expectedReturnRate := 10.0, 5.5

	fmt.Print("Investment Amount: ")

	//investmentAmount already has a value of 1000 initially, but it gets overridden by the value entered by the user
	//Here we directly update the value inside the data container(variable) investmentAmount in this case, but passing its pointer
	//The pointer is a memory reference
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the years: ")
	fmt.Scan(&years)
	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	//Start a new line after outputting the value
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
```
- The fmt.Scan() function is a great function for easily fetching & using user input through the command line.
- But this function also has an important **limitation**:
- You can't (easily) fetch multi-word input values.
- Fetching text that consists of more than a single word is tricky with this function.
- We therefore, need an alternative to fmt.Scan() function

### Exercise
- ![img_23.png](img_23.png)
```go
package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter the revenue amount: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expenses amount: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax - ((earningsBeforeTax) * (taxRate / 100))
	ratio := earningsBeforeTax / earningsAfterTax

	fmt.Print("Your earnings before tax are: ")
	fmt.Println(earningsBeforeTax)
	fmt.Print("Your earnings after tax are: ")
	fmt.Println(earningsAfterTax)
	fmt.Print("Your ratio is: ")
	fmt.Println(ratio)

}

```
### Formatting Strings
- Comments are ignored by Go Compiler and are just there for developer's reference
- Please note fmt.Printf() is similar to string.Format() in C# where we specify placeholders and arguments corresponding to them
```go
futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	//Start a new line after outputting the value
	//fmt.Println("Future value: ", futureValue)
	fmt.Printf("Future Value: %v\nFuture value(adjusted for inflation): %v\n", futureValue, futureRealValue)
	//fmt.Println("Future value(adjusted for inflation):", futureRealValue)
```
- ![img_24.png](img_24.png)
- ![img_25.png](img_25.png)
- Similar to %v, we also have other literals we can use like %f for decimal points.
- We can change our code to print the values with 0 decimal points like this
```go
	//Start a new line after outputting the value
	//fmt.Println("Future value: ", futureValue)
	fmt.Printf("Future Value: %.0f\nFuture value(adjusted for inflation): %.0f\n", futureValue, futureRealValue)
	//fmt.Println("Future value(adjusted for inflation):", futureRealValue)
```
### Creating Formatted Strings
- To create a formatted string we use the Sprintf() function
```go
//Creates a new text string as per the format specifier and returns the resulting string
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value(adjusted for inflation): %.1f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
```
- To build multiline strings, use backticks``
```go
fmt.Printf(`Future Value: %.0f
		Future value(adjusted for inflation): %.0f`, futureValue, futureRealValue)
```

### Understanding Functions
- We have already seen the main() function
- They can be considered as code-on-demand blocks
- main() function is invoked by Go when the Go Program starts
- Similarly, we have Math.Pow(), Printf, PrintLn functions
- ![img_26.png](img_26.png)
- ![img_27.png](img_27.png)
- ![img_28.png](img_28.png)
- A function can be defined using the "func" keyword.
```go
//Here both text and text2 accept a type of string
func outputText(text, text2 string) {
	fmt.Print()
}
```
- We can use custom functions like this
```go
func main() {
    const inflationRate float64 = 6.5
	//If we don't assign a value to investmentAmount variable, it will take the default null value for float64 which is 0.0
	var investmentAmount float64
	years, expectedReturnRate := 10.0, 5.5

	outputText("Investment Amount: ")

	//investmentAmount already has a value of 1000 initially, but it gets overridden by the value entered by the user
	//Here we directly update the value inside the data container(variable) investmentAmount in this case, but passing its pointer
	//The pointer is a memory reference
	fmt.Scan(&investmentAmount)

	outputText("Enter the years: ")
	fmt.Scan(&years)
	outputText("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	//Creates a new text string as per the format specifier and returns the resulting string
	//formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	//formattedRFV := fmt.Sprintf("Future value(adjusted for inflation): %.1f\n", futureRealValue)
	//fmt.Print(formattedFV, formattedRFV)
	//Start a new line after outputting the value
	//fmt.Println("Future value: ", futureValue)
	fmt.Printf(`Future Value: %.0f
		Future value(adjusted for inflation): %.0f`, futureValue, futureRealValue)
	//fmt.Println("Future value(adjusted for inflation):", futureRealValue)

}

// Here both text and text2 accept a type of string
func outputText(text string) {
	fmt.Print(text)
}

```
### Returning Values from Functions and Variable Scope
- In Go, unlike in many other programming languages, you can actually pretty easily return multiple values
simply by separating them with a comma here.
- In Go, as in many other programming languages, any variables or constants that are declared in a function
are scoped to that function and are only available in that function.
- We can see that in practice here
```go
//Return 2 values from the same function(similar to tuples in C#)
func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

//Use the above function like this
futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

```

### An Alternative return value syntax
```go
func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	//return fv, rfv
	return
}
```
### Exercise: Working with Functions
- Separate the profit and loss problem into separate functions
```go
package main

import "fmt"

func main() {
	revenue := getTextFromUser("Enter the revenue amount: ")
	expenses := getTextFromUser("Enter the expenses amount: ")
	taxRate := getTextFromUser("Enter the tax rate: ")
	earningsBeforeTax, earningsAfterTax, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Print("Your earnings before tax are: ")
	fmt.Printf("%.1f", earningsBeforeTax)
	fmt.Print("Your earnings after tax are: ")
	fmt.Printf("%.1f", earningsAfterTax)
	fmt.Printf(`Ratio is %.1f`, ratio)

}

func getTextFromUser(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax - ((earningsBeforeTax) * (taxRate / 100))
	ratio := earningsBeforeTax / earningsAfterTax
	return earningsBeforeTax, earningsAfterTax, ratio
}

```

## Control Structures in Go
- We can use if-else statements or nested if-else statements
- We also can use the "for" keyword for looping. Unlike other programming languages which have while loop, do-while loops, in Go, we just have for loop
```go
fmt.Println("Welcome to Nishant Banking Services")
	fmt.Println("What do you want to do")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanf("%d", &choice)
	//wantsCheckBalance := choice == 1
	if choice == 1 {
		fmt.Println("Your balance is: ", accountBalance)
	} else if choice == 2 {
		fmt.Print("Enter the amount you want to deposit: ")
		var amount float64
		fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("Your cannot deposit an amount less than or equal to zero")
			//No code is executed after this return. It stops the execution of the function
			return
		}
		accountBalance += amount
		fmt.Println("Your updated balance is: ", accountBalance)
	} else if choice == 3 {
		fmt.Print("Enter the amount you want to withdraw: ")
		var amount float64
		fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("Your cannot withdraw an amount less than or equal to zero")
			//No code is executed after this return. It stops the execution of the function
			return
		}

		if amount > accountBalance {
			fmt.Println("You cannot withdraw this amount as you dont have sufficient balance")
			return
		}
		accountBalance -= amount
		fmt.Println("Your updated balance is: ", accountBalance)
	} else {
		fmt.Println("Good bye!")
	}

	fmt.Println("Your choice: ", choice)
```
- For infinite loops just use the "for" keyword without any conditions
- This is similar to while(true) in C#
- We can use "break" keyword to break out of the loop
- We can use the "continue" keyword to skip the current loop iteration and move onto the next iteration
- Besides the for variations introduced in the last lectures, there also is another common variation (which will also be shown again later in the course):
```go
//similar to the while condition in C#
for someCondition {
// do something ...
}
```
- someCondition is an expression that yields a boolean value or a variable that contains a boolean value (i.e., true or false).
- In that case, the loop will continue to execute the code inside the loop body until the condition / variable yields false.
```go
package main

import "fmt"

func main() {
	var accountBalance float64 = 1000
	fmt.Println("Welcome to Nishant Banking Services")

	for {
		fmt.Println("What do you want to do")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		//wantsCheckBalance := choice == 1
		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)
		} else if choice == 2 {
			fmt.Print("Enter the amount you want to deposit: ")
			var amount float64
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Your cannot deposit an amount less than or equal to zero")
				//No code is executed after this return. It stops the execution of the function
				continue
			}
			accountBalance += amount
			fmt.Println("Your updated balance is: ", accountBalance)
		} else if choice == 3 {
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
			fmt.Println("Your updated balance is: ", accountBalance)
		} else {
			fmt.Println("Good bye!")
			//return
			break
		}

	}

	fmt.Println("Thanks for using our bank")
}

```

### switch statements: an alternative to nested if-else statements
- In other languages after each case statement, we need to add "break" keyword so that no other case gets executed
- In Go, we don't have to do that
- If we have a situation where we have to break out of the loop, dont use switch statement, use the if-else statement
- We cannot use the break keyword inside the switch statement as it has a different meaning inside the switch statement
```go
package main

import "fmt"

func main() {
	var accountBalance float64 = 1000
	fmt.Println("Welcome to Nishant Banking Services")

	for {
		fmt.Println("What do you want to do")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		//wantsCheckBalance := choice == 1

		switch choice {
		case 1:
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
			fmt.Println("Your updated balance is: ", accountBalance)
		default:
			fmt.Println("Good bye!")
			fmt.Println("Thanks for using our bank")
			return
			//break
		}
	}
}

```

## Working with Files in Go
- ![img_29.png](img_29.png)
- We can write to files by importing the "os" package
```go
package main

import "fmt"
import "os"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	//0644 is a way of encoding file permissions. It basically means that owner can read/write file
	//other users can only read from it.
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}
```
### Reading from a file
- ![img_30.png](img_30.png)
- We need the strconv package if we want to convert from string to float64
```go
package main

import "fmt"
import "os"
import "strconv"

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	//0644 is a way of encoding file permissions. It basically means that owner can read/write file
	//other users can only read from it.
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() float64 {
	//We are using the discard operator to reject the second value returned by the function
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}
```

### Handling Errors
- In go error handling works differently compared to other programming languages
- In other programming languages, if the app crashes, we are expected to enclose our code in try-catch blocks
- Thats not how it works in Go though. 
- Here errors dont crash the application
- In our case ReadFile() returns an empty byte[] if it doesnot find the file.
- But notice, the function also returns a second value called error variable.
```go
package main

import "fmt"
import "os"
import "strconv"
import "errors"

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	//0644 is a way of encoding file permissions. It basically means that owner can read/write file
	//other users can only read from it.
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	//We are using the discard operator to reject the second value returned by the function
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("cannot read balance from file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("cannot parse balance from file")
	}
	return balance, nil
}

func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("--------")
	}
}
```

### Time to Panic!!
- ![img_31.png](img_31.png)
- In Go, a panic is a mechanism used to handle unexpected or critical errors that make it impossible for a program to continue running normally.
- When a panic occurs, the program begins to unwind the stack, executing any deferred functions, and then it crashes.
- You can trigger a panic explicitly using the panic() function or implicitly when Go encounters certain runtime errors, like accessing an out-of-bounds slice index. 
```go
package main

import "fmt"

func main() {
    defer fmt.Println("This will run before the panic stack unwinds!")

    fmt.Println("About to panic...")
    panic("Something went wrong!") // Trigger panic

    fmt.Println("This will not be printed.")
}
```
- In our code we can do it as:
```go
func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("--------")
		panic("Cant continue, sorry.")
	}
}
```
- ![img_32.png](img_32.png)

### Exercise
- ![img_33.png](img_33.png)
```go
package main

import "fmt"
import "errors"
import "os"

const fileName = "financials.txt"

func main() {
	revenue, err := getTextFromUser("Enter the revenue amount: ")
	if err != nil {
		panic(err)
	}

	expenses, err := getTextFromUser("Enter the expenses amount: ")
	if err != nil {
		panic(err)
	}

	taxRate, err := getTextFromUser("Enter the tax rate: ")
	if err != nil {
		panic(err)
	}

	earningsBeforeTax, earningsAfterTax, ratio := calculateFinancials(revenue, expenses, taxRate)
	WriteFinancialDetailsToFile(earningsBeforeTax, earningsAfterTax, ratio)
	fmt.Print("Your earnings before tax are: ")
	fmt.Printf("%.1f\n", earningsBeforeTax)
	fmt.Print("Your earnings after tax are: ")
	fmt.Printf("%.1f\n", earningsAfterTax)
	fmt.Printf(`Ratio is %.1f`, ratio)
}

func getTextFromUser(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Value cannot be less than or equal to 0")
	}
	return userInput, nil
}

func WriteFinancialDetailsToFile(earningsBeforeTax float64, earningsAfterTax float64, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.2f", earningsBeforeTax, earningsAfterTax, ratio)
	os.WriteFile(fileName, []byte(results), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax - ((earningsBeforeTax) * (taxRate / 100))
	ratio := earningsBeforeTax / earningsAfterTax
	return earningsBeforeTax, earningsAfterTax, ratio
}

```

- In the above module, we learnt about the following
- ![img_34.png](img_34.png)

## Working with Packages
- We have already worked with fmt,os,errors package
- ![img_35.png](img_35.png)

### Splitting code across multiple files in the same package
- We can split the code across different files in the same package
- However, we need to specify the import statement again in different files
- This is because imports are not shared across different files in the same package
```go
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

```
### Why would you use more than one package
- ![img_36.png](img_36.png)
- ![img_37.png](img_37.png)
- ![img_38.png](img_38.png)
- If we were to draw an analogy, packages in Go are conceptually similar to namespaces in C#, though there are some differences in their usage and structure.
- Grouping Code:- In Go, a package is used to group related functions, types, and variables. Every Go file belongs to a package, defined at the top of the file with a package statement.
- In C#, namespaces serve a similar purpose, allowing developers to organize code into logical groups. You use the namespace keyword in C#.

- Importing/Using:- In Go, you import packages using the import keyword, and you refer to their exported names directly, like fmt.Println().
- In C#, you use the using directive to reference namespaces, allowing you to access classes, methods, or properties defined within that namespace, like Console.WriteLine().

- Scope and Visibility:- In Go, visibility of identifiers is determined by whether the name starts with an uppercase letter (public/exported) or lowercase letter (private/unexported).
- In C#, visibility is controlled through access modifiers such as public, private, protected, or internal.

- While Go's packages are tightly coupled with file organization, C# namespaces are more flexible and don't require the code in a namespace to reside in specific files.
- Main reason of using more than one package is **reusability**
- We should make functions more generic and put them in their own package

### Splitting code across multiple packages
- Having 2 packages in the same folder is not allowed in Go
- Every package must be in its own subfolder in Go
- Only functions or variables or constants that's start with an uppercase character are available in other packages.
- In Javascript any function we wish to export, we need to use the "export" keyword.
- In Go, we just need to use "casing". If it is uppercase, it is automatically exported. 
- All lowercase functions or variables or constants are considered private.
- ![img_39.png](img_39.png)
- So the code in fileops.go file is 
```go
package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, filename string) {
	valueText := fmt.Sprint(value)
	//0644 is a way of encoding file permissions. It basically means that owner can read/write file
	//other users can only read from it.
	os.WriteFile(filename, []byte(valueText), 0644)
}

func GetFloatFromFile(filename string) (float64, error) {
	//We are using the discard operator to reject the second value returned by the function
	data, err := os.ReadFile(filename)
	if err != nil {
		return 1000, errors.New("Failed to find file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored value")
	}
	return value, nil
}

```
- The code in bank.go is modified as follows:
```go
package main

import "fmt"
import "examples.com/bank/fileops"

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

```
### Using third party packages
- Go's standard library comes with loads of useful functions and packages built in.
- Sometimes, we need some functionality that is not part of standard library
- So, we can install third party libraries
- For e.g go random data package
- We can use Go's package discovery page to search for pages:  https://pkg.go.dev/
- So we need to install this package.
- Run the following commands in shell
```shell
go get github.com/Pallinder/go-randomdata

```
- Note it gets added to the go.mod file
- ![img_40.png](img_40.png)
- We can use it as follows:
```go
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
}
```

## Understanding Pointers
- ![img_41.png](img_41.png)
- Pointers are variables that store value addresses instead of values.
- ![img_42.png](img_42.png)
- Pointers avoid unnecessary value copies
- We can use pointers to directly mutate values
- ![img_43.png](img_43.png)
- ![img_44.png](img_44.png)
- ![img_45.png](img_45.png)
- Can lead to expected results
- ![img_46.png](img_46.png)
- To get the pointer to a specific variable we use "&" keyword
- To get the value stored inside a pointer, we can do dereferencing by using "*" in front of it
- In Go, we cannot perform pointer arithmetic
```go
package main

import "fmt"

func main() {
	age := 32 //regular variable

	//Has address of age variable
	var agePointer *int
	agePointer = &age
    
	//Get the actual value stored inside the agePointer
	fmt.Println("Age:", *agePointer)
	//Pass the address of the age variable
	adultYears := getAdultYears(&age)
	fmt.Println("Get Adult Years:", adultYears)
}

func getAdultYears(age *int) int {
	//Use dereferencing and get the value stored inside the agePointer
	return *age - 18
}

```
- Now there is no copy of age variable being created. We are just passing the address of the variable and using the value stored inside of it directly
- However, here we are getting no benefit by creating pointers

### Pointers Null Value
- All values in Go have a so-called "Null Value" - i.e., the value that's set as a default if no value is assigned to a variable.
- For example, the null value of an int variable is 0. Of a float64, it would be 0.0. Of a string, it's "".
- For a pointer, it's nil - a special value built-into Go.
- nil represents the absence of an address value - i.e., a pointer pointing at no address / no value in memory.

### Directly Mutate values
- We can store the value of the adultYears directly inside the age variable rather than creating a new adultYears variable
```go
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
	//This basically overwrites what is stored inside the memory location referenced by the pointer
	//Note since we are overriding(mutating) the original value, we dont have to return anything here
	//This is similar to ref keyword in C#
	*age = *age - 18
}

```
- Recall the fmt.Scan(&choice) function, we were already using a pointer. 
- Scan internally dereferences the pointer and overwrites the value stored inside choice with the value entered by the user
- We should not overdo pointers and make our code unnecessarily complex

## Structs and Custom Types
- We can use Structs which is a datatype that helps us to structure data and group related data together
- We can add methods,functions to structs also
- ![img_47.png](img_47.png)
- Structs are structures that allow you to group related together into one single value.
- In Go, we can create instances of a struct type
```go
package main

import "fmt"
import "time"

// Group related fields together
//Define a struct
type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	var appUser user
	//Create instance of a struct
	appUser = user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}
	// ... do something awesome with that gathered data!
	outputUserDetails(appUser)
}

func outputUserDetails(appUser user) {
	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthDate, fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		appUser.createdAt.Year(), appUser.createdAt.Month(), appUser.createdAt.Day(),
		appUser.createdAt.Hour(), appUser.createdAt.Minute(), appUser.createdAt.Second()))
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

```
### Structs and Pointers
- Instead of passing the struct value to the function, we can pass a pointer to the struct.
```go
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
	outputUserDetails(&appUser)
}

func outputUserDetails(u *user) {
	//Note we are not dereferencing the pointer here.
	//fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate, (*u).createdAt)
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}
```
- Important Point is that we are not dereferencing the pointer to the struct here to get its value
- This is something allowed by Go as we can directly access the property values of a struct 
- This can just be done by writing u.firstName, u.lastName etc.
- Ideally we should do it like this: (*u).firstName, (*u).lastName

### Introducing Methods inside a Struct
- We can also attach functions to structs
- A function attached to a struct is called a method
- We can achieve this with the following code:
```go
type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}
```
- Note that right after func we specify the struct to which we need to attach this method to
- This is called **Receiver** argument.
- We can now call this function from the main function as follows:
- Note we don't have to pass any arguments, because that method already has access to all values inside the struct.
```go
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
```

### Mutation Methods
- Note that in the above case when we define a method for a struct, we pass the struct itself as a receiver argument to the method
- This is actually like passing values to the method
- This means the original struct is copied over and a new struct is created which the method then operates on
```go
//Here we are creating a copy of the original struct
func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}
```
- To mutate or change values defined in the original struct, we can pass the struct as a pointer
```go
//Here we are operating on the original struct
func (u *user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}
```
- A more detailed example is like:
```go
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

func (u *user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func (u *user) clearUserName() {
	//Here we are mutating the original struct
	u.firstName = ""
	u.lastName = ""
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
	appUser.clearUserName()
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

```
### Using Creation/Constructor Functions
- So a creation or constructor function is not really a special function in Go
- It is just like any other function in Go which can return a struct after initializing it
- Not only initialization, it can also serve as a central place to do validations for all the values we are creating the struct with
- Just like we saw earlier, it can return a pointer to the struct as it will allow us to mutate the struct directly rather than its copy
```go
type user struct {
firstName string
lastName  string
birthDate string
createdAt time.Time
}


// Not a feature built in go, just a pattern/convention
// Think of constructor functions
func newUser(firstName, lastName, birthdate string) (*user, error) {
//Can be used to do validations in one central place
if firstName == "" || lastName == "" || birthdate == "" {
return nil, errors.New("first name and Last name and birthdate must be provided")
}

return &user{
firstName: firstName,
lastName:  lastName,
birthDate: birthdate,
createdAt: time.Now(),
}, nil
}

//Usage
func main() {
firstName := getUserData("Please enter your first name: ")
lastName := getUserData("Please enter your last name: ")
birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
var appUser *user
appUser, err := newUser(firstName, lastName, birthdate)
if err != nil {
fmt.Println(err)
return
}
// ... do something awesome with that gathered data!
appUser.outputUserDetails()
appUser.clearUserName()
appUser.outputUserDetails()
}

func getUserData(promptText string) string {
fmt.Print(promptText)
var value string
//Will scan the user input on the line only, i.e if enter key is pressed, the value is input as NULL
fmt.Scanln(&value)
return value
}

```

## Structs, Packages and Imports
- We can create a folder named user and copy-paste the user struct code along with its functions into a file called user.go
- Hence, we can create a new package called user
- The Uppercase naming convention doesn't just apply to the Struct name but also to all the fields in the struct
```go
package user

import (
	"errors"
	"fmt"
	"time"
)

// User struct: Group related fields together
type User struct {
	FirstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.FirstName, u.lastName, u.birthDate, u.createdAt)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.lastName = ""
}

// NewUser : Not a feature built in go, just a pattern/convention
// Think of constructor functions
func NewUser(firstName, lastName, birthdate string) (*User, error) {
	//Can be used to do validations in one central place
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name and Last name and birthdate must be provided")
	}

	return &User{
		FirstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}

```
- Now we can arrange the code in structs.go as follows:
```go
package main

import (
	"examples.com/structs/user"
	"fmt"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	var appUser *user.User
	
	appUser = &user.User{
		FirstName: firstName,
	}
	appUser, err := user.NewUser(firstName, lastName, birthdate)
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

```

### Exposing Constructors and a different constructor function name
- We should not be exposing fields of the struct and expose only the constructor function.
- It is convention to name the constructor function as just New()
- Remember errors package? errors.New("Some Error").
```go
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
```

### Struct Embedding
- Build a new Struct that builds upon an existing struct.
- Struct has no classes and no inheritance
- However, we can build a struct on top of another struct
```go
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User     User
}

//Create a constructor function for Admin Struct like this
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

// We can use it like this
admin, _ := user.NewAdmin("email@example.com", "test123")
admin.User.OutputUserDetails()
```
- However, rather explicitly typing the "type" the embedded struct, we can use it anonymously as an **anonymous embedding**
- We are able to directly access the methods or fields(if they are exposed(capital Letter naming)) of the anonymous embedded struct on the parent struct
```go
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

//Now we can use it like this
//We are able to directly access the methods of the anonymous embedded struct on the parent struct
admin, _ := user.NewAdmin("email@example.com", "test123")
admin.OutputUserDetails()
```

### Summary of Structs
- It is a custom type that can group related fields and functions
- We can create constructor functions to outsource the creation of structs in a reusable way and add validations
- We can put structs in packages and casing matters
- If we have a method that wants to edit a struct, we must pass the reference to the struct as a pointer
- In constructor function, we create a pointer to the newly created struct.
- We can also create structs based on other structs also called struct embeddings
- Structs are a very important feature in Go and they are used a lot in writing Go code.

### Creating other custom types and adding methods
- In Go, we can not only use "type" keyword just to create structs, 
- but it can also be used to assign an "alias" to other built-in types
- Why this is required? 
- This is required to create extension methods(just like C#) on built-in types
- For example, lets use the type keyword to add an alias for built-in type:"string"
- We can do this aliasing for int, float64 etc and even for more complex built in types.
```go
//alias
type customString string

//add an extension method by passing the type to a function
func (str customString) log() {
	fmt.Println(str)
}

//Use it like this
var name customString = "John"
name.log()
```

### Using Struct Tags
- Struct Tags are metadata that can be added to struct fields
- This is useful in case when we convert the struct to json as this is used by encoding/json package
```go
type Note struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
```
### Practice Project - Note-Taking App
- We will develop a simple app that will take the user input as Note Title and Note Content
- This will then display the note and also convert the note to json and save it to the file system
- We will create a struct "note" and its related constructor, display() and save() functions
- So for this we will create a separate package called note which will import into our main package
- We will make use of encodings/json package to convert our struct "note" to json and then save it within a file using the "os" package
- To read multi character input from the command line we will make use of bufio package
```go
//Code for Note Package
package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Title: %s, Content: %s, CreatedAt: %s\n", note.Title, note.Content, note.CreatedAt.Format(time.RFC3339))
}

func (note Note) Save() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"
	//Convert data to json
	//This package only converts to json the publicly available data
	jsonText, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonText, 0644)
}

func New(title, content string) (Note, error) {

	//Do validation
	if title == "" || content == "" {
		return Note{}, errors.New("Empty note")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

```
- Code for the main package is as follows:
```go
package main

import (
	"bufio"
	"examples.com/NotetakingApp/note"
	"fmt"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Display User Note
	userNote.Display()

	//Save the file
	err = userNote.Save()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Saved Note successfully")

}

func getNoteData() (string, string) {
	title := getUserInput("Enter the Note Title")
	//if err != nil {
	//	fmt.Println(err)
	//	return "", "", err
	//}
	content := getUserInput("Enter the Note Content")
	//if err != nil {
	//	fmt.Println(err)
	//	return "", "", err
	//}

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	//Create a reader that listens on the command line
	reader := bufio.NewReader(os.Stdin)
	//rune in Go is a single character
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	//fmt.Scan(&value)
	//if value == "" {
	//	return "", errors.New("User input is empty")
	//}
	return text
}

```

## Interfaces and Generic Code
- Interfaces help us to write more flexible code
- ![img_48.png](img_48.png)
- Interfaces can be created anywhere in the Go file
- We need to use the "type" keyword
- An interface is contract that guarantees that a certain value, typically a Struct has a certain method.
- Interface dont have a function body, It just has a function name with its return value
- We also dont have to give the parameters of that function a name, just their type is required
- We can add as many method signatures to interface in Go
- If however, we have an interface with just one method, we have to name as the method name + "r" 
- For eg. an interface with just one method Save() inside it is named saver
```go
//Define an interface
type saver interface {
	Save() error
}

//Use the interface
func saveData(data saver) error {
err := data.Save()
if err != nil {
fmt.Println(err)
return err
}
return nil
}

//For our note taking app for userNote and todo we can implement it as follows

//Save the user note
err = saveData(userNote)
if err != nil {
return
}

//Save the todo
err = saveData(todo)
if err != nil {
return
}

```
- Note that in Go, the above code will automatically work for todos and notes
- This is because unlike other programming languages, we dont necessarily need to explicitly declare that a struct implements a particular interface
- Go automatically looks at the method signature and finds that if the given struct being passed to it, implements the methods defined by the interface, it automatically works
- ![img_49.png](img_49.png)
- This enables us to write more generic code automatically

### Embedded Interfaces
- We can create compound interfaces where we can combine other interfaces with other methods
- The only requirement for this is work automatically is that struct must implement all methods of that compound interface as well as any method defined in the embedded interface
```go
type outputtable interface {
	saver //Embedded interface
	Display()
	//DoSomething(int) string
}

type saver interface {
Save() error
}

func outputData(data outputtable) error {
data.Display()
return saveData(data)
}

func saveData(data saver) error {
err := data.Save()
if err != nil {
fmt.Println(err)
return err
}
return nil
}

//Now we can condense our code for userNote and todo as follows

//Output the userNote
err = outputData(userNote)
if err != nil {
fmt.Println(err)
return
}
fmt.Println("Saved Note successfully")

//Output the todo
err = outputData(todo)
if err != nil {
fmt.Println(err)
return
}
fmt.Println("Saved Todo successfully")
```
- ![img_50.png](img_50.png)

### The special "Any value allowed" type
- Go has a special "any" type which can accept any kind of input
- We can either specify it using "any" keyword or empty interface like interface {}
- Similar to "any" keyword in typescript
```go
//This can be dangerous though
printSomething("Some string")
printSomething(1)

// value can be of any type
// Even Println() accepts any type
func printSomething(value interface{}) {
	fmt.Println(value)
}
```
### Working with type Switches
- Similar to reflection in C#, we can use "type" keyword to get the type of "any" value that is being passed to the function
```go
func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case string:
		fmt.Println("String: ", value)
	case float64:
		fmt.Println("Float: ", value)
	default:
		fmt.Println("Unknown type: ", value)
	}
}

printSomething("Some string")
printSomething(1)
printSomething(1.5)
```
### Extracting Type Information From Values
- The previous value.(type) has a limitation that it can used only within a switch statement
- Alternative to that is to use value(.string) or value(.int)
- It returns the value and also an "ok" boolean which indicates if it is of the specified type
```go
func printSomething(value interface{}) {
intVal, ok := value.(int)
if ok {
intVal = intVal + 1
fmt.Println("Integer: ", intVal)
return
}
floatVal, ok := value.(float64)
if ok {
floatVal = floatVal + 1
fmt.Println("Float: ", intVal)
return
}

stringVal, ok := value.(string)
if ok {
fmt.Println("String: ", stringVal)
return
}
}
```

### Interfaces, Dynamic Types and Limitations
- Look at the following code:
- This is a lot of code to write, we need to check the type and return accordingly
- TO prevent this, Go offers Generics
```go
func add(a, b interface{}) any {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}

	aStr, aIsStr := a.(string)
	bStr, bIsStr := b.(string)

	if aIsStr && bIsStr {
		return aStr + bStr
	}

	return a
}
```
### Generics in Go
- In the above code, we saw how difficult it is to determine the type of value and do operations accordingly
- Go provides us with a generics feature which we can use like this using square brackets []
- Here we can also define that generics can be typeof int, float64 or string type
```go
func addGenerics[T int | float64 | string](a, b T) T {
	return a + b
}

//Go will figure out the return type of this function will be int
result := addGenerics(1, 2)
fmt.Println(result)

```
## Managing Related Data with Arrays, Slices and Maps
- Struct allows us to bundle different pieces of data together to use it as one single value
- Now we will look at collections
- ![img_51.png](img_51.png)
- These collections help us to group data together and manage lists of data
- ![img_53.png](img_53.png)
- An array that holds different values for the same thing
- We can create arrays of built in types as well as custom types like structs
- Even nested arrays are possible
```go
package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [4]string = [4]string{"1", "2", "3"}
	productNames[3] = "4"
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(productNames)
	fmt.Println(prices[2])
}

```
### Selecting parts of Array with Slices
- Slices allow us to create new lists based on other lists of data
```go
func main() {
	var productNames [4]string = [4]string{"1", "2", "3"}
	productNames[3] = "4"
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(productNames)
	fmt.Println(prices[2])

	//Start with the first element at the index 1(include element at index 1)
	//and go upto index 3 and exclude element at index 3, so this will have only 2 values
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
}
```
- We can do even more with slices like this
- Slices can be based on other slices also
```go
func main() {
	var productNames [4]string = [4]string{"1", "2", "3"}
	productNames[3] = "4"
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(productNames)
	fmt.Println(prices[2])

	//Start with the first element at the index 1(include element at index 1)
	//and go upto index 3 and exclude element at index 3, so this will have only 2 values
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

	//We can omit the starting index and write the ending index
	//This will start from the first index i.e 0 and go all the way upto index 3
	//and exclude the element at index 3
	newPrices := prices[:3]
	fmt.Println(newPrices)

	//Can do opposite also
	//Start with element at index 1 and go all the way upto the end element
	testPrices := prices[1:]
	fmt.Println(testPrices)

	//Using highest bound in a slice
	//Go has a special feature where we can go upto lastIndex + 1
	//This allows us to include the last element in the slice also(remember element at ending index is excluded)
	includeLastElementPrices := prices[1:4]
	fmt.Println(includeLastElementPrices)

	//Slices can be used on arrays, but they can also be based on other slices
	firstSlice := prices[1:]      //Returns [9.99, 45.99, 20.0]
	secondSlice := firstSlice[:1] //Returns [9.99]
	fmt.Println(secondSlice)
}
```
#### Slices are a reference/window into an array bit like a pointer
- When we create an array, that array is stored in memory
- When we then create a slice from that array, we get a window into that array
- So, if we modify an element in a slice, we would also modify the same element in the original array
```go
//Slices can be used on arrays, but they can also be based on other slices
	fmt.Println("Original Prices array: ", prices) //Returns [10.99 9.99 45.99 20]
	firstSlice := prices[1:]                       //Returns [9.99, 45.99, 20.0]
	fmt.Println("First Slice: ", firstSlice)
	secondSlice := firstSlice[:1] //Returns [9.99]
	fmt.Println("Second Slice: ", secondSlice)
	secondSlice[0] = 17.99
	fmt.Println("Second Slice after modification: ", secondSlice)
	fmt.Println("Original Prices array second time after modification: ", prices)
```
- ![img_54.png](img_54.png)
- When we create a slice, we dont create a copy of the original array, so we dont have that copy in memory which occupies extra space
- Instead, we still have only one array in memory and our slice is just a tiny reference to that part of that array
- Slicing is a very **memory efficient technique** to select parts of an array
- GO also saves some metadata for our slices, For every slice, we get length and capacity
- Length gives us the number of items in the array/slice
- Capacity 
```go
//Slices can be used on arrays, but they can also be based on other slices
	fmt.Println("Original Prices array: ", prices) //Returns [10.99 9.99 45.99 20]
	firstSlice := prices[1:]                       //Returns [9.99, 45.99, 20.0]
	fmt.Println("First Slice: ", firstSlice)
	secondSlice := firstSlice[:1] //Returns [9.99]
	fmt.Println("Second Slice: ", secondSlice)
	secondSlice[0] = 17.99
	fmt.Println("Second Slice after modification: ", secondSlice)
	fmt.Println("Original Prices array second time after modification: ", prices)
	fmt.Println("Length of first slice: ", len(firstSlice))
	fmt.Println("Capacity of first slice: ", cap(firstSlice))

```
- Look at this code
```go
//Slices can be used on arrays, but they can also be based on other slices
fmt.Println("Original Prices array: ", prices) //Returns [10.99 9.99 45.99 20]
firstSlice := prices[1:]                       //Returns [9.99, 45.99, 20.0]
fmt.Println("First Slice: ", firstSlice)
secondSlice := firstSlice[:1] //Returns [9.99]
fmt.Println("Second Slice: ", secondSlice)
secondSlice[0] = 17.99
fmt.Println("Second Slice after modification: ", secondSlice)
fmt.Println("Original Prices array second time after modification: ", prices)
fmt.Println("Length of first slice: ", len(firstSlice)) //Gives 3
fmt.Println("Capacity of first slice: ", cap(firstSlice)) //Gives 3
fmt.Println("Length of second slice: ", len(secondSlice)) //Gives 1
fmt.Println("Capacity of second slice: ", cap(secondSlice)) //Gives 3
	
```
- Capacity refers to the maximum number of elements the slice can hold without needing to reallocate.
- Length gives the number of elements currently present.
- Capacity gives the total space available for elements before reallocation becomes necessary.

- Refer to the following code:
```go
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
randomFirstSlice := prices[0:4]
fmt.Println("Random First Slice: ", randomFirstSlice)
fmt.Println("Length of random first slice: ", len(randomFirstSlice))
fmt.Println("Capacity of random first slice: ", cap(randomFirstSlice))
randomSecondSlice := randomFirstSlice[1:]
fmt.Println("Random Second Slice: ", randomSecondSlice)
fmt.Println("Length of random second slice: ", len(randomSecondSlice))
fmt.Println("Capacity of random second slice: ", cap(randomSecondSlice))
randomThirdSlice := randomFirstSlice[:1]
fmt.Println("Random Third Slice: ", randomThirdSlice)
fmt.Println("Length of random third slice: ", len(randomThirdSlice))
fmt.Println("Capacity of random third slice: ", cap(randomThirdSlice))
```
- ![img_55.png](img_55.png)
- Note the length of the third slice is 1, but its capacity is 4 because it is built off randomFirstSlice and it has 4 elements it can select from
- Also note, capacity starts from the first element to the last element in the slice off it is based on
- We can select more elements to the right rather than to the left
- A slice in Go is a view over an underlying array. It is defined by three components:
- A pointer to the start of the slice in the underlying array.
- The length (len), which determines the number of accessible elements.
- The capacity (cap), which determines how many elements the underlying array can hold starting from the slice’s pointer.

### Building Dynamic Lists with Slices
- Slices help us to create dynamic arrays
- We cannot always set the length of the array in advance
- In javascript, we always have dynamic arrays
- In Go, we cannot do that
- Solution in Go is by using slices
- When we create a slice, Go automatically creates an underlying array for us
- There is method append() to add elements to a slice and create a new slice or modify the existing slice
- There is no method to remove elements from a slice, since we can always slice a slice :)
```go
func main() {
	//Go automatically creates a slice for us and it also creates an underlying array
	//Here prices is a slice
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[1])
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	//Throws errors as we cannot access indexes that dont exist
	//prices[2] = 11.99

	//Append function not only adds the element to the slice but also to the underlying array
	updatedPrices := append(prices, 14.99)
	//Note the original slice doesnot change, to do that we have to do prices = append(prices,14.99)
	//Also note that the underlying array is automatically updated by Go
	//It is typical to work with slices in Go rather than use arrays
	//If we are sure about the number of the elements in the list, go for arrays
	fmt.Println(updatedPrices, prices)

	//Remove the elements
	//No function as such to remove elements since we already have slices
	prices = prices[1:]
	fmt.Println(prices)
}
```
## Practice Project
```go
// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
```
- Solution is as follows:
```go
type Product struct {
	title string
	id    string
	price float64
}

func main() {
	hobbies := [3]string{"Cricket", "Coding", "Azure"}
	fmt.Println("Hobbies: ", hobbies)
	fmt.Println("First element of Hobbies: ", hobbies[0])
	fmt.Println("Second and third element of Hobbies: ", hobbies[1:3])
	fmt.Println("First and third element of Hobbies: ", [2]string{hobbies[0], hobbies[2]})
	hobbiesSlice := hobbies[0:2]
	hobbiesSliceAnotherWay := hobbies[:2]
	fmt.Println(hobbiesSlice)
	fmt.Println(hobbiesSliceAnotherWay)
	hobbiesSlice = hobbies[1:]
	fmt.Println(hobbiesSlice)

	goals := []string{"Father", "Solution Architect"}
	fmt.Println(goals)
	goals[1] = "Azure Solution Architect"
	fmt.Println(goals)
	goals = append(goals, "Excellent Father")
	fmt.Println(goals)

	products := []Product{
		{title: "Simple Product 1", id: "1", price: 10},
		{title: "Simple Product 2", id: "2", price: 20},
	}
	fmt.Println(products)
	products = append(products, Product{
		title: "Simple Product 3",
		id:    "3",
		price: 30,
	})
	fmt.Println(products)
}
```

## Unpacking List Values
- Sometimes you have an existing Slice and you want to append another Slice or array to it.
- What if we want to add a list of float64 to another list of float64
- We cannot do it directly using the append() function as it can add only one value to a list of values
- So we use a special operator(similar to spread operator in Javascript) to spread open our second list and add its elements one by one to the first list

```go
prices := []float64{10.99, 8.99}
prices = prices[1:] //Gives 8.99
discountPrices := []float64{101.99, 80.99, 20.59}
prices = append(prices, discountPrices...)
fmt.Println(prices) //gives 8.99,101.99,80.99,20.59
```

## Introducing Maps
- Maps are used to group data together
- In Maps we have Key-Value pairs
- It is similar to Dictionary in C# or Map() in Javascript
```go
func MainMap() {
//websites := []string{"google.com", "aws.com", "azure.com"}
//fmt.Println(websites)
websitesMap := map[string]string{
"google": "google.com",
"aws":    "aws.com",
"azure":  "azure.com",
"github": "github.com",
}
fmt.Println(websitesMap["github"])
websitesMap["LinkedIn"] = "linkedin.com"
fmt.Println(websitesMap["LinkedIn"])

//Delete a key
delete(websitesMap, "LinkedIn")
fmt.Println(websitesMap)
}
```

### Maps vs Structs
- In Maps, we can use anything for keys
- We can even have an integer, array or struct as a key
- This gives us more flexibility
- In Structs, we have predefined data structures
- We cannot at runtime in structs add a key-value pair
- We also cannot delete a key-value pair from it
- Struct should be used when we have clearly defined kinds of data that should have exactly the same shape
- When we have a collection of values, and we want to have custom labels,map is our data structure of choice
- Think of maps as arrays where we dont use indexes, but any labels of our choice

### Using the special "make" function
- Using the make function we can define the length and capacity of the array we are making
- Note, that if the capacity is 5 and, we append the 6th element, i.e now the array length is 6, then capacity increases by 5
- So now length of array is 6 and capacity is 10
```go
userNames := make([]string, 2, 5)

	userNames[0] = "William Shakespeare"
	userNames[1] = "John"
	userNames = append(userNames, "Jane")
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Max")
	fmt.Println(userNames)
	fmt.Println(len(userNames))
	fmt.Println(cap(userNames))
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Max")
	fmt.Println(len(userNames))
	fmt.Println(cap(userNames))
```
### "make"ing Maps
- We can just specify the length while using the make() function and not the capacity
- This allows Go to pre-allocate memory for the map
```go
//For maps, we can use make function and just specify the length of the map so that Go can pre-allocate memory
	//This makes it a bit more efficient
	courseRatings := make(map[string]float64, 5)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	fmt.Println(courseRatings)
```
### Working with type aliases
- For this code: map[string]float64, we have to type a lot.
- We can make it shorter by using type aliases like this
```go
type floatMap map[string]float64

func (f floatMap) output() {
	fmt.Println(f)
}

//courseRatings := make(map[string]float64, 5)
courseRatings := make(floatMap, 5)
courseRatings["go"] = 4.7
courseRatings["react"] = 4.8
courseRatings.output()
```

### For loops with Arrays, Slices and Maps
- We always have a requirement to iterate over arrays, slices and maps
- This is similar to "foreach" keyword in C#, we have an equivalent keyword in Go called "range"
```go
courseRatings := make(floatMap, 5)
courseRatings["go"] = 4.7
courseRatings["react"] = 4.8

userNames := make([]string, 2, 5)
userNames[0] = "William Shakespeare"
userNames[1] = "John"

//If we dont care about individual item values or indexes, we can also just write for range userNames
	for index, value := range userNames {
		//fmt.Println(index)
		//fmt.Println(value)
		fmt.Println(userNames[index])
		fmt.Println(value)
		fmt.Println(value == userNames[index])
	}

	for key := range courseRatings {
		//fmt.Println(key)
		//fmt.Println(value)
		fmt.Println(courseRatings[key])
	}
```

## Functions (Deep Dive)
- We will explore advanced function related features
- ![img_56.png](img_56.png)
- Functions are first class values in Go
- We can use functions themselves as parameter values for other functions
- This is similar to delegates in C# or Action or Func keyword in C#
- We can also do type aliasing of functions
```go
package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

```
### Returning Functions as values
- We can return functions as values from the result of a function
```go
func getTransformerFunction() transformFn {
	return double
}

func getTransformerFunctionWithNumbers(numbers *[]int) transformFn {
if (*numbers)[0] == 1 {
return double
} else {
return triple
}
}

numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
moreNumbers := []int{11, 12, 13, 14}

transformerFn1 := getTransformerFunctionWithNumbers(&numbers)
transformerFn2 := getTransformerFunctionWithNumbers(&moreNumbers)

transformedNumbers := transformNumbers(&numbers, transformerFn1)
moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)
fmt.Println(transformedNumbers)
fmt.Println(moreTransformedNumbers)
```

### Anonymous functions
- When we have a case where a function wants another function as parameter value, we can use anonymous functions
- We can define a function when we need it rather than defining it in advance
```go
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	transformed := transformNumbers(&numbers, func(num int) int {
		return num * 2
	})

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

```
- This makes the code more readable

### Understanding Closures
- Closures also use anonymous functions
- They use a specific aspect of anonymous function
- Child scope can access parameters from the parent scope
- Similar to javascript
```go
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

//Usage
double := createTransformer(2)
triple := createTransformer(3)
fmt.Println(double(4))
fmt.Println(triple(5))

doubled := transformNumbers(&numbers, double)
fmt.Println(doubled)
tripled := transformNumbers(&numbers, triple)
fmt.Println(tripled)

```
- In Go, a closure is a function value that references variables from outside its own body. 
- These functions can “close over” variables in their surrounding scope, which means they remember and can manipulate the values of those variables even after the surrounding function has exited. 
- This is a powerful concept that allows for more dynamic and flexible programming
```go
package main

import "fmt"

func main() {
    // Create a function that returns a closure
    adder := func(start int) func(int) int {
        sum := start
        return func(x int) int {
            sum += x
            return sum
        }
    }

    // Initialize the closure with a starting value
    myAdder := adder(10)

    // Use the closure
    fmt.Println(myAdder(5))  // Output: 15
    fmt.Println(myAdder(3))  // Output: 18
    fmt.Println(myAdder(10)) // Output: 28
}
```
- ![img_57.png](img_57.png)

### Recursion
- It is a function that calls itself
- A good example can be a function that calculates factorial of a number
```go
func factorial(num int) int {
	//exit condition
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
```
### Using Variadic Functions
- It is actually quite a simple concept
- It is similar to params keyword in C#
- It takes a list of numbers and converts it into a slice of the defined datatype for us automatically
- Also it is like rest operator in javascript
```go
func main() {
	//numbers := []int {1,10,15}
	
	sum := sumup(1,10,15)
	fmt.Println(sum)
}

func sumup(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
```
- ![img_58.png](img_58.png)
- ![img_59.png](img_59.png)

### Splitting slices into Parameter Values
- We can use the spread operator here
- If we have a slice of numbers, we can pass it to a variadic function by using the spread operator for the slice: numbers defined below
```go
func main() {
	numbers := []int{1, 10, 15}
	sum := sumup("testing", 2, numbers...)
	fmt.Println(sum)
}
func sumup(prestarting string, startingValue int, numbers ...int) int {
	fmt.Println(prestarting)
	sum := startingValue
	for _, number := range numbers {
		sum += number
	}
	return sum
}
```

## Project: Price Calculator
- ![img_60.png](img_60.png)
- ![img_61.png](img_61.png)
- In Go, when you access a struct field through a pointer, Go automatically dereferences the pointer for you
- You don't need to manually dereference the struct. The language simplifies this for convenience. 
- You can still explicitly dereference the pointer 
- We have bufio package which can help us 
- The bufio package in Go is used for buffered I/O operations, which improve efficiency when reading and writing data. It wraps around an io.Reader or io.Writer to provide buffering, reducing the number of direct system calls and improving performance.
- Efficient File Reading/Writing: Allows reading and writing large files in chunks rather than one byte at a time.
- Buffered Input Handling: Useful for handling input streams like user input or network connections more effectively
- Passing a receiver pointer (*T) instead of a receiver struct (T) in Go is often necessary due to these key reasons:
#### Mutability – Modify Struct Fields
- When you pass a struct by value, Go creates a copy, meaning any modifications won't affect the original. 
- Using a pointer receiver allows methods to change the original struct.
```go
package main

import "fmt"

type User struct {
	Name string
}

// Method with value receiver (does NOT modify original struct)
func (u User) UpdateName(newName string) {
	u.Name = newName
}

// Method with pointer receiver (modifies original struct)
func (u *User) UpdateNamePtr(newName string) {
	u.Name = newName
}

func main() {
	user := User{Name: "Alice"}

	user.UpdateName("Bob")   // Value receiver—modifications lost
	fmt.Println(user.Name)   // Output: Alice (no change)

	user.UpdateNamePtr("Bob") // Pointer receiver—modifications persist
	fmt.Println(user.Name)    // Output: Bob
}
```
#### Efficiency – Avoiding Struct Copying
- Passing large structs by value causes unnecessary copying, which wastes memory and CPU cycles. 
- Using a pointer ensures only a reference is passed.
```go
type LargeStruct struct {
Data [1000]int
}

// Using a pointer avoids copying 1000 integers
func (ls *LargeStruct) ProcessData() {
ls.Data[0] = 42 // Modify data directly
}
```

#### Consistency – Pointer for All Methods
- If some methods use a pointer receiver and others don’t, Go treats them differently—leading to unexpected behavior. 
- Using only pointer receivers ensures consistency.

#### We implemented this in our project like this
```go
package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	fmt.Println(result)
	//return result
}

// NewTaxIncludedPriceJob: Constructor Function
// NewTaxIncludedPriceJob: In Go, when you access a struct field through a pointer, Go automatically dereferences the pointer for you
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Error loading prices.txt")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))
	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println(err)
			file.Close()
			return
		}
		prices[lineIndex] = floatPrice
	}

	//fmt.Println(job.InputPrices)
	job.InputPrices = prices
}

```
- We called this package from our main package as follows:
```go
package main

import (
	"examples.com/price-calculator-project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}

```
### Reading and Writing JSON to Files
```go
package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open("prices.txt")
	if err != nil {
		return nil, errors.New("Error opening file: " + err.Error())
	}

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Error scanning file: " + err.Error())
	}

	file.Close()
	return lines, nil
}

func WriteJson(data interface{}, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("Error creating file: " + err.Error())
	}

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Error encoding file: " + err.Error())
	}

	file.Close()
	return nil
}

```

- If we are outputting a result to JSON and we want to exclude some field from the struct, we use the "-" sign
```go
type TaxIncludedPriceJob struct {
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]float64      `json:"tax_included_prices"`
	IOManager         filemanager.FileManager `json:"-"`
}
```
- In the above code, we are excluding the IO Manager

### Working on a swappable struct
- We can want to use more generic, more reusable code
- This can be done by using interfaces
- We can create a separate file called iomanager.go and have this code in there
```go
package iomanager

type IOManager interface {
	WriteResult(data interface{}) error
	ReadLines() ([]string, error)
}

```
- Note this file has the same method signatures are those found in commandManager.go file and fileManager.go file 
- These files are as follows:
- commandManager.go
```go
package cmdmanager

import "fmt"

type CMDManager struct {
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices, Confirm every price with enter")
	var prices []string
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (cmd CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}

```
- FileManager.go is as follows:
```go
package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm *FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("Error opening file: " + err.Error())
	}

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Error scanning file: " + err.Error())
	}

	file.Close()
	return lines, nil
}

func (fm *FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("Error creating file: " + err.Error())
	}

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Error encoding file: " + err.Error())
	}

	file.Close()
	return nil
}

func New(inputFilePath string, outputFilePath string) *FileManager {
	return &FileManager{
		inputFilePath,
		outputFilePath}
}

```
- Next step is that in prices.go file, we need to use the generic interface: ioManager instead of fileManager struct
```go
type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]float64  `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"`
}

//Modify the constructor also
func NewTaxIncludedPriceJob(taxRate float64, fm iomanager.IOManager) *TaxIncludedPriceJob {

return &TaxIncludedPriceJob{
TaxRate:   taxRate,
IOManager: fm,
}
}
```
- Now to swap them automatically we can do something like this in main.go file
```go
package main

import (
	"examples.com/price-calculator-project/filemanager"
	"fmt"

	//"examples.com/price-calculator-project/filemanager"
	"examples.com/price-calculator-project/prices"
	//"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("prices_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, fm)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not process price job")
			fmt.Println(err)
		}
	}
}

```
- As you can see interfaces are implemented implicitly in Go.
- We can swap them out with either FileManager or Command Manager as and when we want
- 



