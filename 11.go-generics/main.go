package main

import "fmt"

func main() {
	intRes := add(1, 2)
	floatRes := add(2.3, 2.4)
	strRes := add("Viaks ", "Negi")
	fmt.Println("Int Result: ", intRes)
	fmt.Println("Float Result: ", floatRes)
	fmt.Println("String Result: ", strRes)
}

func add[T string | int | float64](a, b T) T {
	return a + b
}
