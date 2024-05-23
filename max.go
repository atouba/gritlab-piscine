package piscine

func Max(a []int) int {
	maxx := 0

	if len(a) > 0 {
		maxx = a[0]
	}

	for _, v := range a {
		if maxx < v {
			maxx = v
		}
	}
	return maxx
}
