package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ExtractTetromino(s string) [][]string {
	var tetromino []string
	var tetrominoes [][]string
	if readFile, err := os.Open(s); err != nil {
		fmt.Println(err)
	} else {
		fileScanner := bufio.NewScanner(readFile)

		for fileScanner.Scan() {
			fileLines := fileScanner.Text()

			if fileLines == "" {
				tetrominoes = append(tetrominoes, tetromino)
				tetromino = []string{}
			} else {
				tetromino = append(tetromino, fileLines)
			}
		}
		if len(tetromino)!=0{
			tetrominoes = append(tetrominoes, tetromino)
		}
	}
	return tetrominoes
}
