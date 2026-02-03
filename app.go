package main

import (
	"fmt"
)

func main() {
	revenue := getUserInput("Revenue")
	expenses := getUserInput("Expenses")
	taxRate := getUserInput("Tax Rate")

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText + ": ")
	fmt.Scan(&userInput)
	return userInput
}

func calculateEBT(revenue, expenses float64) float64 {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
}
