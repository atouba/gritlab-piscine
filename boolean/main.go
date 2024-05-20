package main

import "os"
import "github.com/01-edu/z01"

type boolean = int

var (
  lengthOfArg int = len(os.Args[1:])
  EvenMsg = "I have an even number of arguments"
  OddMsg = "I have an odd number of arguments"
  yes = 1
  no = 0
)

func even(nbr int) int {
  if nbr % 2 == 0 {
    return 1
  }
  return 0
}

// -------------------------------------------

func printStr(s string) {
  for _, r := range s {
    z01.PrintRune(r)
  }
  z01.PrintRune('\n')
}

func isEven(nbr int) boolean {
  if even(nbr) == 1 {
    return yes
  } else {
    return no
  }
}

func main() {
  if isEven(lengthOfArg) == 1 {
    printStr(EvenMsg)
  } else {
    printStr(OddMsg)
  }
}
