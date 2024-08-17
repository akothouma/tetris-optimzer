package utils

import (
	"bufio"
	"fmt"
	"os"
)

/* ExtractTetromino reads the sample.txt returning tetrominoes and the number of tetrominoes */
func ExtractTetromino(s string) ([][]string, int) {
	var tetromino []string
	var tetrominoes [][]string
	var lineCount int
	if readFile, err := os.Open(s); err != nil {
		fmt.Println(err)
	} else {
		fileScanner := bufio.NewScanner(readFile)

		for fileScanner.Scan() {
			fileLines := fileScanner.Text()
			if lineCount == 4 && fileLines == "" {
				tetrominoes = append(tetrominoes, tetromino)
				tetromino = []string{}
				lineCount = 0

			} else {
				if fileLines != "" {
					tetromino = append(tetromino, fileLines)
					lineCount++
				}

			}
		}
		if len(tetromino) != 0 {
			tetrominoes = append(tetrominoes, tetromino)
		}

	}
	size := len(tetrominoes)
	return tetrominoes, size
}
