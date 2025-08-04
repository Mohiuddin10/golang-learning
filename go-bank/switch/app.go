package main

import (
	"fmt"

	"example.com/switch/fileOps"
)

const accountFile = "balance.txt"

func main() {
	accountBalance, err := fileOps.GetFloatNumberFromFile(accountFile)
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
				fileOps.WriteFloatToFile(accountBalance, accountFile)
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
				fileOps.WriteFloatToFile(accountBalance, accountFile)
				fmt.Printf("You have successfully withdrawn $%.2f. New balance is: $%.2f\n", WithdrawAmount, accountBalance)
			}
		default:
			fmt.Println("Exiting the service. Thank you!")
			fmt.Println("Thank you for using Go Bank. Goodbye!")
			return
		}

	}

}
