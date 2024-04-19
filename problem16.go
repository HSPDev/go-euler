package main

import (
	"fmt"
	"math"
)

func main() {
	//Sum of digits of 2¹⁰⁰⁰
	number := math.Pow(2, 1000)
	numberText := fmt.Sprintf("%.0f\n", number)
	runes := []rune(numberText)

	sum := 0
	for i := 0; i < len(runes)-1; i++ {
		value := int(runes[i] - '0')
		sum += value
	}

	fmt.Println("Sum is ", sum)
}
