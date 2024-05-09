package piscine

func UltimateDivMod(a *int, b *int) {
	d := *a / *b
	*b = *a % *b
	*a = d
}
