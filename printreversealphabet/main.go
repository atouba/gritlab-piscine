package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var c rune = 'z'
	var n rune = '\n'
	for c != 'a' {
		z01.PrintRune(c)
		c--
	}
	z01.PrintRune(c)
	z01.PrintRune(n)
}
