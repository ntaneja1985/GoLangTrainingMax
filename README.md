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
