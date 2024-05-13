package piscine

func Fibonacci(nbr int) int {
	if nbr < 0 {
		return -1
	}
	if nbr <= 1 {
		return nbr
	}
	return Fibonacci(nbr-1) + Fibonacci(nbr-2)
}
