package main

import (
	"fmt"
)

func main() {
	result := add(1, 2.5)

	fmt.Println(result)
}

func add[T int | float64 | float32](a, b T) T {
	return a + b
}
