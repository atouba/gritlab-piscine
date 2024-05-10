package piscine

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
	var count int = 0

	for i := range s {
		if s[i] == '+' || s[i] == '-' {
			count++
		}
	}
	if count > 1 {
		return false
	}
	for i := index_atoi(s); i < strlen_atoi(s); i++ {
		if rune(s[i]) == '+' || rune(s[i]) == '-' {
			if rune(s[i+1]) > '9' || rune(s[i+1]) < '0' {
				return false
			}
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
