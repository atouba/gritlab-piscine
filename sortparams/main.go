package main

import (
	"os"

	"github.com/01-edu/z01"
)

func nrune(s string, n int) rune {
	if n <= 0 || n > len(s) {
		return 0
	}
	for i, v := range s {
		if i == n-1 {
			return v
		}
	}
	return 0
}

func compare(a, b string) int {
	var j int = 0
	for _, v := range a {
		if v > nrune(b, j+1) {
			return 1
		} else if v < nrune(b, j+1) {
			return -1
		}
		j++
	}
	return 0
}

func PrintStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func sortIntegerTable(table []string) {
	for i := len(table) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if compare(table[j], table[j+1]) == 1 {
				temp := table[j]
				table[j] = table[j+1]
				table[j+1] = temp
			}
		}
	}
}

func main() {
	table := os.Args[1:]
	sortIntegerTable(table)
	for _, v := range table {
		PrintStr(v)
		z01.PrintRune('\n')
	}
}
