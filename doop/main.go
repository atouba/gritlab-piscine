package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Atoi

func index_Atoi(s string) int {
	for i := range s {
		if s[i] != '0' {
			break
		}
		return i
	}
	return 0
}

func strlen_Atoi(s string) int {
	var l int = 0

	for i := range s {
		l++
		i = i
	}
	return l
}

func count_sign(s string) bool {
	for i := 0; i < strlen_Atoi(s); i++ {
		if rune(s[i]) == '+' || rune(s[i]) == '-' {
			if i != 0 {
				return false
			}
		} else if rune(s[i]) < '0' || rune(s[i]) > '9' {
			return false
		}
	}
	return true
}

func Atoi(s string) int {
	var res int = 0
	var sign int = 1

	if !count_sign(s) {
		return 0
	}
	for i := index_Atoi(s); i < strlen_Atoi(s); i++ {
		if rune(s[i]) == '-' {
			sign = -1
			continue
		}
		if rune(s[i]) >= '0' && rune(s[i]) <= '9' {
			res += int(rune(s[i]) - '0')
			if i < strlen_Atoi(s)-1 {
				res *= 10
			}
		}
	}

	return res * sign
}

// isNumeric

func IsNumeric(s string) bool {
	for i, v := range s {
		if i == 0 && v == '-' {
			continue
		}
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

// printNbr

func printStr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}

func printt(strs ...string) {
	for _, v := range strs {
		printStr(v)
	}
}

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		printStr("-9223372036854775808")
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

// doop

func isValidOp(op string) bool {
	if op != "+" && op != "-" && op != "/" && op != "*" && op != "%" {
		return false
	}
	return true
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func isOverFlow(res int, strs ...string) bool {
	for _, v := range strs {
		if v[0] == '-' {
			if len(v) > 20 {
				return false
			}
			if len(v) == 20 && Atoi(v) >= 0 {
				return false
			}
		} else {
			if len(v) > 19 {
				return false
			}
			if len(v) == 19 && Atoi(v) <= 0 {
				return false
			}
		}
		if (v[0] == '*' || v[0] == '+') && (abs(res) < abs(Atoi(v))) {
			return false
		}
	}
	return true
}

func validArgs(args []string) bool {
	if len(args) != 3 || !IsNumeric(args[0]) || !IsNumeric(args[2]) {
		return false
	}
	if !isValidOp(args[1]) {
		return false
	}
	if Atoi(args[2]) == 0 {
		if args[1] == "/" {
			printStr("No division by 0\n")
			return false
		} else if args[1] == "%" {
			printStr("No modulo by 0\n")
			return false
		}
	}
	return true
}

func printOutput(args []string) {
	var res int

	c := args[1][0]
	switch c {
	case '+':
		res = Atoi(args[0]) + Atoi(args[2])
	case '-':
		res = Atoi(args[0]) - Atoi(args[2])
	case '*':
		res = Atoi(args[0]) * Atoi(args[2])
	case '/':
		res = Atoi(args[0]) / Atoi(args[2])
	case '%':
		res = Atoi(args[0]) % Atoi(args[2])
	}
	if !isOverFlow(res, args[0], args[1]) {
		return
	}
	PrintNbr(res)
	z01.PrintRune('\n')
}

func main() {
	if !validArgs(os.Args[1:]) {
		return
	}
	printOutput(os.Args[1:])
}
