package piscine

func convertToBase(n int, base string, res *[]rune) {
	if n < StrLen(base) {
		*res = append(*res, rune(base[n]))
	} else {
		convertToBase(n/StrLen(base), base, res)
		*res = append(*res, rune(base[n%StrLen(base)]))
	}
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	var res []rune

	nbr = ToUpper(nbr)
	n := AtoiBase(nbr, baseFrom)
	convertToBase(n, baseTo, &res)
	return string(res)
}
