package utils

import (
	"os"
)
/* NonEmpty ensures that sample.txt is not an empty file*/
func NonEmptyFile(s string) bool {
	fileStat, _ := os.Stat(s)
	return fileStat.Size() != 0
}
