package main

import (
	"os"

	"github.com/01-edu/z01"
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

func clearB(b *[1024]byte, n int) {
	for i := 0; i < n; i++ {
		b[i] = 0
	}
}

func main() {
	l := len(os.Args)
	var maxi int = 0

	if l == 1 {
		var buffer [1024]byte
		n, _ := os.Stdin.Read(buffer[:])
		printt(string(buffer[0:n]))
		clearB(&buffer, n)
		maxi = n
		for n > 0 {
			n, _ = os.Stdin.Read(buffer[:])
			if maxi < n {
				maxi = n
			}
			printt(string(buffer[0:n]))
			clearB(&buffer, maxi)
		}
	} else {
		for i := 1; i < l; i++ {
			data, err := os.ReadFile(os.Args[i])
			if err == nil {
				printt(string(data))
			} else {
				printt("ERROR: ", err.Error(), "\n")
				os.Exit(1)
			}
		}
	}
}
