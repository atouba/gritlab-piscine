package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb < 2 {
		return 1
	}
	res := nb
	nb--
	for nb > 1 {
		res *= nb
		nb--
	}
	return res
}
