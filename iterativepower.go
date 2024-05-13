package piscine

func IterativePower(nb int, power int) int {
  if nb <= 0 || power < 0 {
    return 0
  }
  if power == 0 {
    return 1
  }
	res := 1
	for power > 0 {
		res *= nb
		power--
	}
	return res
}
