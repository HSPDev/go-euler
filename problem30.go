package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	masterSum := 0
	for i := 10; i < 1000*1000; i++ {
		//Make the number into letters (so we can split digits)
		number := strconv.Itoa(i)

		sum := 0

		//Iterate the digits
		for _, c := range number {
			//Make each digit into a number again for MATH!
			n := float64(c - '0')
			sum += int(math.Pow(n, 5))
		}

		if sum == i {
			fmt.Println("Found one! =", sum)
			masterSum += sum
		}
	}
	fmt.Println("Master sum =", masterSum)
}
