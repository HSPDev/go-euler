package main

import (
	"fmt"
	"runtime"
	"time"
	"unicode/utf8"
)
//This is a playground to find out how we can produce more strings, FASTER!
func main() {
	//Force using all available cores!
	runtime.GOMAXPROCS(runtime.NumCPU())

	//Alphabet to use (case sensitive)
	alphabet := []rune("abcdefghijklmnopqrstuvwxyzæøå")

	start := time.Now()
	
	//Initialize the first byte to be the first rune.
	attempt := make([]rune, 1)
	attempt[0] = alphabet[0]

	for i := 0; i<1000*1000*20;i++ {

		//Rollover next character in array
		rolloverPosition := 0
		for {
			//If we are beyond end of string, we increase size of string.
			//This happens a bit in the beginning, but less and less as the string size increases
			//MARGINAL gains, if any.
			if rolloverPosition >= len(attempt) {
				attempt = append(attempt, alphabet[0])
				break
			}
			//Check if the next character needs to wrap around and carry forward
			//We do this by checking if the item is the last one in the alphabet.
			if attempt[rolloverPosition] == alphabet[len(alphabet)-1] {
				//It did, so reset it and rollover the next character in the next loop.
				attempt[rolloverPosition] = alphabet[0]
				rolloverPosition++
				continue
			}
			//Just roll the character forward and exit, as it's within alphabet
			attempt[rolloverPosition] = alphabet[index(alphabet, attempt[rolloverPosition])+1]
			break
		}
		size := 0
    	for _, r := range attempt {
	        size += utf8.RuneLen(r)
	    }

	    bs := make([]byte, size)

	    count := 0
	    for _, r := range attempt {
	        count += utf8.EncodeRune(bs[count:], r)
	    }
		_ = bs

	}

	elapsed := time.Since(start)
	fmt.Println("Time to run:", elapsed)
}


func index(alphabet []rune, r rune) int {
	for i := 0; i < len(alphabet); i++ {
		if r == alphabet[i] {
			return i
		}
	}
	panic("Rune " + string(r) + " not found in alphabet")
}