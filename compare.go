package piscine

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
