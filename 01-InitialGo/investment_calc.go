package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var principalCapital, expectedAppreciation, termOfInvestment float64
	// var expectedAppreciation float64 = 5.5

	fmt.Print("Please enter your intial captial for investement: ")
	fmt.Scan(&principalCapital)

	fmt.Print("Please enter your expected interest rate: ")
	fmt.Scan(&expectedAppreciation)

	fmt.Print("Please enter how long your investment will appreciate for: ")
	fmt.Scan(&termOfInvestment)

	expectedReturn := principalCapital * math.Pow((1+expectedAppreciation/100), termOfInvestment)
	inflationAdjustedReturn := expectedReturn / math.Pow((1+inflationRate/100), termOfInvestment)

	// fmt.Println("Expected Return : ", expectedReturn)
	fmt.Printf("Expected Return : %.2f\n Inflation adjusted value : %.2f", expectedReturn, inflationAdjustedReturn)
	// fmt.Println("Inflation adjusted Return : ", inflationAdjustedReturn)
}
