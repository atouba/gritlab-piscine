package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func main() {
	for i, v := range os.Args {
		if i == 0 {
			PrintStr(string([]rune(v)[2:]))
		} else {
			PrintStr(v)
		}
		z01.PrintRune('\n')
	}
}
