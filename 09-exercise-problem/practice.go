package main

import "fmt"

// 7
type product struct {
	title string
	id    string
	price float64
}

func main() {
	//1
	hobbies := [3]string{"coding", "drawing", "gaming"}
	fmt.Println("Hobbies : ", hobbies)

	//2
	initialHobby := hobbies[0]
	fmt.Println("First Hobby:", initialHobby)

	hobbyList := hobbies[1:]
	fmt.Println("Secondary hobby list :", hobbyList)

	//3
	initialHobbySlice := hobbies[:2]
	secondaryHobbySlice := hobbies[0:2]

	fmt.Println("initial hobby slice :", initialHobbySlice)
	fmt.Println("secondary hobby slice :", secondaryHobbySlice)

	//4
	tertiaryHobbySlice := initialHobbySlice[1:3]
	fmt.Println("tertiary hobby slice :", tertiaryHobbySlice)

	//5
	goals := []string{"Learn more", "Be more productive"}

	//6
	goals[1] = "Read more"

	goals = append(goals, "make more time for yourself")

	fmt.Println("Goals:", goals)

	//7
	productList := []product{}

	gameproduct := product{
		title: "Playstation",
		id:    "abc-1923",
		price: 435.99,
	}

	readingproduct := product{
		title: "Chainsaw man vol15",
		id:    "csm-2023",
		price: 12.99,
	}

	drawingproduct := product{
		title: "parchment sketchbook",
		id:    "psb-2023",
		price: 25.99,
	}

	productList = append(productList, gameproduct)
	productList = append(productList, readingproduct)
	productList = append(productList, drawingproduct)

	fmt.Println("Product List:", productList)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
