package piscine

func Index(s string, toFind string) int {
	for i := range s {
		if len(s[i:]) >= len(toFind) && Compare(s[i:i+len(toFind)], toFind) == 0 {
			return i
		}
	}
	return -1
}
