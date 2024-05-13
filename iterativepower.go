package piscine

func IterativePower(nb int, power int) int {
	res := 1
	for power > 0 {
		res *= nb
		power--
	}
	return res
}
