package main

import (
  "github.com/01-edu/z01"
)

func main() {
  var c rune = 'a'
  var n rune = '\n'
  for c != 'z' {
    z01.PrintRune(c)
    c++
  }
  z01.PrintRune(c)
  z01.PrintRune(n)
}
