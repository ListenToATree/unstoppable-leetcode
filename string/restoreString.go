package string

func restoreString(s string, indices []int) string {
	res := make([]rune, len(s))
	for i, c := range s {
		res[indices[i]] = c
	}
	return string(res)
}
