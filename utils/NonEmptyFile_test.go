package utils

import "testing"

func TestNonEmptyFile(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"sample.txt is empty", args{"sample1.txt"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NonEmptyFile(tt.args.s); got != tt.want {
				t.Errorf("NonEmptyFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
