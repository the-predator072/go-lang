package main

import "fmt"

type floatMap map[string]float64
type Array []string

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)
	// userNames := []string{}

	userNames[0] = "Julie"
	userNames[1] = "Vikas"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	courseRatings.output()
	// for loop for arrays and slices
	for index, value := range userNames {
		fmt.Println(index, value)
	}
	// for loop for maps
	for key, value := range courseRatings {
		fmt.Println(key, value)
	}

}
