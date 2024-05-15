package piscine

func SortIntegerTable(table []int) {
	for i := len(table) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if table[j] > table[j+1] {
				Swap(&table[j], &table[j+1])
			}
		}
	}
}
//

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
