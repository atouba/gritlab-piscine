package piscine

import "github.com/01-edu/z01"

func addElement(slice []int, elem int) []int {
	n := len(slice)
	nslice := make([]int, n+1)
	copySlice(slice, nslice)
	nslice[n] = elem
	return nslice
}

func copySlice(slice1 []int, slice2 []int) {
	for i, v := range slice1 {
		slice2[i] = v
	}
}

func putDigitsInArr(n int) []int {
	arr := make([]int, 1)
	arr[0] = n % 10
	n /= 10

	for n > 0 {
		arr = addElement(arr, n%10)
		n /= 10
	}
	return arr
}

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	arr := putDigitsInArr(n)
	SortIntegerTable(arr)
	for _, v := range arr {
		z01.PrintRune(rune(v + '0'))
	}
}
