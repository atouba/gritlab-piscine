package main

import "fmt"
import "os"

func main() {
  if len(os.Args) <= 1 {
    fmt.Println("File name missing")
  } else if len(os.Args) > 2 {
    fmt.Println("Too many arguments")
  } else {
    data, err := os.ReadFile(os.Args[1])
    if err != nil {
      fmt.Print(string(data))
    } else {
//       fmt.Print(string(err))
    }
  }
}

