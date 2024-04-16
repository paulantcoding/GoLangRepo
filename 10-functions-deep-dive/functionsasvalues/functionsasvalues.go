package functionsasvalues

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 6, 4, 2}
	// doubledNumbers := transformNumber(&numbers, double)
	// tripledNumbers := transformNumber(&numbers, triple)
	// fmt.Println("Doubled Numbers :", doubledNumbers)
	// fmt.Println("Tripled Numbers :", tripledNumbers)

	trasformerFn1 := getTransformerFunction(&numbers)
	trasformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumber(&numbers, trasformerFn1)
	moretransformedNumbers := transformNumber(&moreNumbers, trasformerFn2)

	fmt.Println("Transgomed numbers :", transformedNumbers)
	fmt.Println("Transgomed numbers :", moretransformedNumbers)

}

func transformNumber(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
