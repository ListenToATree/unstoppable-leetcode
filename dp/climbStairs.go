package dp

func climbStairs(n int) int {
	stairs := make([]int, 46)
	if n < 2 {
		return n
	}
	stairs[1], stairs[2] = 1, 2
	for i := 3; i <= n; i++ {
		stairs[i] = stairs[i-1] + stairs[i-2]
	}
	return stairs[n]
}
