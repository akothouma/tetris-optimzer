package utils

import "strings"

/* TrimHorizontal removes the string contents of the tertomino that have no characters in them*/
func TrimHorizontal(tetromino []string) []string {
	var horizontalTrimmed []string
	for j := 0; j < len(tetromino); j++ {
		ValidRow := false
		for k := len(tetromino[j]) - 1; k >= 0; k-- {
			if strings.Contains(string(tetromino[j][k]), "#") {
				ValidRow = true
				break
			}
		}
		if ValidRow {
			horizontalTrimmed = append(horizontalTrimmed, tetromino[j])
		}
	}
	return horizontalTrimmed
}

/* TrimVertical checks the column in the tetromino for # and marks it as valid*/
func TrimVertical(tetromino []string) []string {
	var i int

	for k := len(tetromino[0]) - 1; k >= 0; k-- {
		columnIsValid := false
		for j := 0; j < len(tetromino); j++ {
			if strings.Contains(string(tetromino[j][k]), "#") {
				columnIsValid = true
				break
			}
		}
		if !columnIsValid {
			for i = range tetromino {
				tetromino[i] = RemoveCharAtIndex(tetromino[i], k)
			}
		}

	}
	return tetromino
}

func RemoveCharAtIndex(input string, index int) string {
	return input[:index] + input[index+1:]
}
