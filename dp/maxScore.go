package dp

import "math"

func maximumScore(nums []int, multipliers []int) int {
	m := len(multipliers)
	memo := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		memo[i] = make([]int, m+1)
		for j := 0; j < m+1; j++ {
			memo[i][j] = math.MinInt64
		}
	}
	return dfs(nums, 0, 0, multipliers, memo)
}

func dfs(nums []int, i int, j int, steps []int, memo [][]int) int {
	if i+j >= len(steps) {
		return 0
	}

	if memo[i][j] != math.MinInt64 {
		return memo[i][j]
	}

	res := math.MinInt64
	res = max(res, nums[i]*steps[i+j]+dfs(nums, i+1, j, steps, memo))
	res = max(res, nums[len(nums)-1-j]*steps[i+j]+dfs(nums, i, j+1, steps, memo))

	memo[i][j] = res

	return res
}
