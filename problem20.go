package main

import (
	"fmt"
	"math/big"
)

func main() {
	//NOTE I GOT BITTEN! Float64 silently overflowed. Need big int package. 100! is insanity

	//Sum of digits of 100!
	number := big.NewInt(100)
	for i := 99; i > 0; i-- {
		number = number.Mul(number, big.NewInt(int64(i)))
	}

	numberText := fmt.Sprintf("%s", number)

	runes := []rune(numberText)

	sum := 0
	for i := 0; i < len(runes)-1; i++ {
		value := int(runes[i] - '0')
		sum += value
	}

	fmt.Println("Sum is ", sum)
}
