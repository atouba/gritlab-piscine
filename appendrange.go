package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var arr []int
	for min < max {
		arr = append(arr, min)
		min++
	}
	return arr
}
