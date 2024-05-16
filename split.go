package piscine

func StrLen(s string) int {
        var l int = 0

        for i := range s {
                l++
                i = i
        }
        return l
}

func NRune(s string, n int) rune {
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

func Compare(a, b string) int {
        var j int = 0
        for _, v := range a {
                if v > NRune(b, j+1) {
                        return 1
                } else if v < NRune(b, j+1) {
                        return -1
                }
                j++
        }
        return 0
}

func Index(s string, toFind string) int {
        for i := range s {
                if len(s[i:]) >= len(toFind) && Compare(s[i:i+len(toFind)], toFind) == 0 {
                        return i
                }
        }
        return -1
}

func Split(s, sep string) []string {
	var arr []string

	for i := 0; i < StrLen(s); {
		for Compare(s[i:i+len(sep)], sep) == 0 {
			i += len(sep)
		}

		if length_word := Index(s[i:], sep); length_word != -1 {
			arr = append(arr, s[i:i+length_word])
		} else {
			arr = append(arr, s[i:])
			break
		}
		i += Index(s[i:], sep)
	}
	return arr
}
