package piscine

// func CompStr(a, b interface{}) bool {
// 	return a == b
// }

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	for curr := l.Head; curr != nil; curr = curr.Next {
		if comp(curr.Data, ref) {
			return &ref
		}
	}
	return nil
}
