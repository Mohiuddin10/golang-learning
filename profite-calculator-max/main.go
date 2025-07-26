package main

import "fmt"

func main() {
	var ravenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter your total revenue: ")
	fmt.Scan(&ravenue)

	fmt.Print("Enter your total expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter your tax rate (in %): ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := ravenue - expenses
	taxAmount := earningsBeforeTax * (taxRate / 100)
	earningAfterTax := earningsBeforeTax - taxAmount
	ratio := earningsBeforeTax / earningAfterTax

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
