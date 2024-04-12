package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var principalCapital, expectedAppreciation, termOfInvestment float64
	// var expectedAppreciation float64 = 5.5

	fmt.Print("Please enter your intial captial for investement: ")
	fmt.Scan(&principalCapital)

	fmt.Print("Please enter your expected interest rate: ")
	fmt.Scan(&expectedAppreciation)

	fmt.Print("Please enter how long your investment will appreciate for: ")
	fmt.Scan(&termOfInvestment)

	// expectedReturn := principalCapital * math.Pow((1+expectedAppreciation/100), termOfInvestment)
	// inflationAdjustedReturn := expectedReturn / math.Pow((1+inflationRate/100), termOfInvestment)
	expectedReturn, inflationAdjustedReturn := calculateValuesReturn(principalCapital, expectedAppreciation, termOfInvestment)

	formattedExpectedReturn := fmt.Sprintf("Expected Return : %.2f\n", expectedReturn)
	formattedIAReturn := fmt.Sprintf("Inflation adjusted value : %.2f", inflationAdjustedReturn)

	fmt.Print(formattedExpectedReturn)
	fmt.Print(formattedIAReturn)
	// fmt.Println("Expected Return : ", expectedReturn)
	//	fmt.Printf("Expected Return : %.2f\n Inflation adjusted value : %.2f", expectedReturn, inflationAdjustedReturn)
	// fmt.Println("Inflation adjusted Return : ", inflationAdjustedReturn)
}

// func calculateValuesReturn(principalCapital, expectedAppreciation, termOfInvestment float64) (float64, float64) {
// 	returningValue := principalCapital * math.Pow((1+expectedAppreciation/100), termOfInvestment)
// 	returnedInflationValue := returningValue / math.Pow((1+inflationRate/100), termOfInvestment)
// 	return returningValue, returnedInflationValue
// }

func calculateValuesReturn(principalCapital, expectedAppreciation, termOfInvestment float64) (returningValue float64, returnedInflationValue float64) {
	returningValue = principalCapital * math.Pow((1+expectedAppreciation/100), termOfInvestment)
	returnedInflationValue = returningValue / math.Pow((1+inflationRate/100), termOfInvestment)
	return
}
