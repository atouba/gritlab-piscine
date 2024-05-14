package piscine

import "github.com/01-edu/z01"

func isUnique(base string) bool {
	for i, v := range base {
		if Index(string(base[i+1:]), string(v)) != -1 {
			return false
		}
	}
	return true
}

func isBaseValid(base string) bool {
	if StrLen(base) < 2 || !isUnique(base) {
		return false
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	base_length := StrLen(base)

	if !isBaseValid(base) {
		PrintStr("NV")
		return
	}
	if nbr < 0 {
		z01.PrintRune('-')
		nbr *= -1
	}
	if nbr >= base_length {
		PrintNbrBase(nbr/base_length, base)
		z01.PrintRune(rune(base[nbr%base_length]))
	} else if nbr < base_length {
		z01.PrintRune(rune(base[nbr]))
	}
}
