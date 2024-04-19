package main

import (
	"fmt"
)

func main() {
	//Done using https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes

	primeNumberToFind := 10001     //The prime number to find (first prime is 2)
	upperLimit := 10 * 1000 * 1000 //Find primes under this size (search all numbers under this size)

	sieve := make([]bool, upperLimit) //Initializes to false, so we strike non-primes out with true
	primeIndex := 0
	for i := 2; i < upperLimit-3; i++ {

		//If already striked out, skip this number
		if sieve[i-2] {
			continue
		}

		//Not striked out, is next number in list (and is a prime number)
		primeIndex += 1
		if primeIndex == primeNumberToFind {
			fmt.Println("Found wanted prime no.", primeNumberToFind, " with value ", i)
			return
		}

		//Strike out all numbers multiple of this.
		for y := (i * 2); y < upperLimit-3; y += i {
			sieve[y-2] = true
		}
	}
}
