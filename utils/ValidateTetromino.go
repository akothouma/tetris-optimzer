package utils

import (
	"fmt"
	"os"
)

/* Is ValidTetromino checks each tetromino in the tetromino slice and validates it based on number of #,connections and length of each string in the tetromino*/
func IsValidTetromino(tetrominoes [][]string) bool {
	for i := 0; i < len(tetrominoes); i++ {
		var connections int
		var count int
		for j := 0; j < len(tetrominoes[i]); j++ {
			for k := 0; k < len(tetrominoes[i][j]); k++ {
				if len(tetrominoes[i][j]) != 4 { // checks length of each  line of tetromino
					fmt.Fprintln(os.Stderr, `ERROR`)
					os.Exit(0)
				}
				if tetrominoes[i][j][k] == '#' {
					count++
					// check connection to neighbouring cells
					if k > 0 && tetrominoes[i][j][k-1] == '#' {
						connections++
					}
					if k < len(tetrominoes[i][j])-1 && tetrominoes[i][j][k+1] == '#' {
						connections++
					}
					if j > 0 && tetrominoes[i][j-1][k] == '#' {
						connections++
					}
					if j < len(tetrominoes[i])-1 && tetrominoes[i][j+1][k] == '#' {
						connections++
					}
				}
			}
		}

		if count != 4 || connections != 6 && connections != 8 {
			fmt.Fprintln(os.Stderr, `ERROR`)
			os.Exit(0)
		}

	}

	return true
}
