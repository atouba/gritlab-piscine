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
	var res int64 = 0

	for i := index(s); i < StrLen(s); i++ {
		res += int64(rune(s[i]) - '0')
		res *= 10
	}

	return int(res / 10)
}
