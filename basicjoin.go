package piscine

func BasicJoin(elems []string) string {
	var res string

	for _, v := range elems {
		res += v
	}
	return res
}
