package piscine

func ListSize(l *List) int {
	var count int

	for it := l.Head; it != nil; it = it.Next {
		count++
	}
	return count
}
