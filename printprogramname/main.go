package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}

func main() {
	PrintStr(string([]rune(os.Args[0])[2:]))
	z01.PrintRune('\n')
}
