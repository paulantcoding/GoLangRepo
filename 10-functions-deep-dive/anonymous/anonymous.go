package anonymous

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}

	double := createTransformer(2)
	triple := createTransformer(3)

	//anonymous function
	transformedNumbers := transformNumber(&numbers, func(number int) int { return number * 2 })

	doubled := transformNumber(&numbers, double)
	tripled := transformNumber(&numbers, triple)

	fmt.Println("Transformed numbers :", transformedNumbers)
	fmt.Println("Double numbers :", doubled)
	fmt.Println("Triple numbers :", tripled)
}

func transformNumber(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

//closures
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
