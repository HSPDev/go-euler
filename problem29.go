package main

import (
	"fmt"
	"math/big"
)

// Biggest hurdle with this task is that every single type overflowed and simple slices.contains() broke with big integers.
func main() {
	const lower int64 = 2
	const upper int64 = 100

	numbers := make([]big.Int, 0)
	for a := lower; a <= upper; a++ {
	bLoop:
		for b := lower; b <= upper; b++ {
			var n big.Int
			n.Exp(big.NewInt(a), big.NewInt(b), nil)

			for i := 0; i < len(numbers); i++ {
				if numbers[i].Cmp(&n) == 0 {
					continue bLoop
				}
			}
			numbers = append(numbers, n)
		}
	}

	fmt.Println(len(numbers))
}
