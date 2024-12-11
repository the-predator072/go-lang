package main

import "fmt"

func main() {
	// DYNAMIC ARRAYS WITH SLICES
	prices := []float64{10.99, 23.24}
	fmt.Println(prices)
	prices[1] = 11.4
	// append returns new array
	prices = append(prices, 4.5)
	fmt.Println(prices)
	// shift
	prices = prices[1:]
	fmt.Println(prices)
}

//             STATIC ARRAYS AND SLICES
// func main() {
// 	// float array
// 	prices := [4]float64{10.0, 11.1, 2.2, 5.4}
// 	// string array
// 	products := [5]string{"vikas", "negi"}
// 	fmt.Println(prices)
// 	fmt.Println(products)
// 	//slicing the array
// 	slicedPrices := prices[1:3]
// 	fmt.Println(slicedPrices)
// 	slicedPrices1 := prices[:3]
// 	fmt.Println(slicedPrices1)
// 	slicedPrices2 := prices[2:]
// 	fmt.Println(slicedPrices2)
// 	//slices are basically sub arrays i.e., we can create create slice from another slice
// 	//when we make slice we don't have copy of the original array any changes to slice will reflect in the original array
// 	fmt.Println(len(prices))
// 	fmt.Println(len(slicedPrices2))
// 	fmt.Println(cap(prices))
// 	fmt.Println(cap(slicedPrices2))
// }
