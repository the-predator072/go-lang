package main

import "fmt"

type transformFn func(int) int

func main() {
	slice := []int{1, 2, 3, 4, 6}
	slice2 := []int{4, 2, 4, 1, 4, 5}
	fmt.Println(slice)
	doubled := transformSlice(&slice, double)
	fmt.Println(doubled)
	tripled := transformSlice(&slice2, triple)
	fmt.Println(tripled)
	transformer1 := getTransformerFn(&slice)
	transformer2 := getTransformerFn(&slice2)
	transformedSlice1 := transformSlice(&slice, transformer1)
	transformedSlice2 := transformSlice(&slice2, transformer2)
	fmt.Println(transformedSlice1)
	fmt.Println(transformedSlice2)
}

func transformSlice(s *[]int, transform transformFn) []int {
	dNumber := []int{}

	for _, value := range *s {
		dNumber = append(dNumber, transform(value))
	}
	return dNumber
}

func getTransformerFn(s *[]int) transformFn {
	if (*s)[0] == 1 {
		return double
	}
	return triple
}

func double(n int) int {
	return n * 2
}

func triple(n int) int {
	return n * 3
}
