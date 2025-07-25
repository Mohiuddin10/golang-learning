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
}
