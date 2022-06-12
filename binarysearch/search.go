package binarysearch

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		pivot := left + (right-left)/2

		if nums[pivot] > target {
			right = pivot - 1
		} else if nums[pivot] < target {
			left = pivot + 1
		} else {
			return pivot
		}

	}

	return -1

}
