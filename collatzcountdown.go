package piscine

func CollatzCountdown(start int) int {
	var count int

	if start <= 0 {
		return -1
	} else {
		for start != 1 {
			if start%2 == 0 {
				start /= 2
			} else {
				start = (start * 3) + 1
			}
			count++
		}
	}
	return count
}
