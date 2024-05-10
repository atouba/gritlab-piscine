package piscine

func index(s string) int {
	for i := range s {
		if s[i] != '0' {
			break
		}
		return i
	}
	return 0
}

func BasicAtoi(s string) int {
	var res int = 0

	//	for i := 0; i < StrLen(s); i++ {
	for i := index(s); i < StrLen(s); i++ {
		if rune(s[i]) >= '0' && rune(s[i]) <= '9' {
			res += int(rune(s[i]) - '0')
			if i < StrLen(s)-1 {
				res *= 10
			}
		}
	}

	return res
}
