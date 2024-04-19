package main

import (
	"fmt"
)

func main() {
	//Find smallest number evenly dividable by 1-20. If it has to be dividable by 2, it has to even.
OUTER:
	for i := 2; i > 0; i += 2 {
		//Is always dividable by 1.
		//Is always diviable by 2 as we only increment by that.
		for j := 3; j <= 20; j++ {
			if i%j != 0 {
				continue OUTER
			}
		}
		fmt.Println("Winner! = ", i)
		break
	}
}
