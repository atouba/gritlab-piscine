package piscine

func swap(s1 *string, s2 *string) {
	var temp string

	temp = *s1
	*s1 = *s2
	*s2 = temp
}

func SortWordArr(table []string) {
	for i := len(table) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if table[j] > table[j+1] {
				swap(&table[j], &table[j+1])
			}
		}
	}
}
