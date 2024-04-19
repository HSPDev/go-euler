package main

import (
	"fmt"
	"math"
)

//Triangle numbers. Triangle numbers are made by adding all previous numbers together.
//1,3,6,10,15,21,28
//+2,+3,+4 e.g.

//28 has 5 diviors
//Find the first one with over 500 diviors

func countDivisors(number int) int {
	count := 1 //All can be divided by 1

	//The naive approach is to loop up until the number, but if we instead utilize that
	//the divisors come in pairs at top and bottom ends, we can drastically reduce runtime by only working up until the square root
	//We just have to assume that every single divisor is actually 2, unless it's the actual square root as thats it's perfectly divisable.
	//E.g. 100 = 1,100 2,50 4,25 10,10
	//E.g. 16 = 1,16 2,8 4,4 ((That's the same, one divisor, don't cpunt double))
	for i := 2; i <= int(math.Ceil(math.Sqrt(float64(number)))); i++ {
		if number%i == 0 {
			count++
			if number/i != i {
				count++
			}
		}
	}
	return count
}
func main() {
	number := 0
	for i := 1; i < 1000*1000; i++ {
		number += i
		divisors := countDivisors(number)

		if i%100 == 0 {
			fmt.Printf("Triangle number %d is %d and has %d divisors.\n", i, number, divisors)
		}

		if divisors > 500 {
			fmt.Printf("Found first number over 500 divisors ( = %d)!\n", number)
			return
		}
	}
}
