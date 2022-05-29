package array

import "sort"

//// Time: O(n^2) Space: O(1)
//// brute force
//func twoSumLessThanK(nums []int, k int) int {
//	sort.Ints(nums)
//	max := -1
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			sum := nums[i] + nums[j]
//			if sum >= k {
//				break
//			} else if sum > max {
//				max = sum
//			}
//		}
//	}
//	return max
//}

// Time: O(nlogn) Space: O(n)
func twoSumLessThanK(nums []int, k int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	max := -1
	for l < r {
		sum := nums[l] + nums[r]
		if sum < k {
			max = Max(sum, max)
			l++
		} else {
			r--
		}
	}
	return max
}
