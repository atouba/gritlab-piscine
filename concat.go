package piscine

func Concat(str1 string, str2 string) string {
	var i int
	res := make([]byte, len([]byte(str1))+len([]byte(str2)))

	for i = range str1 {
		res[i] = str1[i]
	}
	for j := range str2 {
		res[i] = str2[j]
		i++
	}
	return string(res)
}
