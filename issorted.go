package piscine

func compare(x, y int) int {
	if x > y {
		return 1
	} else if x < y {
		return -1
	}
	return 0
}

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := range len(a) - 1 {
		if f(a[i], a[i+1]) > 0 {
			return false
		}
	}
	return true
}
