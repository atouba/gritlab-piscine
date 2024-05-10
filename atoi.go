package piscine

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
	for i := index(s); i < StrLen(s); i++ {
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
	for i := index(s); i < StrLen(s); i++ {
		if rune(s[i]) == '-' {
			sign = -1
			continue
		}
		if rune(s[i]) >= '0' && rune(s[i]) <= '9' {
			res += int(rune(s[i]) - '0')
			if i < StrLen(s)-1 {
				res *= 10
			}
		}
	}

	return res * sign
}
