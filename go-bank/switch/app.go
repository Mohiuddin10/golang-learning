package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountFile = "balance.txt"

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountFile)
	if err != nil {
		return 1000, errors.New("could not read balance file, initializing with default balance")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("could not parse balance, initializing with default balance")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountFile, []byte(balanceText), 0644)
}

func main() {
	accountBalance, err := readBalanceFromFile()
	if err != nil {
		fmt.Println("Error reading balance:", err)
	}
	fmt.Println("====Welcome to Go Bank====")

	for {
		presentMenu()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		fmt.Println("You selected option:", choice)

		switch choice {
		case 1:
			fmt.Printf("Your current balance is: $%.2f\n", accountBalance)
		case 2:
			var dipositeAmount float64
			fmt.Print("Enter the amount to deposit: ")
			fmt.Scan(&dipositeAmount)
			if dipositeAmount <= 0 {
				fmt.Println("Deposit amount must be greater than zero.")
				continue
			} else {
				accountBalance += dipositeAmount
				writeBalanceToFile(accountBalance)
			}

			fmt.Printf("You have successfully deposited $%.2f. New balance is: $%.2f\n", dipositeAmount, accountBalance)
		case 3:
			var WithdrawAmount float64
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&WithdrawAmount)
			if WithdrawAmount <= 0 {
				fmt.Println("Withdrawal amount must be greater than zero & less than account balance.")
				continue
			} else if WithdrawAmount > accountBalance {
				fmt.Println("Insufficient balance for this withdrawal.")
			} else {
				accountBalance -= WithdrawAmount
				writeBalanceToFile(accountBalance)
				fmt.Printf("You have successfully withdrawn $%.2f. New balance is: $%.2f\n", WithdrawAmount, accountBalance)
			}
		default:
			fmt.Println("Exiting the service. Thank you!")
			fmt.Println("Thank you for using Go Bank. Goodbye!")
			return
		}

	}

}
