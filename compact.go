package piscine

func isZero(s string) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}

func countNonZeros(strs *[]string) int {
	var count int

	for _, v := range *strs {
		if !isZero(v) {
			count++
		}
	}
	return count
}

func Compact(ptr *[]string) int {
	nonZeros := countNonZeros(ptr)
	if nonZeros == 0 {
		*ptr = []string{}
		return 0
	} else {
		newPtr := make([]string, nonZeros)
		j := 0
		for _, v := range *ptr {
			if isZero(v) {
				continue
			}
			newPtr[j] = v
			j++
		}
		*ptr = newPtr
	}
	return nonZeros
}
