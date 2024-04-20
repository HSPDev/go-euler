package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// They provided us 5000 names in a CSV file to work with.
func readNamesFromCsv() []string {
	f, err := os.Open("./0022_names.txt")
	if err != nil {
		panic("Unable to read input file")
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic("Unable to parse file as CSV for")
	}

	if len(records) != 1 {
		panic("Malformed data, we expect no columns, just pure CSV list of names.")
	}

	return records[0]
}

// Returns position 1-26 (A-Z) for the string letter (as byte).
func alphabeticPosition(letter byte) uint8 {
	//Uppercase A-Z is code 65-90 ASCII
	for i := 65; i <= 90; i++ {
		if letter == byte(i) {
			return uint8(i - 64)
		}
	}
	panic("Unknown rune received! " + string(letter))
}

// Returns sorted list of names, done using a variant of selection sort.
func sortAlphabetically(names []string) []string {
	//Loop all elements
	for i := 0; i < len(names); i++ {
		firstName := names[i]
		firstPosition := i
		//Loop all elements ahead of current one, find the smallest one.
		for j := i + 1; j < len(names); j++ {
			//Extract the current name
			test := names[j]

			//Loop all letters of both characters. If one is longer and they are equal, the shortest go first.
			minLength := min(len(firstName), len(test))
			for l := 0; l < minLength; l++ {
				alphabeticPosForTest := alphabeticPosition(test[l])
				alphabeticPosForCurrentBest := alphabeticPosition(firstName[l])
				//If the alphabetic position comes before our current best guess, this has to be moved up.
				if alphabeticPosForTest < alphabeticPosForCurrentBest {
					firstName = test
					firstPosition = j
					break
				}
				if alphabeticPosForTest > alphabeticPosForCurrentBest {
					//We dont need to compare the rest of the letters, as this one is already out.
					break
				}
				//If they are equal, the current contender is kept in the first spot, but the might get moved up in the next one.
				//Instance, current = AAC and test is AAB - Iteration 1 is equal, 2 is equal, but 3 is less than, so AAB goes before AAC.
				//e.g. = loop again!
			}
			//If they start out the same, but the currently testing name is shorter, move it up!
			if firstName[0:minLength] == test[0:minLength] && len(test) < len(firstName) {
				firstName = test
				firstPosition = j
			}
		}

		//If we have a change, swap!
		if firstPosition != i {
			names[i], names[firstPosition] = names[firstPosition], names[i]
		}

	}
	return names
}

// Returns the "score" based on the sum of alphabetic position of letters in name
func alphabeticScore(name string) int {
	var score int = 0
	for i := 0; i < len(name); i++ {
		score += int(alphabeticPosition(name[i]))
	}
	return score
}

func main() {
	records := readNamesFromCsv()
	records = sortAlphabetically(records)

	totalScore := 0
	for i := 0; i < len(records); i++ {
		//Score is based on position in alphabetic sorted list multiplied by the name score
		totalScore += ((i + 1) * alphabeticScore(records[i]))
	}

	fmt.Printf("Total score is %d\n", totalScore)
}
