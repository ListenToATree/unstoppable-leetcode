package array

func countHillValley1(nums []int) int {
	count, pre := 0, nums[0]

	for i := 1; i < len(nums)-2; i++ {
		cur, next := nums[i], nums[i+1]
		if pre < cur && cur > next {
			count++
		} else if pre > cur && cur < next {
			count++
		} else if cur == next {
			nums[i+1] = nums[i]
			continue
		}
		pre = cur
	}
	return count
}

func countHillValley(nums []int) int {
	result, n := 0, len(nums)

	for i := 1; i < n-1; i++ {
		prev, next := i-1, i+1
		for prev >= 0 && nums[i] == nums[prev] {
			prev--
		}
		for next < n && nums[i] == nums[next] {
			i++
			next++
		}
		if prev >= 0 &&
			next < n &&
			((nums[prev] < nums[i] && nums[i] > nums[next]) ||
				(nums[prev] > nums[i] && nums[i] < nums[next])) {
			result++
		}
	}

	return result
}
