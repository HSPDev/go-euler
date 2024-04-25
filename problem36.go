package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(numberString string) bool {
	runes := []rune(numberString)
	for forward, reverse := 0, len(runes)-1; forward < reverse; forward, reverse = forward+1, reverse-1 {
		if runes[forward] != runes[reverse] {
			return false
		}
	}
	return true
}

func toBase2(number int) string {
	result := ""
	for number > 0 {
		remainder := number % 2
		number = number / 2
		result = strconv.Itoa(remainder) + result
	}
	return result
}

// If somebody put a gun to my head to optimize this, I would find a way to avoid working with strings so much.
func main() {
	sum := 0
	for i := 1; i < 1000*1000; i++ {
		b2 := toBase2(i)
		if isPalindrome(b2) && isPalindrome(strconv.Itoa(i)) {
			fmt.Println("Found one!", i)
			sum += i
		}
	}
	fmt.Println(sum)
}
