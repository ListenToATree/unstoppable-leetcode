package dp

func tribonacci(n int) int {
	tibs := make([]int, 38)
	tibs[0], tibs[1], tibs[2] = 0, 1, 1
	if n < 3 {
		return tibs[n]
	}
	for i := 3; i <= n; i++ {
		tibs[i] = tibs[i-1] + tibs[i-2] + tibs[i-3]
	}
	return tibs[n]
}
