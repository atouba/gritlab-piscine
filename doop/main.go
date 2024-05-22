package main

import (
	"os"
)

// Atoi

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
	for i := 0; i < strlen_Atoi(s); i++ {
		if rune(s[i]) == '-' && i != 0 {
      return 0
    } else if rune(s[i]) == '-' && i == 0 {
			sign = -1
		} else if rune(s[i]) >= '0' && rune(s[i]) <= '9' {
			if (res > 0 && sign == -1) || (res < 0 && sign == 1) {
				return 0
			}
			res += int(rune(s[i]) - '0') * sign
			if i < strlen_Atoi(s)-1 {
				res *= 10
			}
		}
	}

	return res 
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

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		os.Stdout.Write([]byte("9223372036854775808"))
		return
	}
	if n < 0 {
		os.Stdout.Write([]byte("-"))
		n *= -1
	}
	if n >= 10 {
		PrintNbr(n / 10)
		os.Stdout.Write([]byte(string((n % 10) + '0')))
	} else if n < 10 {
		os.Stdout.Write([]byte(string(n + '0')))
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

// func isOverFlow(res int, strs ...string) bool {
// 	for _, v := range strs {
// 		if v[0] == '-' {
// 			if len(v) > 20 {
// 				return false
// 			}
// 			if len(v) == 20 && Atoi(v) >= 0 {
// 				return false
// 			}
// 		} else {
// 			if len(v) > 19 {
// 				return false
// 			}
// 			if len(v) == 19 && Atoi(v) <= 0 {
// 				return false
// 			}
// 		}
// 		if (v[0] == '*' || v[0] == '+') && (res ) {
// 			return false
// 		}
// 	}
// 	return true
// }

func isActualNum(num string) bool {
	for _, v := range num {
		if v >= '1' && v <= '9' {
			return true
		}
	}
	return false
}

func validArgs(args []string) bool {
	if len(args) != 3 || !IsNumeric(args[0]) || !IsNumeric(args[2]) {
		return false
	}
	if !isValidOp(args[1]) {
		return false
	}
	num1 := Atoi(args[0])
	num2 := Atoi(args[2])
	if (num1 == 0 && isActualNum(args[0])) || (num2 == 0 && isActualNum(args[2])) {
		return false
	}
	if num2 == 0 && len(args[2]) == 1 {
		if args[1] == "/" {
			os.Stdout.Write([]byte("No division by 0\n"))
			return false
		} else if args[1] == "%" {
			os.Stdout.Write([]byte("No modulo by 0\n"))
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
	// 	if !isOverFlow(res, args[0], args[1]) {
	// 		return
	// 	}
	//   os.Stdout.WriteString([]byte(string(res)))
	PrintNbr(res)
  os.Stdout.Write("\n")
}

func main() {
	if !validArgs(os.Args[1:]) {
		return
	}
	printOutput(os.Args[1:])
}
