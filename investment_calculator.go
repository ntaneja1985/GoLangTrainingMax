package main

import "fmt"
import "math"

// We can keep some variable declarations at the file level also called Global variables
// However, we cannot use the shortcut ( :=) while defining variables and would need to use "var" keyword
const inflationRate float64 = 6.5

func main() {
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

	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
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

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	//return fv, rfv
	return
}
