package dp

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//// brute force
//// Time: O(n^2) Space: O(n)
//func maxSubArray(nums []int) int {
//	max := math.MinInt
//	for i := 0; i < len(nums); i++ {
//		cur := 0
//		for j := i; j < len(nums); j++ {
//			cur += nums[j]
//			max = Max(max, cur)
//		}
//	}
//	return max
//}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum, cur := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		cur = Max(cur+nums[i], nums[i])
		sum = Max(sum, cur)
	}
	return sum
}
