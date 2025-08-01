package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	readStoredValues()
	ravenue, err1 := getUserInput("Enter your total revenue: ")
	// if err != nil {
	// 	handleError(err)
	// 	return
	// }
	expenses, err2 := getUserInput("Enter your total expenses: ")
	// if err != nil {
	// 	handleError(err)
	// 	return
	// }
	taxRate, err3 := getUserInput("Enter your tax rate (in %): ")
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Error:", err1)
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
	storeCalculatedValue(earningsBeforeTax, earningAfterTax, ratio)
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

func storeCalculatedValue(ebt, profit, ratio float64) {
	result := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f/n", ebt, profit, ratio)

	// Here you can store the values in a database or a file
	os.WriteFile("calculated_values.txt", []byte(result), 0644)
}

func readStoredValues() {
	data, err := os.ReadFile("calculated_values.txt")
	if err != nil {
		fmt.Println("Error reading stored values:", err)

	}
	fmt.Println("Stored Values:\n", string(data))
}
