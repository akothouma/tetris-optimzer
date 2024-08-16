package utils

import (
	"reflect"
	"testing"
)

func TestExtractTetromino(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  [][]string
		want1 int
	}{
		{"Extracting tetrominoes from sample1.txt",args{"sample1.txt"},[][]string{{"...#","...#","...#","...#"},{"....","....","....","####"},{".###","...#","....","...."},{"....","..##",".##.","...."},{"....",".##.",".##.","...."},{"....","....","##..",".##."},{"##..",".#..",".#..","...."},{"....","###.",".#..","...."}},8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ExtractTetromino(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractTetromino() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExtractTetromino() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
