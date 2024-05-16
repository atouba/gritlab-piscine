package main

import "os"
import "fmt"


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

func firstLetter(s string) rune {
  for _, v := range s {
    if v >= 'a' && v <= 'z' {
      return v
    }
  }
  return rune(0)
}

func isOption(s string) (string, rune) {
  var strs1 = []string{"--insert=", "-i="}
  var strs2 = []string{"--order", "-o", "--help", "-h"}

  for _, v := range strs1 {
    if index(s, v) == 0 {
      return s[len(v):], 'i'
    }
  }
  for _, v := range strs2 {
    if Compare(s, v) == 0 {
      return "", firstLetter(s)
    }
  }
  return s, 0
}

func help() {
  fmt.Println("--insert")
  fmt.Println("  -i")
  fmt.Println("\t This flag inserts the string into the string passed as argument.")

  fmt.Println("--order")
  fmt.Println("  -o")
  fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func firstNonZeroChar(s []rune) (rune, int) {
  for i, v := range s {
    if v > 0 {
      return v, i
    }
  }
  return 0, 0
}

func minimumRuneIndex(s []rune) (rune, int) {
  var minim rune
  var indexMinim int = 0

  minim, indexMinim = firstNonZeroChar(s);
  for i, v := range s {
    if minim > v && v != 0 {
      minim = v
      indexMinim = i
    }
  }
  return minim, indexMinim
}

func realLength(s []rune) int {
  var count int = 0

  for _, v := range s {
    if v > 0 {
      count++
    }
  }
  return count
}

func printSortedString(s []rune) {
  var minim rune
  var indexMinim int

  for realLength(s) > 0 {
    minim, indexMinim = minimumRuneIndex(s)
    fmt.Print(string(minim))
    s[indexMinim] = 0
  }
}

func main() {
  var str string
  var bigStr string
  var op rune
  var previousOp rune
  var isOrder bool

  if len(os.Args) == 1 {
    help()
    return
  } else {
    for i := 1; i < len(os.Args); i++ {
      str, op = isOption(os.Args[i])
      if op == 0 && previousOp == 'i' {
        bigStr = str + bigStr
      } else if op == 'i' || op == 0 {
        bigStr = bigStr + str
      } else if op == 'o' {
        isOrder = true
        continue
      } else if op == 'h' {
        help()
        return
      }
      previousOp = op
    }
  }
  if isOrder {
    printSortedString([]rune(bigStr))
  } else { 
    fmt.Println(bigStr)
  }
}
