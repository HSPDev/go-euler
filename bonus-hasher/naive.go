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
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	//Hex encoded md5 hash we want to crack (My dog = ramos)
	target := "db3b992995b36a9d2ac616ea2867b14a"

	fmt.Println("Attempting to solve:", target)

	start := time.Now()
	fmt.Println(bruteforce(alphabet, target))
	elapsed := time.Since(start)
	fmt.Println("Time to run:", elapsed)
}

// Returns bruteforced MD5 hash
func bruteforce(alphabet string, target string) string {
	//We convert the md5 string (hex encoded) into pure bytes
	targetBytes, _ := hex.DecodeString(target)

	//Initialize the first byte to be the first rune.
	attempt := make([]byte, 1)
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
		hash := md5.Sum(attempt)
		//We perform fast equality checks using the built in bytes comparison
		if bytes.Equal(hash[:], targetBytes) {
			return string(attempt)
		}
	}
}
func index(alphabet string, b byte) int {
	for i := 0; i < len(alphabet); i++ {
		if b == byte(alphabet[i]) {
			return i
		}
	}
	panic("Rune " + string(b) + " not found in alphabet:" + alphabet)
}
