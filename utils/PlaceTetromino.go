package utils

/*
DrawBoard initiates the placing of tetrominoes,It calls the reccursive function
PlaceTetrominoes. It returns board should all the tetrominoes get placed or
nil should no solution with current board size is reached
*/
func DrawFinalBoard(board [][]string, tetrominoes [][]string) [][]string {
	if PlaceTetrominoes(board, tetrominoes, 0) {
		return board
	}
	return nil
}

/* PlaceTetrominoes goes through each cell on the board and checks if the tetromino is placable at that cell.
If it's possible,it's placed  and the function reccurs to place the next tetromino.If its possible to place the
next tetromino,The tetromino is placed otherwise it is removed.*/

func PlaceTetrominoes(board [][]string, tetrominoes [][]string, i int) bool {
	if len(tetrominoes) == i {
		return true
	}
	tetromino := tetrominoes[i]
	for row := range board {
		for col := range board[row] {
			if IsPlacable(board, tetromino, row, col) {
				PlaceTetromino(board, tetromino, row, col)
				if PlaceTetrominoes(board, tetrominoes, i+1) {
					return true
				}
				RemoveTetromino(board, tetromino, row, col)
			}
		}
	}
	return false
}

/* PlaceTetromino ranges through a tetromino finding the non-empty tetromino cells and places them in the board cells relative to the tetromino cell value */
func PlaceTetromino(board [][]string, tetromino []string, x, y int) {
	for relRow, row := range tetromino {
		for relCol, ch := range row {
			if tetromino[relRow][relCol] != '.' {
				board[relRow+x][relCol+y] = string(ch)
			}
		}
	}
}

/* RemoveTetromino ranges through a tetromino finding the non-empty tetromino cells and removes  them from the board cells relative to the tetromino cell value */
func RemoveTetromino(board [][]string, tetromino []string, x, y int) {
	for relRow, row := range tetromino {
		for relCol, ch := range row {
			if ch != '.' {
				board[x+relRow][y+relCol] = "."
			}
		}
	}
}

/*  Is Placable ranges through the tetromino and finds the non-empty tetromino cells and checks if their  relative cell position on the board is within board bounds or is occupied*/
func IsPlacable(board [][]string, tetromino []string, x, y int) bool {
	var boardX, boardY int
	for relRow, row := range tetromino {
		for relCol, ch := range row {
			if ch != '.' {
				boardX, boardY = x+relRow, y+relCol
				if boardX >= len(board) || boardY >= len(board[0]) || board[boardX][boardY] != "." {
					return false
				}
			}
		}
	}
	return true
}
