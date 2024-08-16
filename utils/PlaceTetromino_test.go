package utils

import "testing"

func TestPlaceTetrominoes(t *testing.T) {
	type args struct {
		board       [][]string
		tetrominoes [][]string
		i           int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"PlaceTetrominoes tests",args{[][] string{{"......"},{"......"},{"......"},{"......"},{"......"},{"......"}},
		[][]string{{"A","A","A","A"},{"BBBB"},{"CCC","..C"},{".DD","DD."},{"EE","EE"},{"FF.",".FF"},{"GG",".G",".G"},{"HHH",".H."}},8},true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlaceTetrominoes(tt.args.board, tt.args.tetrominoes, tt.args.i); got != tt.want {
				t.Errorf("PlaceTetrominoes() = %v, want %v", got, tt.want)
			}
		})
	}
}
