package piscine

func ListReverse(l *List) {
	for i := 0; i < ListSize(l)/2; i++ {
		node1 := ListAt(l.Head, i)
		node2 := ListAt(l.Head, ListSize(l)-i-1)

		data1 := node1.Data
		data2 := node2.Data

		temp := data1
		node1.Data = data2
		node2.Data = temp
	}
}
