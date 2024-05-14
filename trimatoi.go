package piscine

//	func index_atoi(s string) int {
//		for i := range s {
//			if s[i] != '0' {
//				break
//			}
//			return i
//		}
//		return 0
//	}
//
//	func strlen_atoi(s string) int {
//		var l int = 0
//
//		for i := range s {
//			l++
//			i = i
//		}
//		return l
//	}
//
//	func count_sign(s string) bool {
//		for i := 0; i < strlen_atoi(s); i++ {
//			if rune(s[i]) == '+' || rune(s[i]) == '-' {
//				if i != 0 {
//					return false
//				}
//			} else if rune(s[i]) < '0' || rune(s[i]) > '9' {
//				return false
//			}
//		}
//		return true
//	}
func isNum(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func TrimAtoi(s string) int {
	var res int = 0
	var sign int = 1

	for i := 0; i < strlen_atoi(s); i++ {
		if rune(s[i]) == '-' {
			sign = -1
			continue
		}
		if isNum(rune(s[i])) {
			res *= 10
			res += int(rune(s[i]) - '0')
		}
	}

	return res * sign
}
