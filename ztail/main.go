package main

import (
	"fmt"
	"os"
)

func index_atoi(s string) int {
	for i := range s {
		if s[i] != '0' {
			break
		}
		return i
	}
	return 0
}

func strlen_atoi(s string) int {
	var l int = 0

	for i := range s {
		l++
		i = i
	}
	return l
}

func count_sign(s string) bool {
	for i := 0; i < strlen_atoi(s); i++ {
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
	for i := index_atoi(s); i < strlen_atoi(s); i++ {
		if rune(s[i]) == '-' {
			sign = -1
			continue
		}
		if rune(s[i]) >= '0' && rune(s[i]) <= '9' {
			res += int(rune(s[i]) - '0')
			if i < strlen_atoi(s)-1 {
				res *= 10
			}
		}
	}

	return res * sign
}

func main() {
	var ret_value int

	l := len(os.Args)
	if l != 1 {
		c := Atoi(os.Args[2])
		for i := 3; i < l; i++ {
			data, err := os.ReadFile(os.Args[i])
			if err == nil {
				if i != 3 {
					fmt.Printf("\n")
				}
				fmt.Printf("==> %s <==\n", os.Args[i])
				fmt.Printf(string(data[len([]byte(data))-c:]))
			} else {
				fmt.Printf(err.Error())
				fmt.Printf("\n")
				ret_value = 1
			}
		}
	}
	os.Exit(ret_value)
}
