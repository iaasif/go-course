package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println("Hello, World!")
	const inflationRate = 0.07
	var inverstmentAmount float64
	years := 10.0
	expectedReturnrate := 5.5

	fmt.Print("investment Amount: ")
	fmt.Scanln(&inverstmentAmount)

	futureValue := inverstmentAmount * math.Pow(1+expectedReturnrate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate, years)

	fmt.Println("Future Value:", futureValue)
	fmt.Println("Future Real Value:", futureRealValue)
}
