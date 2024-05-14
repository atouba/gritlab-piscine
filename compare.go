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

func Compare(a, b string) int {
	var j int = 0
	for _, v := range a {
		if v > NRune(b, j+1) {
			return 1
		} else if v < NRune(b, j+1) {
			return -1
		}
		j++
	}
	return 0
}
