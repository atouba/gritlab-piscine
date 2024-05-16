package main

import (
	"os"

	"github.com/01-edu/z01"
	//   "fmt"
)

func strLen(s string) int {
	var l int = 0

	for i := range s {
		l++
		i = i
	}
	return l
}

func strRev(s string) string {
	var revstr string
	str_len := strLen(s)

	var revs []rune = make([]rune, str_len)
	var j int = 0
	for i := str_len - 1; i >= 0; i-- {
		revs[j] = rune(s[i])
		j++
	}

	revstr = string(revs)
	return revstr
}

func bigStr(args []string) string {
	var bigStr string

	for i, v := range args {
		if i != 0 {
			bigStr += " " + v
		} else {
			bigStr += v
		}
	}
	return bigStr
}

func isVowel(c rune) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
		return true
	}
	return false
}

// This is used to return the first vowel in the reversed string
// to replace the vowel in the normal string with the vowel in the
// reversed one (respectively)
func firstVowel(arr []rune) rune {
	for i, v := range arr {
		if isVowel(v) {
			arr[i] = 'x'
			return v
		}
	}
	return 0
}

func main() {
	var bStr []rune
	var bStrRev []rune

	if len(os.Args) > 1 {
		bStr = []rune(bigStr(os.Args[1:]))
		bStrRev = []rune(strRev(string(bStr)))
		for i, v := range bStr {
			if isVowel(v) {
				bStr[i] = firstVowel(bStrRev)
			}
			z01.PrintRune(bStr[i])
		}
	}
	z01.PrintRune('\n')
}
