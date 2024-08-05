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
		tetrominoes := utils.ExtractTetromino(os.Args[1])
		if utils.IsValidTetromino(tetrominoes) {
			var trimmedTetrominoes [][]string
			for _, tetromino := range tetrominoes {
				trimH := utils.TrimHorizontal(tetromino)
				trimV := utils.TrimVertical(trimH)
				trimmedTetrominoes = append(trimmedTetrominoes, trimV)
			}
			fmt.Println(utils.RenameTetromino(trimmedTetrominoes))
		}
	} else {
		fmt.Fprintln(os.Stderr, `ERROR`)
	}
}
