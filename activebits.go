package piscine

func ActiveBits(n int) int {
	count := 0
	if n < 0 {
		n *= -1
		count++
	}

	for n != 0 {
		if n&1 == 1 {
			count++
		}
		n = n >> 1
	}
	return count
}
