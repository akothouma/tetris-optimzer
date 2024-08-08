package main

import (
	"fmt"
	"os"
	"strings"

	"tetris/utils"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, `ERROR`)
		return
	}
	if !strings.HasSuffix(os.Args[1], ".txt") {
		fmt.Fprintln(os.Stderr, `ERROR`)
		return
	}
	if utils.NonEmptyFile(os.Args[1]) {
		tetrominoes, size := utils.ExtractTetromino(os.Args[1])
		var renamedTetrominoes [][]string
		var trimmedTetrominoes [][]string
		if utils.IsValidTetromino(tetrominoes) {

			for _, tetromino := range tetrominoes {
				trimH := utils.TrimHorizontal(tetromino)
				trimV := utils.TrimVertical(trimH)
				trimmedTetrominoes = append(trimmedTetrominoes, trimV)
			}
			renamedTetrominoes = utils.RenameTetromino(trimmedTetrominoes)
		}
		var finalBoard [][]string
		for {
			board := utils.DrawInitialBoard(size)
			finalBoard = utils.DrawFinalBoard(board, renamedTetrominoes)
			if finalBoard != nil {
				break
			}
			size++
		}
		utils.PrintBoard(finalBoard)
	} else {
		fmt.Fprintln(os.Stderr, `ERROR`)
	}
}
