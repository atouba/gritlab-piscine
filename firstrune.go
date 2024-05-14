package piscine

func FirstRune(s string) rune {
	for _, v := range s {
		return v
	}
	return 0
}
