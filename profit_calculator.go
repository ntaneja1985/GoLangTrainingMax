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
