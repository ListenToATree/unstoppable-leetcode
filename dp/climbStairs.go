package dp

//func climbStairs(n int) int {
//	stairs := make([]int, 46)
//	if n < 2 {
//		return n
//	}
//	stairs[1], stairs[2] = 1, 2
//	for i := 3; i <= n; i++ {
//		stairs[i] = stairs[i-1] + stairs[i-2]
//	}
//	return stairs[n]
//}

func climbStairs(n int) int {
	c := map[int]int{0: 1, 1: 1}
	return dp4Climb(n, n, c)
}

func dp4Climb(i int, n int, c map[int]int) int {
	if i <= 1 {
		return 1
	}

	if v, ok := c[i]; ok {
		return v
	}

	c[i] = dp4Climb(i-2, n, c) + dp4Climb(i-1, n, c)
	return c[i]
}
