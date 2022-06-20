package string

func lengthOfLastWord(s string) int {
	space, result := 0, 0
	for i, n := range s {
		if n == ' ' {
			space = i + 1
		} else {
			result = i + 1 - space
		}
	}
	return result
}
