package piscine

func isalpha(c rune) bool {
	if (c < '0' || c > '9') && (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
		return false
	}
	return true
}

func tolower(s string) string {
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

func islower(s string) bool {
	for _, v := range s {
		if v < 'a' || v > 'z' {
			return false
		}
	}
	return true
}

func Capitalize(s string) string {
	var lowercase_s []rune = []rune(tolower(s))
	for i := range lowercase_s {
		if (islower(string(lowercase_s[i])) && i == 0) || (islower(string(lowercase_s[i])) && !isalpha(lowercase_s[i-1])) {
			lowercase_s[i] -= 32
		}
	}
	return string(lowercase_s)
}
