package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

func main() {
	var loopRef int
	// var balanceVal float64 = 2500
	loopRef = 1
	fmt.Println("Welcome to your banking app")
	for loopRef != 4 {
		desiredAction := welcomeMessage()
		performAction(desiredAction)
		loopRef = desiredAction
	}
}

func welcomeMessage() (uiInput int) {
	fmt.Println("Which operation would you like to perform?")
	fmt.Println("(Please enter the number corresponding to your desired action)")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Println("Your choice :")
	fmt.Scan(&uiInput)
	return uiInput
}

func performAction(actionNumber int) {
	var balanceVal, err = getBalance()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("***********************")
		//panic("Failure had arose, for security we have terminated the process")
	}
	switch actionNumber {
	case 1:
		//display balance
		fmt.Printf("Your current balance is : %0.2f\n", balanceVal)
	case 2:
		//deposit money
		var depositAmount float64
		fmt.Println("How much would you like to deposit? :")
		fmt.Scan(&depositAmount)
		if depositAmount <= 0 {
			fmt.Print("Your have entered an invalid amount, please retry with a positive a mount")
		} else {
			balanceVal += depositAmount
			fmt.Printf("Your updated balance is : %0.2f\n", balanceVal)
			writeBalance(balanceVal)
		}
	case 3:
		//withdraw money
		var withdrawAmount float64
		fmt.Println("How much would you like to withdraw? :")
		fmt.Scan(&withdrawAmount)
		if withdrawAmount > balanceVal || withdrawAmount <= 0 {
			fmt.Println("Your request is invalid, you have entered too high an amount based on your balance or a negative number, this action cannot be processed")
		} else {
			balanceVal -= withdrawAmount
			fmt.Printf("Your updated balance is : %0.2f\n", balanceVal)
			writeBalance(balanceVal)
		}
	case 4:
		//display balance
		fmt.Println("Exiting your app, have a nice day")
	default:
		fmt.Println("The choise you selected is not identified, please retry")
	}

}

func writeBalance(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}

func getBalance() (float64, error) {
	data, err := os.ReadFile(balanceFile)

	if err != nil {

		return 1000, errors.New("Failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value")
	}
	return balance, nil
}
