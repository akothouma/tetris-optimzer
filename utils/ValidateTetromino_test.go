package utils

import (
	"testing"
)

func TestIsValidTetromino(t *testing.T) {
	t.Run("Check validity of tetromino", func(t *testing.T) {
	tetromino:= []string{
		"...#",
		"...#",
		"...#",
		"...#",	
}
tetrominoes:=[][] string{tetromino}
	

		if got := IsValidTetromino(tetrominoes); got != true {
			t.Errorf("IsValidTetromino() = %v, want true", got)
		
		}
	})

}