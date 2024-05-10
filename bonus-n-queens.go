package main

import (
	"fmt"
	"math"
)

//This is NOT an Euler problem, that I know of, but I found the N-queens problem on an article about backtracking
//and I wanted to try implementing that myself.

func main() {
	n := 4

	//Initialize board to zeros.
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}

	result, success := backtrack(board, 0)
	fmt.Println("Solved:", success)

	printBoard(result)
}

func backtrack(board [][]bool, newPlacement int) (newBoard [][]bool, validSolution bool) {
	n := len(board)

	//End of board. Is this a valid solution?
	if newPlacement >= n*n {
		sumQueens := 0
		for y := 0; y < n; y++ {
			for x := 0; x < n; x++ {
				if board[y][x] {
					sumQueens++
				}
			}
		}
		newBoard = board
		validSolution = sumQueens >= n
		return
	}

	//We index the placement into rows/columns. Index 5, in a grid 4x4 would mean row index 1 (second), column 0 (first)
	y := int(newPlacement / n)
	x := newPlacement % n

	if isLegalPlacement(board, x, y) {
		board[y][x] = true
		testBoard, testValid := backtrack(board, n*(y+1))
		if testValid {
			newBoard = testBoard
			validSolution = testValid
			return
		} else {
			board[y][x] = false
			return backtrack(board, newPlacement+1)
		}
	}

	return backtrack(board, newPlacement+1)
}

func isLegalPlacement(board [][]bool, x int, y int) bool {
	//Find all current placed queens, and analyze against those, instead of constantly looping the board X/Y/Diagonals which checks the same points over and over
	for ty := 0; ty < len(board); ty++ {
		for tx := 0; tx < len(board[ty]); tx++ {
			if board[ty][tx] {
				//This is a queen! Test against it!

				//Same row or column as our attempt, it can strike.
				if ty == y || tx == x {
					return false
				}

				//Check the diagional by testing if the coordinates are proportional.
				//E.g. 1,2 and 3,4 = ( 1-3 = -2) and (2-4 = -2)
				//The absolute comparison ensures it works both ways
				if math.Abs(float64(ty-y)) == math.Abs(float64(tx-x)) {
					return false
				}
			}
		}
	}
	return true
}

// Helper function as debugging aid
func printBoard(board [][]bool) {
	fmt.Println("Chess board:")
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if board[y][x] {
				fmt.Print("X ")
			} else {
				fmt.Print("- ")
			}
		}
		fmt.Println("")
	}
}
