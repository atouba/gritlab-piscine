package piscine

func Join(strs []string, sep string) string {
	var res string

	for i, v := range strs {
		res += v
		if i < len(strs)-1 {
			res += sep
		}
	}
	return res
}
