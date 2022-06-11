package dp

func rob(nums []int) int {
	cache := make(map[int]int)
	return dp(nums, cache, len(nums)-1)
}

func dp(nums []int, cache map[int]int, i int) int {
	if i == 0 {
		return nums[0]
	}
	if i == 1 {
		return max(nums[0], nums[1])
	}
	if _, ok := cache[i]; !ok {
		cache[i] = max(dp(nums, cache, i-1), dp(nums, cache, i-2)+nums[i])
	}
	return cache[i]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
