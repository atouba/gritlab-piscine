package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var arr []int = make([]int, max-min)
	for i := 0; min < max; i++ {
		arr[i] = min
		min++
	}
	return arr
}
