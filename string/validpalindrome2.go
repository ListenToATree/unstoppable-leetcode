package string

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return check(s, i, j-1) || check(s, i+1, j)
		}
		i++
		j--
	}
	return true
}

func check(s string, i int, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
