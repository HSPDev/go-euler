package main

import (
	"fmt"
	"math"
)

func main() {
	var squareSum float64 = 0 //Sum of individually squared numbers
	var sum float64 = 0       //Sum of all numbers
	for i := 1; i <= 100; i++ {
		sum += float64(i)
		squareSum += math.Pow(float64(i), 2)
	}
	sumSquared := math.Pow(sum, 2)

	result := sumSquared - squareSum
	fmt.Println("Done! Difference is ", int(result))
}
