package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	fmt.Print("Enter the amount you want to invest: ")
	fmt.Scan(&investmentAmount)

	var interestRate float64
	fmt.Print("Enter the expected annual interest rate (in %): ")
	fmt.Scan(&interestRate)

	var years float64
	fmt.Print("Enter the number of years you want to invest for: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+interestRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future Value of Investment: ", futureValue)
	fmt.Println("Future Real Value of Investment: ", futureRealValue)
}
