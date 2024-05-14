package piscine

func NRune(s string, n int) rune {
	if n <= 0 || n > len(s) {
		return 0
	}
	for i, v := range s {
		if i == n-1 {
			return v
		}
	}
	return 0
}
