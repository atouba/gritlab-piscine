package piscine

import "github.com/01-edu/z01"

func minimum(n int) int {
	var res int = 0

	for i := 0; i < n-1; i++ {
		res *= 10
		res += i + 1
	}
	return res
}

func maximum(n int) int {
	var ten int = 1
	var maxi int = 0
	var max_digit int = 9

	for i := n; i > 0; i-- {
		maxi += (max_digit * ten)
		ten *= 10
		max_digit--
	}
	return maxi
}

func printNbr(nbr int) {
	if nbr >= 10 {
		PrintNbr(nbr / 10)
		z01.PrintRune(rune('0' + (nbr % 10)))
	} else if nbr < 10 {
		z01.PrintRune(rune('0' + nbr))
	}
}

func lenght_num(nbr int) int {
	var length int = 0

	for nbr != 0 {
		nbr /= 10
		length++
	}
	return length
}

func printnum(nbr int, n int) {
	var length int = 0

	if nbr == 0 {
		z01.PrintRune('0')
		return
	}
	length = lenght_num(nbr)
	if length < n {
		z01.PrintRune('0')
	}
	printNbr(nbr)
}

func tens(t int) int {
	var teny int = 1

	for teny < t {
		teny *= 10
	}
	return teny / 10
}

func checkComb(n int) bool {
	var x int = 0

	for n > 9 {
		x = n % 10
		n /= 10
		if n%10 >= x {
			return false
		}
	}
	return true
}

func PrintCombN(n int) {
	var maxim int = maximum(n)

	for start := minimum(n); start <= maximum(n); start++ {
		if checkComb(start) {
			printnum(start, n)
			if start != maxim {
				printstr(", ")
			}
		}
	}
	z01.PrintRune('\n')
}
