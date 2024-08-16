package utils

import (
	"reflect"
	"testing"
)

func TestRenameTetromino(t *testing.T) {
	type args struct {
		tetrominoes [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"rename tetromino from # to corresponding alphabet",args{[][]string{{"#","#","#","#"},{"####"},{"###","..#"},{".##","##."},{"##","##"},{"##.",".##"},{"##",".#",".#"},{"###",".#."}}},
		[][]string{{"A","A","A","A"},{"BBBB"},{"CCC","..C"},{".DD","DD."},{"EE","EE"},{"FF.",".FF"},{"GG",".G",".G"},{"HHH",".H."}}},
	
}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RenameTetromino(tt.args.tetrominoes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RenameTetromino() = %v, want %v", got, tt.want)
			}
		})
	}
}

