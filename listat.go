package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	var count int

	for it := l; it != nil; it = it.Next {
		if count == pos {
			return it
		}
		count++
	}
	return nil
}
