package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var c rune = '0'
	var n rune = '\n'
	for c <= '9' {
		z01.PrintRune(c)
		c++
	}
	z01.PrintRune(n)
}
