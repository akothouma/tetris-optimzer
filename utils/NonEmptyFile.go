package utils

import "os"

func NonEmptyFile(s string)bool{
	fileStat,_:=os.Stat(s)
    return fileStat.Size()!=0
}