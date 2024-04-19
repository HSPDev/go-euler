package main

import (
	"fmt"
)

// Gets the next collatz number after the one provided.
func nextCollatzNumber(number uint64) uint64 {
	if number%2 == 0 {
		return number / 2
	}
	return number*3 + 1
}

// Returns how long a sequence a given collatz start number ran for.
func sequenceLength(startNumber uint64) int {
	number := startNumber
	counter := 1 //First term is the start number
	for {
		counter++
		number = nextCollatzNumber(number)
		if number == 1 {
			break
		}
	}
	return counter
}
func main() {
	recordHolder := 0
	record := 0
	for i := 3; i < (1000 * 1000); i++ {
		length := sequenceLength(uint64(i))
		if length > record {
			recordHolder = i
			record = length
		}
	}

	fmt.Printf("Number %d has longest chain length %d\n", recordHolder, record)

}
