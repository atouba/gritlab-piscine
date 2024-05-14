package piscine

func ToUpper(s string) string {
	var res string
	var letter rune

	for _, v := range s {
		letter = v
		if v >= 'a' && v <= 'z' {
			letter -= 32
		}
		res += string(letter)
	}
	return res
}
