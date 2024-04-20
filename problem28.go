package main

import (
	"fmt"
)

// Spiraling problem seems crazy, but the period between diagonals is fixed.
// It starts out at 2, goes to 4, then to 6 e.g., increasing by 2 every 4 "walls" - We don't need to draw, just count the period.
func main() {
	const size int = 1001 //Length of the spiral X and Y

	diagonalSum := 1 //Our result (We must manually add the center)

	i := 1           //Counter for the spiral
	period := 2      //The current period (how long each segment is)
	wallCounter := 0 //Finished wall, increment period every 4 walls.
	for {
		i += period
		diagonalSum += i
		wallCounter++
		if wallCounter >= 4 {
			period += 2
			wallCounter = 0

			//When period is going to be bigger than our intended wall size, we are done
			if period > size {
				break
			}
		}
	}
	fmt.Println("diagonalSum = ", diagonalSum)
}
