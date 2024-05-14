package piscine

func ToLower(s string) string {
	var res string
	var letter rune

	for _, v := range s {
		letter = v
		if v >= 'A' && v <= 'Z' {
			letter += 32
		}
		res += string(letter)
	}
	return res
}
