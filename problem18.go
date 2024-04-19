package main

import (
	"fmt"
)

// Find path with largest value through trees top->bottom like this one - 3 + 7 + 4 + 9 =23
//	  3
//	 7 4
//	2 4 6
// 8 5 9 3
//It comes left aligned. Rules is as follows moving from top to bottom. Index in current row, can access same index, and index+1 in next row only.
//We need to make an effecient solution, not just bruteforce it, as they hint its repeated in problem 67 with 100+ lines and 2^99 options. FUCK ! :-)
//I considered Dijkstras which I did in school for my final project, but it cannot calculate most expensive path.
//I considered breadth first and depth first searches, but after reviewing the docs (long time no see), I found that they cannot handle leafs changing
//their values. They wouldn't find me the optimal most expensive path AFAIK?

//I'm going to invert the tree and roll the most expensive values upwards, and keep summing them.
//In my example I would make 8 5 9 3 roll up into 2 4 6. 8 is largest for 2. 9 is largest for 4. 9 is largest for 6. So 10, 13, 15.
//Afterwards 10, 13, 15, should be rolled up into 7 and 4. 13 is largest for 7, 15 is largest for 4. = 20, 19
//Roll up into 3. 20 is largest. 23 is the solution.

func main() {
	fmt.Println("Finding largest tree path")

	tree := [][]int{
		{75},
		{95, 64},
		{17, 47, 82},
		{18, 35, 87, 10},
		{20, 04, 82, 47, 65},
		{19, 01, 23, 75, 03, 34},
		{88, 02, 77, 73, 07, 63, 67},
		{99, 65, 04, 28, 06, 16, 70, 92},
		{41, 41, 26, 56, 83, 40, 80, 70, 33},
		{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
		{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
		{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
		{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
		{63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23},
	}

	//0 indexed, so len -1, but -2 because we start rolling the second last line (with the last one)
	//We assume well formed tree (1 more in every line), so no error checking.
	for y := len(tree) - 2; y >= 0; y-- {
		for x := 0; x < len(tree[y]); x++ {
			//Find max value in line underneath current one, in tree, in positions adjecent to current column
			tree[y][x] += max(tree[y+1][x], tree[y+1][x+1])
		}
	}

	//Top value should have the max value.
	fmt.Println("Found max value: ", tree[0][0])
}
