package main

import "fmt"

func main() {
	fmt.Println(factorialWithLoop(6))
	fmt.Println(factorialWithRecursion(6))
}

func factorialWithRecursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialWithRecursion(n-1)
}

func factorialWithLoop(n int) (res int) {
	res = 1
	for i := 0; i < n; i++ {
		res *= n - i
	}
	return
}
