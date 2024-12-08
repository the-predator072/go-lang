package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalnceFileName = "balance.txt"

func writeBalanceToFile(data float64) {
	balance := fmt.Sprint(data)
	os.WriteFile(accountBalnceFileName, []byte(balance), 0644)
}

func fetchBalanceFromFile() (balance float64) {
	data, _ := os.ReadFile(accountBalnceFileName)
	balanceTxt := string(data)
	balance, _ = strconv.ParseFloat(balanceTxt, 64)
	return
}

func main() {
	var accountBalance float64 = fetchBalanceFromFile()
	fmt.Println("Welcom TO Go Bank")
	for {
		fmt.Println("What do you want to do")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your Account Balance is: ", accountBalance)
		case 2:
			var amount float64
			fmt.Print("Enter the amount to be deposited: ")
			fmt.Scan(&amount)
			accountBalance += amount
			writeBalanceToFile(accountBalance)
			fmt.Printf("Amonut of %.2f deposited to your account, current balance is %.2f\n", amount, accountBalance)
		case 3:
			var amount float64
			fmt.Print("Enter the amount to be withdrawn: ")
			fmt.Scan(&amount)
			accountBalance -= amount
			writeBalanceToFile(accountBalance)
			fmt.Printf("Amonut of %.2f deposited to your account, current balance is %.2f\n", amount, accountBalance)
		default:
			fmt.Print("Exited Application")
		}
	}

}
