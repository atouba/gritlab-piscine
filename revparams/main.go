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

func revParam(args []string, i int) {
	if i < len(args) {
		revParam(args, i+1)
		PrintStr(args[i])
		z01.PrintRune('\n')
	}
}

func main() {
	revParam(os.Args[1:], 0)
}
