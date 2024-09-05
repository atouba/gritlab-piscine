package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	temp := l.Head
	for ; l.Head != nil && l.Head.Next != nil; l.Head = l.Head.Next {
		for l.Head.Next != nil && l.Head.Next.Data == data_ref {
			l.Head.Next = l.Head.Next.Next
		}
	}
	l.Head = temp
}
