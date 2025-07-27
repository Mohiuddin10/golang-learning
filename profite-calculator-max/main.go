package main

import "fmt"

func main() {

	ravenue := getUserInput("Enter your total revenue: ")
	expenses := getUserInput("Enter your total expenses: ")
	taxRate := getUserInput("Enter your tax rate (in %): ")

	// earningsBeforeTax := ravenue - expenses
	// taxAmount := earningsBeforeTax * (taxRate / 100)
	// earningAfterTax := earningsBeforeTax - taxAmount
	// ratio := earningsBeforeTax / earningAfterTax

	earningsBeforeTax, earningAfterTax, ratio := calculateProfit(ravenue, expenses, taxRate)

	fmt.Println("Earning Before Tax: ", earningsBeforeTax)
	fmt.Println("Profit: ", earningAfterTax)
	fmt.Println("Ratio: ", ratio, "%")

	fmt.Printf("EBT: %0.0f\nProfit: %0.0f\nRatio: %0.0f\n", earningsBeforeTax, earningAfterTax, ratio)

	// formatted string   => That put a string + value in a variable
	formattedEbt := fmt.Sprintf("Formatted EBT: %.0f\n", earningsBeforeTax)
	formattedProfit := fmt.Sprintf("Formatted Profit: %.0f\n", earningAfterTax)
	formattedRatio := fmt.Sprintf("Formatted Ratio: %.0f\n", ratio)

	fmt.Print(formattedEbt, formattedProfit, formattedRatio)

	// ===> If we want to use multiline string we can use `` but at that time \n will nor work

}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateProfit(ravenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := ravenue - expenses
	earningAfterTax := earningsBeforeTax - (1 - taxRate/100)
	ratio := earningsBeforeTax / earningAfterTax
	return earningsBeforeTax, earningAfterTax, ratio
}
