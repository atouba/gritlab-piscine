package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
  var ret rune
  if nb < 0 {
    ret = 'T'
  } else {
    ret = 'F'
  }
  z01.PrintRune(ret)
  ret = '\n'
  z01.PrintRune(ret)
}

