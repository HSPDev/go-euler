package main

import "fmt"

func main() {
	const maxLimit = 4 * 1000 * 1000

	//Holds sum of even valued numbers
	sum := 2 //We start out with 2, as we skip those

	//Holds fibbonaci counters
	previous := [2]int{1, 2}
	for {
		current := previous[0] + previous[1]
		if current > maxLimit {
			break
		}

		previous[0] = previous[1]
		previous[1] = current

		if current%2 == 0 {
			sum += current
		}
	}

	fmt.Println("Done! Sum is ", sum)
}
