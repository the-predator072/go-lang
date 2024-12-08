package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	tax := revenue * taxRate / 100
	ebt := revenue - expenses
	eat := ebt - tax

	fmt.Println(ebt)
	fmt.Println(eat)
}
