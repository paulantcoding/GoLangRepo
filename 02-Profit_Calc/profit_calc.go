package main

import (
	"fmt"
)

func main() {
	var revenueTotal, expensesTotal, taxRate float64
	revenueTotal = getUserInput("Please enter the current total revenue : ")
	expensesTotal = getUserInput("Please enter the current total expenses : ")
	taxRate = getUserInput("Please enter the current tax rate : ")

	earningsBeforeTax, earningsAfterTax, currentRatio := calculateEarnings(revenueTotal, expensesTotal, taxRate)
	fmt.Println("Earning before Tax :", earningsBeforeTax)
	fmt.Println("Current profit :", earningsAfterTax)
	fmt.Println("Current Ratio :", currentRatio)

}

// retrieve information from user
func getUserInput(text string) (uiEntry float64) {
	fmt.Print(text)
	fmt.Scan(&uiEntry)
	return uiEntry
}

// calculates and returns earning before tax
// earning after tax
// and current ratio
func calculateEarnings(revenueTotal, expensesTotal, taxRate float64) (ebt, eat, cr float64) {
	ebt = revenueTotal - expensesTotal
	eat = ebt * (1 - taxRate/100)
	cr = ebt / eat
	return
}
