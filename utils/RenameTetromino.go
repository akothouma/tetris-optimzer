package utils

/* RenameTetromino takes the tetromino and gives it alphabetical ASCII value calculated from its index in the tetrominoes slice*/
func RenameTetromino(tetrominoes [][]string) [][]string {
	var renamedTetrominoes [][]string
	for i, array := range tetrominoes {
		var renamedTetromino []string
		var result string
		for _, str := range array {
			var char rune
			for _, ch := range str {
				if ch == '#' {
					char = rune(i + 65)
				} else {
					char = ch
				}
				result += (string(char))
			}
			renamedTetromino = append(renamedTetromino, result)
			result = ""
		}

		renamedTetrominoes = append(renamedTetrominoes, renamedTetromino)
	}
	return renamedTetrominoes
}
