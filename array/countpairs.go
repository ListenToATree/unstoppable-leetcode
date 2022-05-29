package array

func countPairs(nums []int, k int) int {
	ans, n := 0, len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] && (i*j)%k == 0 {
				ans++
			}
		}
	}
	return ans
}
