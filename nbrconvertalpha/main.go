package main

import (
	"os"

	"github.com/01-edu/z01"
)

func strLen(s string) int {
	var l int = 0

	for i := range s {
		l++
		i = i
	}
	return l
}

func index(s string) int {
	for i := range s {
		if s[i] != '0' {
			break
		}
		return i
	}
	return 0
}

func basicAtoi(s string) int {
	var res int = 0

	for i := index(s); i < strLen(s); i++ {
		if rune(s[i]) >= '0' && rune(s[i]) <= '9' {
			res += int(rune(s[i]) - '0')
			if i < strLen(s)-1 {
				res *= 10
			}
		} else {
			return 0
		}
	}

	return res
}

func PrintStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func letter(n int, isupp bool) rune {
	c := 'a' + n - 1
	if isupp {
		return rune(c - 32)
	}
	return rune(c)
}

func main() {
	var n int
	isupp := false
	i := 1
	var v string

	if len(os.Args) == 1 {
		return
	}
	if os.Args[1] == "--upper" {
		isupp = true
		i = 2
	}
	for i, v = range os.Args {
		if (i > 0 && isupp == false) || (i > 1 && isupp == true) {
			n = basicAtoi(v)
			if n > 0 && n <= 26 {
				z01.PrintRune(letter(n, isupp))
			} else {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
