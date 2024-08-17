package utils

import (
	"reflect"
	"testing"
)

func TestTrimHorizontal(t *testing.T) {
	type args struct {
		tetromino []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"horizontally trim unwanted . to optimize tetromino shape", args{[]string{"....", "....", "....", "####"}}, []string{"####"}},
		{"horizontally trim unwanted . to optimize tetromino shape", args{[]string{"....", ".###", ".#..", "...."}}, []string{".###", ".#.."}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimHorizontal(tt.args.tetromino); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimHorizontal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimVertical(t *testing.T) {
	type args struct {
		tetromino []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"vertically trim unwanted . to optimize tetromino shape", args{[]string{"...#", "...#", "...#", "...#"}}, []string{"#", "#", "#", "#"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimVertical(tt.args.tetromino); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimVertical() = %v, want %v", got, tt.want)
			}
		})
	}
}
