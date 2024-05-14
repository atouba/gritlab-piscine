package piscine

func IsPrintable(s string) bool {
	for _, v := range s {
		if v < 32 {
			return false
		}
	}
	return true
}
