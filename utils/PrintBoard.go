package utils

import (
	"fmt"
	"math"
)
/* DrawInitialBoard makes the starting board size with reference to number of tetrominoes*/
func DrawInitialBoard(size int) [][]string {
	initialSize := int(math.Ceil((math.Sqrt(float64(size) * 4))))
	board := make([][]string, initialSize)
	for row := range board {
		board[row] = make([]string, initialSize)
		for col := range board[row] {
			board[row][col] = "."
		}
	}
	return board
}

func PrintBoard(board [][]string) {
	for _, row := range board {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
}
