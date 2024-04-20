package main

import (
	"fmt"
	"math"
)

func sumOfDivisors(number int) int {
	sum := 0

	//You can never divide something properly with less than half its value, as that would be <2
	limit := int(math.Ceil(float64(number)))
	for i := 1; i < limit; i++ {
		//test if i is a proper divisor for our number
		if number%i == 0 {
			sum += i
		}
	}

	return sum
}

func main() {
	fmt.Println("Amicable numbers")
	sum := 0
	for i := 1; i < 10*1000; i++ {
		//Find amicable numbers (where the sum of divisors's sum of divisors is the original number)
		test := sumOfDivisors(i)

		//Sum of divisors for some numbers (like 6, 1+2+3 = 6, actually means the number is the same, and it is not a pair - gotcha!)
		if i == sumOfDivisors(test) && i != test {
			fmt.Printf("%d\n", i)
			sum += i
		}
	}
	fmt.Printf("Sum of amicalable numers = %d\n", sum)
}
