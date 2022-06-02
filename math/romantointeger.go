package math

var roman = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	sum, i := 0, 0

	for i < len(s)-1 {
		val := s[i]
		nextVal := s[i+1]

		if val < nextVal {
			sum += int(nextVal - val)
			i += 2
		} else {
			sum += int(val)
			i++
		}
	}
	sum += int(s[i])
	return sum
}
