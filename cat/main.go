package main

import (
	"fmt"
  "github.com/01-edu/z01"
	"os"
)

func printStr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}

func printt(strs ...string) {
  for _, v := range strs {
    printStr(v)
  }
}

func main() {
  l := len(os.Args)
  if l == 1 {
  } else {
    for i := 1; i < l; i++ {
      file, err := os.Open(os.Args[i])
      if err != nil {
        printt("ERROR: ", err.Error(), "\n")
        os.Exit(1)
      } else {
        fmt.Printf("Type: %T\n", file)
        fmt.Println(file)
      }
    }
  }
}

