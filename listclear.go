package piscine

// why having it := l.Head, and using it
// doesn't work
func ListClear(l *List) {
	for l.Head != nil {
		temp := l.Head.Next
		l.Head = nil
		l.Head = temp
	}
	l.Tail = nil
	l = nil
}
