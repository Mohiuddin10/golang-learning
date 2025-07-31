package main

import (
	"errors"
	"fmt"
	"os"
)

func handleError(err error) {
	fmt.Println("Error:", err)
}

func main() {
	readStoredValues()
	ravenue, err := getUserInput("Enter your total revenue: ")
	if err != nil {
		handleError(err)
		return
	}
	expenses, err := getUserInput("Enter your total expenses: ")
	if err != nil {
		handleError(err)
		return
	}
	taxRate, err := getUserInput("Enter your tax rate (in %): ")
	if err != nil {
		handleError(err)
		return
	}

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
	storeCalculatedValue(formattedEbt, formattedProfit, formattedRatio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("input must be greater than zero")

	} else {
		return userInput, nil
	}
}

func calculateProfit(ravenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := ravenue - expenses
	earningAfterTax := earningsBeforeTax - (1 - taxRate/100)
	ratio := earningsBeforeTax / earningAfterTax
	return earningsBeforeTax, earningAfterTax, ratio
}

func storeCalculatedValue(formattedEbt, formattedProfit, formattedRatio string) {
	ebt := fmt.Sprint(formattedEbt)
	profit := fmt.Sprint(formattedProfit)
	ratio := fmt.Sprint(formattedRatio)

	// Here you can store the values in a database or a file
	os.WriteFile("calculated_values.txt", []byte(ebt+"\n"+profit+"\n"+ratio), 0644)
}

func readStoredValues() {
	data, err := os.ReadFile("calculated_values.txt")
	if err != nil {
		fmt.Println("Error reading stored values:", err)

	}
	fmt.Println("Stored Values:\n", string(data))
}
