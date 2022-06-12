package binarysearch

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

const picked = 1

func guessNumber(n int) int {
	l, r := 0, n
	for l < r {
		m := l + (r-l)/2
		res := guess(m)
		if res == 0 {
			return m
		} else if res > 0 {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

func guess(num int) int {
	if num > picked {
		return -1
	}
	if num < picked {
		return 1
	}
	return 0
}
