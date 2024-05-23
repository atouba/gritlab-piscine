package piscine

func AdvancedSortWordArr(table []string, f func(a, b string) int) {
	for i := len(table) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if f(table[j], table[j+1]) > 0 {
				swap(&table[j], &table[j+1])
			}
		}
	}
}
