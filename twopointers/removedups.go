package twopointers

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, prev := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == prev {
			continue
		}
		nums[slow] = nums[i]
		slow++
		prev = nums[i]
	}
	return slow
}
