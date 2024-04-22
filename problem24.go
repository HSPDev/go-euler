package main

import (
	"fmt"
	"strconv"
)

// All permutations of 0 1 2 3 4 5 6 7 8 9 in lexographic order.
// find the millionth one.
// First problem I had to google to learn about permutations and Lehmans coding. It makes sense and this is my own implementation.
func main() {

	//Find millionth permutation
	target := 1000 * 1000

	//Numbers to pluck from when doing the permutation, we'll remove each integer as we use them
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Finding", target, "permutation of source array", numbers)

	//First permutation is the current one, so subtract one.
	target -= 1

	//The resulting string
	string := ""

	for i := 9; i >= 0; i-- {
		//Find factorioal of i, to find the value of the left most position.
		f := factorial(i)
		c := int(target / f) //Round it down, to find the absolute value, dont care about remainder
		target = target % f  //Use the remainder for the next iteration, the integer to the right

		//Grab this number, next in our result
		string += strconv.Itoa(numbers[c])

		fmt.Println("Next index:", c, "with current numbers", numbers)

		//Grab all items up intil this index, and all the items +1 index after this (to preverse order)
		// (spread the final array to append it to the first ,as we can't just concat two arrays)
		numbers = append(numbers[:c], numbers[c+1:]...)
	}

	fmt.Println("Resulting permutation:", string)
}

func factorial(number int) int {
	if number <= 1 {
		return 1
	}
	return number * factorial(number-1)
}
