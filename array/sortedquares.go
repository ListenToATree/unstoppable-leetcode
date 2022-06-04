package array

// Time: O(n) Space: O(n)
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	l, r := 0, len(nums)-1

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[l] >= nums[r] {
			res[i] = nums[l] * nums[l]
			l++
		} else {
			res[i] = nums[r] * nums[r]
			r--
		}
	}
	return res
}
