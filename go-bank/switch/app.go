package main

import (
	"fmt"
	"os"
)

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	accountBalance := 1000.0
	fmt.Println("====Welcome to Go Bank====")

	for {

		fmt.Println("Please enter your service code to continue...")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

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
