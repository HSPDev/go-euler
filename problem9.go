package main

import (
	"fmt"
	"math"
)

func main() {
	//Pythogran Tripplet.
	//a < b < c, a+b+c = 1000, a² + b² = c²
	//I got no fucking clue, I suspect it can be resolved as a differential equation, but it's small scale
	//so I can do it nummerically.
	//I assume all variables satifies >=1

	//Start at 997, because minimum values for a is 1, b is 2 and c = 1000-1-2
	for c := (1000 - 1 - 2); c > 1; c-- {
		for b := c - 1; b > 1; b-- {
			for a := b - 1; a > 1; a-- {
				if a+b+c != 1000 {
					continue
				}
				fmt.Println("Testing ", a, b, c)

				if math.Pow(float64(a), 2)+math.Pow(float64(b), 2) == math.Pow(float64(c), 2) {
					fmt.Println("Found the solution!", a*b*c)
					return
				}
			}
		}
	}
}
