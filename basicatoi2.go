package piscine

func BasicAtoi2(s string) int {
	var res int = 0

	//	for i := 0; i < StrLen(s); i++ {
	for i := index(s); i < StrLen(s); i++ {
		if rune(s[i]) >= '0' && rune(s[i]) <= '9' {
			res += int(rune(s[i]) - '0')
			if i < StrLen(s)-1 {
				res *= 10
			}
		} else {
      return 0
    }
	}

	return res
}
