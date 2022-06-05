package string

func removeDuplicates(s string) string {
	if len(s) == 0 {
		return ""
	}
	var stack []rune
	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, c)
			continue
		}
		if c == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, c)
	}
	return string(stack)
}
