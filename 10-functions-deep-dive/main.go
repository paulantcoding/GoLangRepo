package main

import "fmt"

//variadic functions
func main() {

	numbers := []int{10, 1, 15, 40, -5}
	sum := sumup(1, numbers...)

	anotherSum := sumup(1, 10, 15)

	fmt.Println(sum)
	fmt.Println(anotherSum)

}

func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum

}
