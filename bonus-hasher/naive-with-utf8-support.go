package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {

	//Alphabet to use (case sensitive)
	alphabet := []rune("abcdefghijklmnopqrstuvwxyzæøå")

	//Hex encoded md5 hash we want to crack ("æblet" = Danish for "The apple" to test utf8)
	const target string = "179eddcf882aef741257145e46e8820f"

	fmt.Println("Attempting to solve:", target)

	start := time.Now()
	fmt.Println(bruteforce(alphabet, target))
	elapsed := time.Since(start)
	fmt.Println("Time to run:", elapsed)
}

// Returns bruteforced MD5 hash
func bruteforce(alphabet []rune, target string) string {
	//We convert the md5 string (hex encoded) into pure bytes
	targetBytes, _ := hex.DecodeString(target)

	//Initialize the first byte to be the first rune.
	attempt := make([]rune, 1)
	attempt[0] = alphabet[0]

	for {

		//Rollover next character in array
		rolloverPosition := 0
		for {
			//If we are beyond end of string, we increase size of string.
			if rolloverPosition >= len(attempt) {
				attempt = append(attempt, alphabet[0])
				break
			}
			//Check if the next character needs to wrap around and carry forward
			n := index(alphabet, attempt[rolloverPosition]) + 1
			if n >= len(alphabet) {
				//It did, so reset it and rollover the next character in the next loop.
				attempt[rolloverPosition] = alphabet[0]
				rolloverPosition++
				continue
			}
			//Just roll the character forward and exit, as it's within alphabet
			attempt[rolloverPosition] = alphabet[n]
			break
		}

		//Hash our byte array.
		hash := md5.Sum([]byte(string(attempt)))
		//We perform fast equality checks using the built in bytes comparison
		if bytes.Equal(hash[:], targetBytes) {
			return string(attempt)
		}
	}
}
func index(alphabet []rune, r rune) int {
	for i := 0; i < len(alphabet); i++ {
		if r == alphabet[i] {
			return i
		}
	}
	panic("Rune " + string(r) + " not found in alphabet")
}
