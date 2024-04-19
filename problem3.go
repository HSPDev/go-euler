package main

import (
	"fmt"
)

func main() {
	//The prime factors of 13195 are 5, 7, 13 and 29
	//Find largest prime factor of 600851475143
	//I discovered that given primes can only be divided by themselves and 1,
	//And you do prime factorization by dividing by increasingly larger prime numbers (where remainder is 0),
	//it's possible to ignore whatever anything is a prime, and just start from 2 and loop upwards.

	i := 600851475143
	test := 2
	for {
		if i%test == 0 {
			i /= test
			if i == 1 {
				fmt.Println("DONE! Result =", test)
				break
			}
			test = 2 //Reset back to first prime.
			continue
		}

		test++ //Increment, would be better if only primes but fuck that.
	}
}
