package piscine

func compareXY(x, y int) int {
	if x > y {
		return 1
	} else if x < y {
		return -1
	}
	return 0
}

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a)-2; i++ {
		if f(a[i], a[i+1]) != f(a[i+1], a[i+2]) {
			return false
		}
	}
	return true
}
