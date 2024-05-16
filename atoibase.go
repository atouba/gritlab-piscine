package piscine

func strLen(s string) int {
	var l int = 0

	for i := range s {
		l++
		i = i
	}
	return l
}

func nRune(s string, n int) rune {
	if n <= 0 || n > len(s) {
		return 0
	}
	for i, v := range s {
		if i == n-1 {
			return v
		}
	}
	return 0
}

func compare(a, b string) int {
	var j int = 0
	for _, v := range a {
		if v > nRune(b, j+1) {
			return 1
		} else if v < nRune(b, j+1) {
			return -1
		}
		j++
	}
	return 0
}

func index_b(s string, toFind string) int {
	for i := range s {
		if len(s[i:]) >= len(toFind) && compare(s[i:i+len(toFind)], toFind) == 0 {
			return i
		}
	}
	return -1
}

func is_unique(base string) bool {
	for i, v := range base {
		if index_b(string(base[i+1:]), string(v)) != -1 {
			return false
		}
	}
	return true
}

func is_baseValid(base string) bool {
	if strLen(base) < 2 || !is_unique(base) || index_b(base, "+") != -1 || index_b(base, "-") != -1 {
		return false
	}
	return true
}

func AtoiBase(s string, base string) int {
	var res int = 0
	var sign int = 1
	var j int = 0

	if !is_baseValid(base) {
		return 0
	}
	for i := strLen(s) - 1; i >= 0; i-- {
		res += index_b(base, string(s[i])) * IterativePower(strLen(base), j)
		j++
	}
	return res * sign
}
