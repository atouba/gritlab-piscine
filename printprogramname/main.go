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
	PrintStr(string([]rune(os.Args[0])[2:]))
	z01.PrintRune('\n')
}
