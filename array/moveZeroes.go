package array

func moveZeroes(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		nums[index] = nums[i]
		index++
	}
	for ; index < len(nums); index++ {
		nums[index] = 0
	}
}
