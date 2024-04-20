package main

import (
	"fmt"
	"time"
)

func main() {

	start, _ := time.Parse("2006-01-02", "1901-01-01")
	end, _ := time.Parse("2006-01-02", "2000-12-31")
	sundayOnFirstCounter := 0
	for d := start; d.After(end) == false; d = d.AddDate(0, 0, 1) {
		if time.Weekday(0) == d.Weekday() && d.Day() == 1 {
			fmt.Println("Søndag 1st", d.Format("2006-01-02"))
			sundayOnFirstCounter++
		}
	}
	fmt.Println("Total Sunday on 1st count = ", sundayOnFirstCounter)

}
