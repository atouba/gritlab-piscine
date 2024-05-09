package piscine

import "github.com/01-edu/z01"

func printstr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		printstr("-9223372036854775808")
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n *= -1
	}
	if n >= 10 {
		PrintNbr(n / 10)
		z01.PrintRune(rune('0' + (n % 10)))
	} else if n < 10 {
		z01.PrintRune(rune('0' + n))
	}
}
