package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) display() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Paul"
	userNames[1] = "Flynn"
	userNames = append(userNames, "Iola")
	userNames = append(userNames, "Mylar")

	fmt.Println(userNames)

	// courseRatings := make(map[string]float64, 3)
	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["rust"] = 3.8
	courseRatings["springboot"] = 4.8

	// fmt.Println(courseRatings)
	courseRatings.display()

	for index, value := range userNames {
		fmt.Println("Index :", index)
		fmt.Println("Value :", value)
	}

	for course, rating := range courseRatings {
		fmt.Println("Course :", course)
		fmt.Println("Rating :", rating)
	}

}
