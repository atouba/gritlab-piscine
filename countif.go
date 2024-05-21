package piscine

func CountIf(f func(string) bool, tab []string) int {
	var count int

	for _, v := range tab {
		if f(v) {
			count++
		}
	}
	return count
}
