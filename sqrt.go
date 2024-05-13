package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	res := 0
	for (res * res) < nb {
		res++
	}
	if res*res == nb {
		return res
	}
	return 0
}
