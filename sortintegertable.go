package piscine

func SortIntegerTable(table []int) {
	for i := len(table) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if table[j] > table[j+1] {
				Swap(&table[j], &table[j+1])
			}
		}
	}
}
