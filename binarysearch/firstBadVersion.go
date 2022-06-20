package binarysearch

func firstBadVersion(n int) int {
	l, r := 1, n
	for l < r {
		m := l + (r-l)/2
		if isBadVersion(m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func isBadVersion(version int) bool {
	if version >= 4 {
		return true
	}
	return false
}
