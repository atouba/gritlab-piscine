package piscine

func IsAlpha(s string) bool {
	for _, v := range s {
		if (v < '0' || v > '9') && (v < 'a' || v > 'z') && (v < 'A' || v > 'Z') {
			return false
		}
	}
	return true
}
