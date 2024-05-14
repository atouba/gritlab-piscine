package piscine

func Concat(str1 string, str2 string) string {
	var i int
	res := make([]rune, len(str1)+len(str2))

	for i = range str1 {
		res[i] = NRune(str1, i+1)
	}
	for j := range str2 {
		res[i] = NRune(str2, j+1)
		i++
	}
	return string(res)
}
