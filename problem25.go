package main

import (
	"fmt"
	"math/big"
)

//Fibbonaci - find index for first number with 1000 digits
func main() {

	sequence := make([]*big.Int, 0)
	sequence = append(sequence, big.NewInt(1))
	sequence = append(sequence, big.NewInt(1))

	for i := 3; i<1000*1000; i++ {
		nextNumber := sequence[i-3].Add(sequence[i-3], sequence[i-2])
		sequence = append(sequence, nextNumber)

		asString := fmt.Sprintf("%d", nextNumber)
		if len(asString) == 1000 {
			fmt.Printf("Fibbonaci number %d = %d\n", i, nextNumber)
			fmt.Println("Found the first one at 1000 digits!")
			return
		}
	}

}
