package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	sum := sumUp(slice)
	fmt.Println(sum)
	sum2 := variadicSumUp(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(sum2)
	anotherSum2 := variadicSumUp(1, slice...)
	fmt.Println(anotherSum2)
}

func sumUp(s []int) (res int) {
	res = 0
	for _, value := range s {
		res += value
	}
	return
}

func variadicSumUp(startValue int, numbers ...int) (res int) {
	res = startValue * 10
	for _, value := range numbers {
		res += value
	}
	return
}
