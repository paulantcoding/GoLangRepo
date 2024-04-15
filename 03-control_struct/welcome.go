package main

import "fmt"

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
