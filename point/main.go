package main

// func printStr(s string) {
// 	for _, r := range s {
// 		z01.PrintRune(r)
// 	}
// 	z01.PrintRune('\n')
// }

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// func setPoint(ptr point) {
// 	ptr.x = 42
// 	ptr.y = 21
// }

func main() {
	points := &point{}

	setPoint(points)

	print("x = ", points.x, ", y = ", points.y)
	println()
}
