package piscine

func ConcatParams(args []string) string {
	var res string = args[0]

	for i := 1; i < len(args); i++ {
		res += "\n" + args[i]
	}
	return res
}
