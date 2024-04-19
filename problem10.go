package main

import (
	"fmt"
)

func main() {
	//Done using https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes

	upperLimit := 2 * 1000 * 1000 //Find sum of all primes below 2 million

	sieve := make([]bool, upperLimit) //Initializes to false, so we strike non-primes out with true
	primeSum := 0

	for i := 2; i < upperLimit-3; i++ {

		//If already striked out, skip this number
		if sieve[i-2] {
			continue
		}

		//Strike out all numbers multiple of this.
		for y := (i * 2); y < upperLimit-3; y += i {
			sieve[y-2] = true
		}

		//It's a prime though! So we add it to the sum.
		primeSum += i
	}
	fmt.Println("Sum of primes below", upperLimit, " is = ", primeSum)
}
