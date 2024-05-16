package piscine

func SplitWhiteSpaces(s string) []string {
	var arr []string

	for i := 0; i < StrLen(s); i++ {
		for s[i] == ' ' {
			i++
		}

		if length_word := Index(s[i:], " "); length_word != -1 {
			arr = append(arr, s[i:i+length_word])
		} else {
			arr = append(arr, s[i:])
			break
		}
		i += Index(s[i:], " ")
	}
	return arr
}
