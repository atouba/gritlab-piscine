package main

import "github.com/01-edu/z01"

func  printcomma() {
  var comma rune

  comma = ','
  z01.PrintRune(comma)
}

func  printspace() {
  var space rune

  space = ' '
  z01.PrintRune(space)
}

func  printn(n int) {
  var x, y rune

  x = rune('0' + (n / 10));
  y = rune('0' + (n % 10));
  z01.PrintRune(x)
  z01.PrintRune(y)
}

func PrintComb2() {
  x := 0
  y := 0

  for ; x <= 98; x++ {
    for y = x + 1; y <= 99; y++ {
      printn(x)
      printspace()
      printn(y)
      if x != 98 {
        printcomma()
        printspace()
      }
    }
  }
      
}

func main() {
  PrintComb2()
}

