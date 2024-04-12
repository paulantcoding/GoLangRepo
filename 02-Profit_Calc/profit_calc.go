package main

import (
	"errors"
	"fmt"
	"os"
)

const profitFile = "profitFile.txt"

// validate user input
//show error message and exit if invalid
//no negative or 0 numners
//store results into file

func main() {
	var revenueTotal, expensesTotal, taxRate float64
	revenueTotal, err := getUserInput("Please enter the current total revenue : ")
	if err != nil {
		fmt.Print(err)
		return
	}
	expensesTotal, err2 := getUserInput("Please enter the current total expenses : ")
	if err2 != nil {
		fmt.Print(err2)
		return
	}
	taxRate, err3 := getUserInput("Please enter the current tax rate : ")
	if err3 != nil {
		fmt.Print(err3)
		return
	}

	earningsBeforeTax, earningsAfterTax, currentRatio := calculateEarnings(revenueTotal, expensesTotal, taxRate)
	saveResultsInFile(earningsBeforeTax, earningsAfterTax, currentRatio)
	fmt.Println("Earning before Tax :", earningsBeforeTax)
	fmt.Println("Current profit :", earningsAfterTax)
	fmt.Println("Current Ratio :", currentRatio)

}

// retrieve information from user
func getUserInput(text string) (float64, error) {
	var uiEntry float64
	fmt.Print(text)
	fmt.Scan(&uiEntry)
	if uiEntry <= 0 {
		return 0, errors.New("Value must be higher than 0")
	}
	return uiEntry, nil
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

func saveResultsInFile(ebt, eat, cr float64) {
	resultsToText := fmt.Sprintf("EBT: %.2f\nEAT: %.2f\nRatio %.2f\n", ebt, eat, cr)
	os.WriteFile(profitFile, []byte(resultsToText), 0644)
}
