package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(number int) bool {
	tester := strconv.Itoa(number)
	runes := []rune(tester)
	for forward, reverse := 0, len(runes)-1; forward < reverse; forward, reverse = forward+1, reverse-1 {
		if runes[forward] != runes[reverse] {
			return false
		}
	}
	return true
}
func main() {
	largestPalindrome := 0
	compound := [2]int{0, 0}
	//A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 times 99
	//Find the largest palindrome made from the product of two $3$-digit numbers.
	for a := 999; a > 100; a-- {
		for b := a; b > 100; b-- {
			test := a * b
			if isPalindrome(test) {
				if test > largestPalindrome {
					largestPalindrome = test
					compound[0] = a
					compound[1] = b
				}
			}
		}
	}
	fmt.Println("Found large palindrome!", largestPalindrome, "using ", compound[0], "x", compound[1])
}
