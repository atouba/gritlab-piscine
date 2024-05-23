package piscine

func Rot14(s string) string {
	b := []rune(s)

	for i, v := range s {
		if v > rune('z'-14) && v <= 'z' {
			b[i] = 'a' + 14 - ('z' - v) - 1
		} else if v > rune('Z'-14) && v <= 'Z' {
			b[i] = 'A' + 14 - ('Z' - v) - 1
		} else if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			b[i] = v + 14
		} else {
			b[i] = v
		}
	}
	//   return string(b)
	return string(b)
}
