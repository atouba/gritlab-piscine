package piscine

import "github.com/01-edu/z01"

func addElement(slice []int, elem int) []int {
  n := len(slice)
  slice = slice[0 : n + 1]
  slice[n] = elem
  return slice
}

// n > 0
func putDigitsInArr(n int) []int {
  var arr = make([]int, 1)
  arr[0] = n % 10
  n /= 10

  for n > 0 {
    arr = addElement(arr, n % 10)
    n /= 10
  }
  return arr
}

func PrintNbrInOrder(n int) {
  var arr = putDigitsInArr(n)
  SortIntegerTable(arr)
  for _, v := range arr {
    z01.PrintRune(rune(v + '0'))
  }
}

