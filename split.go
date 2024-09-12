package piscine

func Split(s, sep string) []string {
	var arr []string

	for i := 0; i < StrLen(s); {
		for i + len(sep) <= len(s) && Compare(s[i:i+len(sep)], sep) == 0 {
			i += len(sep)
		}

		if length_word := Index(s[i:], sep); length_word != -1 {
			arr = append(arr, s[i:i+length_word])
		} else {
			arr = append(arr, s[i:])
			break
		}
		i += Index(s[i:], sep)
	}
	return arr
}
