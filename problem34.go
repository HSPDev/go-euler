package main

import (
	"fmt"
	"strconv"
)

// Pure bruteforce approach, reused code from problem 30.
func main() {
	masterSum := 0
	for i := 10; i < 1000*1000*10; i++ {
		//Make the number into letters (so we can split digits)
		number := strconv.Itoa(i)

		sum := 0

		//Iterate the digits
		for _, c := range number {
			//Make each digit into a number again for MATH!
			n := int(c - '0')
			sum += factorial(n)
		}

		if sum == i {
			fmt.Println("Found one! =", sum)
			masterSum += sum
		}
	}
	fmt.Println("Master sum =", masterSum)
}

func factorial(number int) int {
	if number <= 1 {
		return 1
	}
	return number * factorial(number-1)
}
