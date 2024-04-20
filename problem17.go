package main

import (
	"fmt"
	"math"
	"strings"
)

// Find the total length of all words used to write out numbers from 1 to 1000 in plain English.
func main() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		text := spellNumber(i)
		textStripped := strings.ReplaceAll(text, " ", "")
		textStripped = strings.ReplaceAll(textStripped, "-", "")
		length := len(textStripped)
		sum += length
		fmt.Printf("Number %d is %s with length %d.\n", i, textStripped, length)
	}
	fmt.Printf("Total length = %d\n", sum)
}

// Write number as english
func spellNumber(number int) string {

	//All the "custom" words with weird wordings
	words := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}
	if number <= len(words) {
		return words[number-1]
	}

	//prefix-words for 10's starting from 20 to 90. Hundreds are different.
	tennerWords := []string{
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}

	if number < 100 {
		tens := int(math.Floor(float64(number / 10)))
		tenWord := tennerWords[tens-2]
		remainder := number % 10
		if remainder == 0 {
			return tenWord
		}
		return fmt.Sprintf("%s-%s", tenWord, spellNumber(remainder))
	}

	if number < 1000 {
		hundreds := int(math.Floor(float64(number / 100)))
		remainder := number % 100
		if remainder == 0 {
			return fmt.Sprintf("%s %s", spellNumber(hundreds), "hundred")
		}
		return fmt.Sprintf("%s %s and %s", spellNumber(hundreds), "hundred", spellNumber(remainder))
	}

	//If it had to go higher, I would extract the above code for hundreds as thousands (i tried it)
	//but it also need refactoring to handle and vs. , when spelling it out (comma first, and only the last one)
	//It could probably also extract tens, and they could probably generalized fill in the last between 12 and 20,
	//but at that point we are almost doing Enterprise Fizzbuzz instead of solving the problem.
	if number == 1000 {
		return "one thousand"
	}

	panic(fmt.Sprintf("No clue what number %d is.", number))
}
