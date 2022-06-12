package binarysearch

func mySqrt(x int) int {
	l := 0
	r := x/2 + 1
	for l < r {
		m := l + (r-l)/2
		if m*m > x {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}
