package main

import (
	"fmt"
)

func main() {
	var revenueTotal, expensesTotal, taxRate float64

	fmt.Print("Please enter the current totatl revenue : ")
	fmt.Scan(&revenueTotal)

	fmt.Print("Please enter the current total expenses : ")
	fmt.Scan(&expensesTotal)

	fmt.Print("Please enter the current tax rate : ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenueTotal - expensesTotal

	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)

	currentRatio := earningsBeforeTax / earningsAfterTax

	fmt.Println("Earning before Tax :", earningsBeforeTax)
	fmt.Println("Current profit :", earningsAfterTax)
	fmt.Println("Current Ratio :", currentRatio)

}
