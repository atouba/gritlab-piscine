package piscine

func StrLen(s string) int {
	var l int = 0

	for i := range s {
		l++
		i = i
	}
	return l
}
