package piscine

func StrRev(s string) string {
	var revstr string
	var str_len = StrLen(s)

	var revs []rune = make([]rune, str_len)
	var j int = 0
	for i := str_len - 1; i >= 0; i-- {
		revs[j] = rune(s[i])
		j++
	}

	revstr = string(revs)
	return revstr
}
