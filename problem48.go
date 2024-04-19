package main

import (
	"fmt"
	"math/big"
)

// Find last 10 digits of series (1¹ + 2² - 1000¹⁰⁰⁰)
func main() {
	sum := big.NewInt(1) //1¹ = 1
	for i := 2; i <= 1000; i++ {
		current := big.NewInt(int64(i))
		current.Exp(current, current, nil)
		sum.Add(sum, current)
	}

	sumText := fmt.Sprintf("%s", sum)

	fmt.Println("Last 10 is ", sumText[len(sumText)-10:])
}
