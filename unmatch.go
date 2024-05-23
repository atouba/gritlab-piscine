package piscine

func count(a []int, n int) int {
	var counter int

	for _, v := range a {
		if v == n {
			counter++
		}
	}
	return counter
}

func Unmatch(a []int) int {
	res_default := -1

	for _, v := range a {
		if count(a, v)%2 == 1 {
			return v
		}
	}
	return res_default
}
