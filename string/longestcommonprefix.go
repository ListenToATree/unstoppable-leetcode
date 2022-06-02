package string

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, el := range strs {
		prefix = lcpHelper(prefix, el)
	}
	return prefix
}

func lcpHelper(s1, s2 string) string {
	result := make([]byte, 0)
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			break
		}
		result = append(result, s1[i])
	}
	s := string(result)
	return s
}
