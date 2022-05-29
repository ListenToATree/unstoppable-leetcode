package array

// Time: O(n) Space: O(1)
func findTheDifference(s string, t string) rune {
	rs, rt := []rune(s), []rune(t)
	r := rt[len(s)]

	for i, _ := range rs {
		r -= rs[i]
		r += rt[i]
	}

	return r
}
