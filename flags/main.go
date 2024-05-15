package main

import "os"
// import "fmt"


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

func index(s string, toFind string) int {
	for i := range s {
		if len(s[i:]) >= len(toFind) && Compare(s[i:i+len(toFind)], toFind) == 0 {
			return i
		}
	}
	return -1
}

func isOption(s string) string {
  var strs = []string{"--insert=", "-i=", "--order", "-o", "--help", "-h"}
  for i, v := range strs {
    if index(os.Args[i], v) == 0 {
      return os.Args[i][len(v):]
    }
  }
  return ""
}

func main() {
	for i, v := range os.Args {
		if i > 0 {
		}
	}
}
