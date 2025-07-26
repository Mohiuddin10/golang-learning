package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64
	fmt.Print("Enter the amount you want to invest: ")
	fmt.Scan(&investmentAmount)

	var interestRate float64
	fmt.Print("Enter the expected annual interest rate (in %): ")
	fmt.Scan(&interestRate)

	var years float64
	fmt.Print("Enter the number of years you want to invest for: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, interestRate, years)

	// var futureValue = investmentAmount * math.Pow(1+interestRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Printf("Future Value of Investment: %.2f\n", futureValue)
	fmt.Printf("Future Real Value of Investment: %.2f\n", futureRealValue)

	c := add(5, 10)
	fmt.Printf("Sum: %d\n", c)
}

// ===> Return multiple values from one func

func calculateFutureValue(investmentAmount, interestRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+interestRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

// alternative return value syntex
func add(a, b int) (sum int) {
	sum = a + b
	return
}
