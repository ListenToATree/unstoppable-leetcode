package array

func runningSum(nums []int) []int {
	res := make([]int, len(nums))
	sum := 0

	for i := 0; i < len(nums); i++ {
		res[i] = sum + nums[i]
		sum += nums[i]
	}

	return res
}
