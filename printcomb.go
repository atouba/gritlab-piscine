package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var a, b, c rune
	var e, f, g rune

	a = '0'
	b = '1'
	c = '2'

	e = ','
	f = ' '
  g = '\n'

	for ; a <= '7'; a++ {
		for b = a + 1; b <= '8'; b++ {
			for c = b + 1; c <= '9'; c++ {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)

				if a != '7' {
					z01.PrintRune(e)
					z01.PrintRune(f)
				} else {
          z01.PrintRune(g)
        }
			}
		}
	}
}
